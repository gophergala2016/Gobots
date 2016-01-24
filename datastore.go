package main

import (
	"bytes"
	"encoding/gob"
	"errors"
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
	UserBucket = []byte("Users")
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
	return nil, errDatastoreNotImplemented
}

func (db *dbImpl) lookupAIToken(token string) (*aiInfo, error) {
	return nil, errDatastoreNotImplemented
}

// Games
func (db *dbImpl) startGame(ai1, ai2 aiID, init botapi.Board) (gameID, error) {
	return "", errDatastoreNotImplemented

}

func (db *dbImpl) addRound(id gameID, round botapi.Replay_Round) error {
	return errDatastoreNotImplemented
}

func (db *dbImpl) lookupGame(id gameID) (botapi.Replay, error) {
	_, seg, _ := capnp.NewMessage(capnp.SingleSegment(nil))
	replay, _ := botapi.NewReplay(seg)
	return replay, nil
}

var errDatastoreNotImplemented = errors.New("gobots: datastore operation not implemented")
