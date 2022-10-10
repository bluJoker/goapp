package models

import (
    "crypto/sha1"
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gomodule/redigo/redis"
    "log"
    "math/rand"
)

var Db *sql.DB
var Pool *redis.Pool  //创建redis连接池

func init() {
    var err error
    Db, err = sql.Open("mysql", "root:lwyu22195840@@tcp(localhost:3306)/goapp?charset=utf8&parseTime=True")
    if err != nil {
        log.Fatal(err)
    }

    Pool = &redis.Pool{
        MaxIdle:16,                        //最初的连接数量
        // MaxActive:1000000,              //最大连接数量
        MaxActive:0,                       //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
        IdleTimeout:300,                   //连接关闭时间 300秒 （300秒不使用自动关闭）
        Dial: func() (redis.Conn ,error){  //要连接的redis数据库
            return redis.Dial("tcp","localhost:6379")
        },
    }

    return
}

// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
    u := new([16]byte)
    _, err := rand.Read(u[:])
    if err != nil {
        log.Fatalln("Cannot generate UUID", err)
    }

    // 0x40 is reserved variant from RFC 4122
    u[8] = (u[8] | 0x40) & 0x7F
    // Set the four most significant bits (bits 12 through 15) of the
    // time_hi_and_version field to the 4-bit version number.
    u[6] = (u[6] & 0xF) | (0x4 << 4)
    uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
    return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
    cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
    return
}

func rediss() {
    c := Pool.Get()
    defer c.Close()
    _, err := c.Do("Set", "username", "jack")
    if err != nil {
        fmt.Println(err)
        return
    }
    flag, err := c.Do("exists", "username")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%T\n", flag)
    fmt.Println(flag)
}
