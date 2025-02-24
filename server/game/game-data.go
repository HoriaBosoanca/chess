package game

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/teris-io/shortid"

	"mychess/chess"
)

var Games map[string]*Game = make(map[string]*Game)
var GlobalMu sync.Mutex

type Game struct {
	State     string
	GameID    string
	WhiteInit bool
	BlackInit bool
	Board     chess.Board
	GameMu    sync.Mutex
}

func CreateGame() (game *Game) {
	game = &Game{State: "pending black", GameID: shortid.MustGenerate()}
	GlobalMu.Lock()
	defer GlobalMu.Unlock()
	Games[game.GameID] = game
	return game
}

func JoinGame(gameID string) (game *Game, err error) {
	GlobalMu.Lock()
	defer GlobalMu.Unlock()
	game, exists := Games[gameID]
	if !exists {
		return &Game{}, errors.New("!Game not found")
	}
	game.GameMu.Lock()
	defer game.GameMu.Unlock()
	game.State = "waiting for white accept"
	return game, nil
}

func Write(conn *websocket.Conn, writeMutex *sync.Mutex, data interface{}) {
	writeMutex.Lock()
	defer writeMutex.Unlock()
	conn.WriteJSON(data)
}

type ErrMsg struct {
	Error string `json:"error"`
}

type Notification struct {
	Type string `json:"type"`
	Message string `json:"message"`
}