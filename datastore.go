package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io"
	"strconv"
	"time"

	"zombiezen.com/go/capnproto2"

	"github.com/boltdb/bolt"
	"github.com/gophergala2016/Gobots/botapi"
)

type datastore interface {
	// Users
	createUser(u userID) error
	loadUser(u userID) (*player, error)

	// AIs
	createAI(u userID, info *aiInfo) (id aiID, token string, err error)
	listAIsForUser(u userID) ([]*aiInfo, error)
	lookupAI(id aiID) (*aiInfo, error)
	lookupAIToken(token string) (*aiInfo, error)

	// Games
	startGame(ai1, ai2 aiID, init botapi.Board) (gameID, error)
	addRound(id gameID, round botapi.Replay_Round) error
	lookupGame(id gameID) (botapi.Replay, error)
}

type dbImpl struct {
	*bolt.DB
}

func initDB(dbName string) (datastore, error) {
	db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		for _, b := range [][]byte{UserBucket, GameBucket, AIBucket, TokensBucket} {
			if _, err := tx.CreateBucketIfNotExists(b); err != nil {
				return err
			}
		}

		return nil
	})

	return &dbImpl{db}, err
}

// userID is just the user's GitHub API access token. Let's hope it's unique
// enough #YOLO
type userID string

type aiID string

type gameID string

type aiInfo struct {
	id    aiID
	Nick  string
	token string

	wins   int
	losses int
}

var (
	UserBucket   = []byte("Users")
	GameBucket   = []byte("Games")
	AIBucket     = []byte("AI")
	TokensBucket = []byte("AISecretToken")
)

func (db *dbImpl) createUser(uID userID) error {
	u := username(uID)
	p := player{
		Name: u,
	}

	return db.Update(func(tx *bolt.Tx) error {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)

		if err := enc.Encode(p); err != nil {
			return err
		}

		b := tx.Bucket(UserBucket)
		return b.Put([]byte(uID), buf.Bytes())
	})
}

func (db *dbImpl) loadUser(uID userID) (*player, error) {
	var p *player
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(UserBucket)
		dat := b.Get([]byte(uID))

		buf := bytes.NewReader(dat)
		dec := gob.NewDecoder(buf)

		return dec.Decode(&p)
	})

	return p, err
}

// AIs
func (db *dbImpl) createAI(u userID, info *aiInfo) (id aiID, token string, err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(AIBucket)
		idNum, err := b.NextSequence()
		if err != nil {
			return err
		}
		id = aiID(strconv.FormatUint(idNum, 10))
		token = genName(32)
		newInfo := &aiInfo{
			id:    id,
			Nick:  info.Nick,
			token: token,
		}
		var buf bytes.Buffer
		if err := gob.NewEncoder(&buf).Encode(newInfo); err != nil {
			return err
		}
		if err := b.Put([]byte(id), buf.Bytes()); err != nil {
			return err
		}

		tb := tx.Bucket(TokensBucket)
		if err := tb.Put([]byte(token), []byte(id)); err != nil {
			return err
		}

		return nil
	})
	return
}

func (db *dbImpl) listAIsForUser(u userID) ([]*aiInfo, error) {
	return nil, errDatastoreNotImplemented
}

func (db *dbImpl) lookupAI(id aiID) (*aiInfo, error) {
	var info *aiInfo
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(AIBucket)
		dat := b.Get([]byte(id))

		buf := bytes.NewReader(dat)
		return gob.NewDecoder(buf).Decode(&info)
	})
	return info, err
}

func (db *dbImpl) lookupAIToken(token string) (*aiInfo, error) {
	var info *aiInfo
	err := db.View(func(tx *bolt.Tx) error {
		tb := tx.Bucket(TokensBucket)
		idBytes := tb.Get([]byte(token))
		if len(idBytes) == 0 {
			// TODO not found error
		}

		b := tx.Bucket(AIBucket)
		dat := b.Get(idBytes)

		buf := bytes.NewReader(dat)
		return gob.NewDecoder(buf).Decode(&info)
	})
	return info, err
}

// Games
func (db *dbImpl) startGame(ai1, ai2 aiID, init botapi.Board) (gameID, error) {
	var gID gameID
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(GameBucket)

		msg, s, err := capnp.NewMessage(capnp.SingleSegment(nil))
		if err != nil {
			return err
		}
		r, err := botapi.NewRootReplay(s)
		if err != nil {
			return err
		}
		gID = newGameID(b)
		r.SetGameId(string(gID))
		r.SetInitialBoard(init)

		data, err := msg.Marshal()
		if err != nil {
			return err
		}

		return b.Put([]byte(gID), data)
	})

	return gID, err

}

func (db *dbImpl) addRound(id gameID, round botapi.Replay_Round) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(GameBucket)
		key := []byte(id)
		data := b.Get(key)
		msg, err := capnp.Unmarshal(copyBytes(data))
		if err != nil {
			return err
		}
		orig, err := botapi.ReadRootReplay(msg)
		if err != nil {
			return err
		}
		newMsg, err := addReplayRound(orig, round)
		if err != nil {
			return err
		}
		newData, err := newMsg.Marshal()
		if err != nil {
			return err
		}
		return b.Put(key, newData)
	})
}

func addReplayRound(orig botapi.Replay, round botapi.Replay_Round) (*capnp.Message, error) {
	newMsg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}
	newReplay, err := botapi.NewRootReplay(seg)
	if err != nil {
		return nil, err
	}
	gid, err := orig.GameId()
	if err != nil {
		return nil, err
	}
	if err := newReplay.SetGameId(gid); err != nil {
		return nil, err
	}
	initBoard, err := orig.InitialBoard()
	if err != nil {
		return nil, err
	}
	if err := newReplay.SetInitialBoard(initBoard); err != nil {
		return nil, err
	}
	origRounds, err := orig.Rounds()
	if err != nil {
		return nil, err
	}
	rounds, _ := botapi.NewReplay_Round_List(seg, int32(origRounds.Len())+1)
	for i := 0; i < origRounds.Len(); i++ {
		if err := rounds.Set(i, origRounds.At(i)); err != nil {
			return nil, err
		}
	}
	if err := rounds.Set(rounds.Len()-1, round); err != nil {
		return nil, err
	}
	return newMsg, nil
}

func (db *dbImpl) lookupGame(id gameID) (botapi.Replay, error) {
	var r botapi.Replay
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(GameBucket)
		data := b.Get([]byte(id))

		msg, err := capnp.Unmarshal(copyBytes(data))
		if err != nil {
			return err
		}
		r, err = botapi.ReadRootReplay(msg)
		return err
	})

	if err == io.EOF {
		err = errDatastoreNotFound
	}
	return r, err
}

func copyBytes(b []byte) []byte {
	bb := make([]byte, len(b))
	copy(bb, b)
	return bb
}

var errDatastoreNotImplemented = errors.New("gobots: datastore operation not implemented")
var errDatastoreNotFound = errors.New("gobots: datastore entity not found")

// NOTE: Definitely definitely only call this from inside a transaction
func newGameID(b *bolt.Bucket) gameID {
	id, err := b.NextSequence()
	if err != nil {
		// Screw it, they're getting a random ID
		return gameID(genName(64))
	}
	return gameID(strconv.FormatUint(id, 16))
}
