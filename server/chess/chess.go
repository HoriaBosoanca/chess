package chess

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func HandleEndpoint(router *mux.Router) {
	router.HandleFunc("/ws", play)
	
	router.HandleFunc("/ws", OptionsHandler).Methods("OPTIONS")
}

// ws://domain/ws?gameID=create - create joinable game
// ws://domain/ws?gameID=abcd - join created game
func play(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if(err != nil) {
		log.Println("!Error upgrading:", err)
		return
	}

	var game *Game
	var myPlayer *Player
	gameID := r.URL.Query().Get("gameID")
	if gameID == "create" {
		game = CreateGame()
		myPlayer = MakePlayer(connection, "white")
	} else {
		gameToJoin, err := JoinGame(gameID)
		if err != nil {
			log.Println("!Error joining game:", err)
			return
		}
		game = gameToJoin
		myPlayer = MakePlayer(connection, "black")
	}

	// write game id
	myPlayer.Write(GameID{GameID: game.GameID})

	for {
		switch game.State {
		case "waiting for white accept":
			if myPlayer.Color == "white" {
				myPlayer.Write(GameStart{YourColor: "white"})
				game.State = "waiting for black accept"
			}
		case "waiting for black accept":
			if myPlayer.Color == "black" {
				myPlayer.Write(GameStart{YourColor: "black"})
				game.State = "white"
			}
		case myPlayer.Color:
			myPlayer.Write("it's your turn")
			var move Move
			myPlayer.Conn.ReadJSON(&move)

			if myPlayer.Color == "white" {
				game.State = "black"
			}
			if myPlayer.Color == "black" {
				game.State = "white"
			}
		}
	}
}

type GameID struct {
	GameID string `json:"gameID"`
}
type GameStart struct {
	YourColor string `json:"yourColor"`
}
type Move struct {
	Move string `json:"move"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}