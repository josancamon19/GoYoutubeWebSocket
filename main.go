package main

import (
	"GoWebSocketsYoutube/data"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// The Upgrade function will take in an incoming request and upgrade the request
// into a websocket connection
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// this line allows other origin hosts to connect to our
	// websocket server
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// creates our websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	// returns our new websocket connection
	return ws, nil
}

func Writer(conn *websocket.Conn) {
	for {
		ticker := time.NewTicker(5 * time.Second)
		for t := range ticker.C {
			fmt.Printf("Updating Stats: %+v\n", t)
			items, err := data.GetSubscribers()
			if err != nil {
				panic(err)
			}
			// next we marshal our response into a JSON string
			jsonString, err := json.Marshal(items)
			if err != nil {
				fmt.Println(err)
			}
			if err := conn.WriteMessage(websocket.TextMessage, jsonString); err != nil {
				fmt.Println(err)
				return
			}

		}
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"static/index.html")
}

func PubStats(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrade(w, r)
	if err != nil {
		fmt.Fprint(w, err)
	}

	go Writer(ws)
}
func setupRoutes() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/stats", PubStats)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	setupRoutes()
}
