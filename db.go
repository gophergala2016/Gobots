package main

import "github.com/gophergala2016/Gobots/botapi"

type datastore interface {
	// Users
	createUser( /* user ID/creds */ ) error

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
