// client project main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func ReceiveFromServer(conn *websocket.Conn) {
	//buffer := make([]byte, 1400)

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			break
		} else if msg == nil {
			continue
		}

		fmt.Print(">> " + string(msg))

	}
}

func SendToServer(conn *websocket.Conn, in chan string) {

	for {
		select {
		case msg := <-in:
			err := conn.WriteMessage(websocket.BinaryMessage, []byte(msg))

			if err != nil {
				log.Println(err)
			}
		}
	}

}

func main() {
	in := make(chan string)
	reader := bufio.NewReader(os.Stdin)
	SetupCommands()

	addr := "still-temple-68994.herokuapp.com"

	conn, _, err := websocket.DefaultDialer.Dial("wss://"+addr+"/ws", nil)
	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	go SendToServer(conn, in)
	go ReceiveFromServer(conn)

	for {
		msg, err := reader.ReadString('\n')

		if err != nil {
			log.Println(err)
		} else if msg == "" {
			continue
		} else if msg[0] == '#' {
			ParseCommand(msg)
			continue
		}

		in <- msg
	}

}
