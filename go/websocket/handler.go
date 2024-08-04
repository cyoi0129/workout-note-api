package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var rooms = Room{}

func ServeWs(c *gin.Context) {

	topic := c.Param("topic")
	queryUser := c.Query("user")
	if queryUser == "" {
		queryUser = "0"
	}

	log.Printf("topic: %s", topic)

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Printf("Failed to set websocket upgrade: %+v\n", err)
	}

	defer ws.Close()

	client := &Client{Ws: ws}

	rooms.AddSubscription(&Subscription{Topic: topic, Client: client})

	for {
		_, msg, err := ws.ReadMessage()

		message_str := string(msg)
		message_json := fmt.Sprintf(`{"id": %s, "message": "%s"}`, queryUser, message_str)
		message_byte := []byte(message_json)

		if err != nil {
			log.Printf("ReadMessage Error. ERROR: %+v\n", err)
			break
		}

		rooms.Publish(message_byte, topic)
	}
}
