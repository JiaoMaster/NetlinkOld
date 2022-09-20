package controler

import (
	"NetLinkOld/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"log"
	"net/http"
	"sync"
	"time"
	"unsafe"
)

func NewCommit(c *gin.Context) {
	var upgrader *websocket.Upgrader
	upgrader = &websocket.Upgrader{

		ReadBufferSize: 1024,

		WriteBufferSize: 1024,

		// 解决跨域问题

		CheckOrigin: func(r *http.Request) bool {

			return true

		},
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	defer conn.Close()
	if err != nil {

		log.Print("upgrade:", err)

		return

	}

	log.Println("ws connect...")
	log.Println(conn.RemoteAddr())
	_, receive, err := conn.ReadMessage()

	if err != nil {

		log.Println(err)

		return

	}

	log.Println("ws receive : ", string(receive))
	models.NM.S.Lock()
	models.NM.M[string(receive)] = ""
	models.NM.S.Unlock()
	connect := &Connetion{
		con: conn,
	}
	ticker := time.NewTicker(time.Second * 10)
	go connect.dance(ticker, conn, &err)
	for {

		//接受消息
		models.NM.S.Lock()
		if _, ok := models.NM.M[string(receive)]; ok && models.NM.M[string(receive)] != "" {
			connect.mutex.Lock()

			err = connect.con.WriteMessage(websocket.TextMessage, StringToBytes(models.NM.M[string(receive)]))
			models.NM.M[string(receive)] = ""

			connect.mutex.Unlock()
		}
		models.NM.S.Unlock()
		if err != nil {
			models.NM.S.Lock()
			delete(models.NM.M, string(receive))
			models.NM.S.Unlock()
			log.Println(err)
			zap.L().Error("WebSocket conn.WriteMessage err :", zap.Error(err))
			break
		}

	}
	log.Println("notice end")
}

type Connetion struct {
	con   *websocket.Conn
	mutex sync.Mutex
}

func (con *Connetion) dance(ticker *time.Ticker, c *websocket.Conn, err *error) {
	for {
		<-ticker.C
		*err = c.SetWriteDeadline(time.Now().Add(10 * time.Second))
		//fmt.Println(time.Now().Format(time.UnixDate))
		if *err != nil {
			log.Printf("ping error: %s\n", (*err).Error())
			break
		}

		con.mutex.Lock()
		if *err = c.WriteMessage(websocket.PingMessage, nil); *err != nil {
			log.Printf("ping error: %s\n", (*err).Error())
			break
		}
		con.mutex.Unlock()

	}
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
