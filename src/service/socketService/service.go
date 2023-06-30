package socketService

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"sm-access/src/service/monitoringService/usbService"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Пропускаем любой запрос
	},
}

type Server struct {
	clients       map[*websocket.Conn]bool
	handleMessage func(message []byte) // хандлер новых сообщений
}

func StartServer(handleMessage func(message []byte)) *Server {
	server := Server{
		make(map[*websocket.Conn]bool),
		handleMessage,
	}

	serverUrl := fmt.Sprintf("%s:%s", os.Getenv("SOCKET_SERVER_HOST"), os.Getenv("SOCKET_SERVER_PORT"))

	http.HandleFunc("/", server.echo)
	go http.ListenAndServe(serverUrl, nil) // Уводим http сервер в горутину

	return &server
}

func (server *Server) echo(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
	defer connection.Close()

	server.clients[connection] = true        // Сохраняем соединение, используя его как ключ
	defer delete(server.clients, connection) // Удаляем соединение

	devices := usbService.GetConnectedDevices()
	for _, device := range devices {
		server.SendToSocket([]byte(fmt.Sprintf("%s: true", device.ID)))
	}

	for {
		mt, message, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break // Выходим из цикла, если клиент пытается закрыть соединение или связь прервана
		}

		go server.handleMessage(message)
	}
}

func (server *Server) SendToSocket(message []byte) {
	for conn := range server.clients {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("SendToSocket() error: %s", err.Error())
		}
	}
}
