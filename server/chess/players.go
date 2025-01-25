package chess

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/teris-io/shortid"
)

var Games map[string]*Game = make(map[string]*Game)
var GlobalMu sync.Mutex

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
		return &Game{}, errors.New("!Game not found.")
	}
	game.GameMu.Lock()
	defer game.GameMu.Unlock()
	game.State = "waiting for white accept"
	return game, nil
}

type Game struct {
	State     string
	GameID    string
	WhiteInit bool
	BlackInit bool
	Board     [][]string
	GameMu    sync.Mutex
}

type Player struct {
	Conn     *websocket.Conn
	ID       string
	Color    string
	PlayerMu sync.Mutex
}

func (p *Player) Write(data interface{}) {
	p.PlayerMu.Lock()
	defer p.PlayerMu.Unlock()
	p.Conn.WriteJSON(data)
}

func MakePlayer(conn *websocket.Conn, color string) (player *Player) {
	player = &Player{Conn: conn, Color: color, ID: shortid.MustGenerate()}
	return player
}
