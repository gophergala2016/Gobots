package main

import (
	"bytes"
	"encoding/gob"
	"errors"
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
		_, err := tx.CreateBucketIfNotExists([]byte("Users"))
		return err
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
	nick  string
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

		buf := bytes.NewBuffer(dat)
		dec := gob.NewDecoder(buf)

		return dec.Decode(&p)
	})

	return p, err
}

// AIs
func (db *dbImpl) createAI(u userID, info *aiInfo) (id aiID, token string, err error) {
	return aiID(0), "", errDatastoreNotImplemented
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
	// TODO not found error
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
	_, err := db.lookupGame(id)
	if err != nil {
		return err
	}
	// TODO: Add round to Replay
	return errDatastoreNotImplemented
}

func (db *dbImpl) lookupGame(id gameID) (botapi.Replay, error) {
	var r botapi.Replay
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(GameBucket)
		data := b.Get([]byte(id))

		msg, err := capnp.Unmarshal(data)
		if err != nil {
			return err
		}
		r, err = botapi.ReadRootReplay(msg)
		return err
	})

	return r, err
}

var errDatastoreNotImplemented = errors.New("gobots: datastore operation not implemented")

// NOTE: Definitely definitely only call this from inside a transaction
func newGameID(b *bolt.Bucket) gameID {
	id, err := b.NextSequence()
	if err != nil {
		// Screw it, they're getting a random ID
		return gameID(genName(64))
	}
	return gameID(strconv.FormatUint(id, 16))
}
