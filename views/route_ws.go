// websockets.go
package views

import (
    "fmt"
    "net/http"
    "github.com/gorilla/websocket"
    "strconv"
    "goapp/models"
    "github.com/gomodule/redigo/redis"
    "log"
    "encoding/json"
    "github.com/bitly/go-simplejson"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type receiveData struct {
    Event string `json:"event"`
    Uuid string `json:"uuid"`
}

type sendData struct {
    Event string `json:"event"`
    Uuid string `json:"uuid"`
    Email string `json:"email"`
    Photo string `json:"photo"`
}

var conn websocket.Conn
var room_name string = ""
var rediConn redis.Conn
const RoomCapacity int = 3
var clients = make(map[string][]*websocket.Conn)

func WsConnect(w http.ResponseWriter, r *http.Request) {
    conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

    // Print the message to the console
    fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), "ws connecting!")

    rediConn = models.Pool.Get() //从连接池，取一个链接
    //defer rediConn.Close() //函数运行结束 ，把连接放回连接池

    // 找到一个可用房间
    for i := 0; i < 1000; i++ {
        name := "room-" + strconv.Itoa(i)
        if exist, _ := redis.Int(rediConn.Do("exists", name)); exist == 0 {
            room_name = name
            fmt.Println("New room: ", name)
            break
        } else {
            slen, _ := redis.Int(rediConn.Do("scard", name))
            fmt.Printf("%s, %d\n", name, slen)
            if slen < RoomCapacity {
                room_name = name
                break
            }
        }
    }

    fmt.Println(room_name)
    if room_name == "" { // 房间不够了
        return
    }

    //连接建立后，服务器向该client发送：在redis中存储的组内已有玩家信息
    vals,err := redis.Values(rediConn.Do("smembers", room_name))
    if err != nil {
        log.Fatal(err)
    }

    for _, val := range vals {
        fmt.Printf("[%T]v = %s\n", val, val)
        js, _ := simplejson.NewJson(val.([]byte))

        data := sendData{}
        data.Event = "create_player"
        data.Uuid = js.Get("uuid").MustString()
        data.Email = js.Get("email").MustString()
        data.Photo = js.Get("photo").MustString()
        fmt.Println("conn: ", conn)
        conn.WriteJSON(&data)
    }

    clients[room_name] = append(clients[room_name], conn);

    fmt.Println("clients: ", clients)

    go readMsg(conn)

    //models.Pool.Close() //关闭连接池
}

func create_player(js *simplejson.Json) {
    fmt.Println("goserver: create_player")
    type rediData struct {
        Uuid string `json:"uuid"`
        Email string `json:"email"`
        Photo string `json:"photo"`
    }
    redidata := rediData{
        Uuid: js.Get("uuid").MustString(),
        Email: js.Get("email").MustString(),
        Photo: js.Get("photo").MustString(),
    }
    rediJson, err := json.Marshal(redidata)
    if err != nil {
        fmt.Println("生成json字符串错误")
    }
    fmt.Println("rediJson: ", rediJson, "====room_name: ", room_name)
    _, err = rediConn.Do("sadd", room_name, string(rediJson))
    if err != nil {
        fmt.Println("redis.Do(sadd) err: ", err)
    }

    // [向房间内所有连接的前端]群发create_player消息
    conns := clients[room_name]
    for _, con := range conns {
        // TODO:
        //if con==m.conn{  //自己发送的信息，不用再发给自己
        //continue
        //}
        data := sendData{}
        data.Event = "create_player"
        data.Uuid = js.Get("uuid").MustString()
        data.Email = js.Get("email").MustString()
        data.Photo = js.Get("photo").MustString()
        con.WriteJSON(&data)
    }
}


func readMsg(conn *websocket.Conn) {
    for {
        _, read_data, _ := conn.ReadMessage()

        js, _ := simplejson.NewJson(read_data)

        event := js.Get("event").MustString()
        if event == "create_player" {
            fmt.Printf("%T %s\n", js.Get("event"), js.Get("event").MustString())
            fmt.Printf("%T %s\n", js.Get("uuid"), js.Get("uuid").MustString())
            fmt.Printf("%T %s\n", js.Get("email"), js.Get("email").MustString())
            fmt.Printf("%T %s\n", js.Get("photo"), js.Get("photo").MustString())
            create_player(js)
        }

        // Print the message to the console
        //fmt.Printf("[%s]%s sent: %s\n", string(msgType), conn.RemoteAddr(), string(msg))
    }
}
