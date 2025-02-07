package chess

import (
	"log"
	"net/http"
	"sync"

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
	writeMutex := sync.Mutex{}
	if(err != nil) {
		log.Println("!Error upgrading:", err)
		return
	}
	
	var game *Game
	color := ""
	gameID := r.URL.Query().Get("gameID")
	if gameID == "create" {
		game = CreateGame()
		color = "white"
	} else {
		gameToJoin, err := JoinGame(gameID)
		if err != nil {
			Write(connection, &writeMutex, ErrMsg{Error: "Invalid game ID"})
			connection.Close()
			return
		}
		game = gameToJoin
		color = "black"
	}

	// write game id
	log.Println("test")
	Write(connection, &writeMutex, struct {
		GameID string `json:"gameID"`
	}{
		GameID: game.GameID,
	})

	for {
		switch game.State {
		case "waiting for white accept":
			if color == "white" {
				Write(connection, &writeMutex, struct {
					YourColor string `json:"yourColor"`
				}{
					YourColor: "white",
				})
				game.State = "waiting for black accept"
			}
		case "waiting for black accept":
			if color == "black" {
				Write(connection, &writeMutex, struct {
					YourColor string `json:"yourColor"`
				}{
					YourColor: "black",
				})
				game.State = "white"
			}
		case color:
			Write(connection, &writeMutex, "it's your turn")
			var move struct {
				Move string `json:"move"`
			}
			connection.ReadJSON(&move)
			if color == "white" {
				game.State = "black"
			}
			if color == "black" {
				game.State = "white"
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}