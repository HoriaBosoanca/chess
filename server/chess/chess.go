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
			Write(connection, &writeMutex, Notification{
				Type: "error",
				Message: "Invalid game ID",
			})
			connection.Close()
			return
		}
		game = gameToJoin
		color = "black"
	}

	// write game id
	Write(connection, &writeMutex, Notification{
		Type: "gameID",
		Message: game.GameID,
	})

	for {
		if color == "white" {
			switch game.State {
			case "waiting for white accept":
				Write(connection, &writeMutex, Notification{
					Type: "color",
					Message: "white",
				})
				game.State = "waiting for black accept"
			case "notify start white":
				Write(connection, &writeMutex, Notification{
					Type: "start",
				})
				game.State = "notify start black"
			case "white's turn":
				Write(connection, &writeMutex, Notification{
					Type: "turn",
				})
				var move struct {
					Move string `json:"move"`
				}
				connection.ReadJSON(&move)
				// execute move
				game.State = "black"
			}
		}

		if color == "black" {
			switch game.State {
			case "waiting for black accept":
				Write(connection, &writeMutex, Notification{
					Type: "color",
					Message: "black",
				})
				game.State = "notify start white"
			case "notify start black":
				Write(connection, &writeMutex, Notification{
					Type: "start",
				})
				game.State = "white's turn"
			case "black's turn":
				Write(connection, &writeMutex, Notification{
					Type: "move",
				})
				var move struct {
					Move string `json:"move"`
				}
				connection.ReadJSON(&move)
				// execute move
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