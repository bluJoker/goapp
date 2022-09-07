package models

import (
    "time"
    "fmt"
)

type User struct {
    Id        int
    Uuid      string
    Name      string
    Email     string
    Password  string
    Photo     string
    CreatedAt time.Time
}

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
    // Postgres does not automatically return the last insert id, because it would be wrong to assume
    // you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
    // information from postgres.

    // w2: 在写SQL中，经常会有诸如更新了一行记录，之后要获取更新过的这一行: Oracle和PostgreSQL支持RETURNING，但MySQL不支持.
    //statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
    statement := "insert into users (uuid, name, email, password, photo, created_at) values (?, ?, ?, ?, ?, ?)"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()

    // w2: 分开来做，先insert
    res, err := stmt.Exec(createUUID(), user.Name, user.Email, Encrypt(user.Password), user.Photo, time.Now())
    if err != nil {
        fmt.Println(err)
    }

    id, err := res.LastInsertId() //获取刚插入的自增主键id
    if err != nil {
        //danger(err, "Get lastInsertId failed")
    }
    user.Id = int(id)

    // use QueryRow to return a row and scan the returned id into the User struct
    //err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
    //fmt.Println(user)

    // w2: 再去找刚插入的记录，users[0]只赋值了Name、Email和Password三个字段，将查找的其余字段(Id、Uuid、CreatedAt: 均为动态赋值，存到users[0])
    sqlStr2 := "select uuid, created_at from users where id = ?"
    row := Db.QueryRow(sqlStr2, id)
    row.Scan(&user.Uuid, &user.CreatedAt)
    fmt.Println(user)

    return
}

// Get a single user given the email
func UserByEmail(email string) (user User, err error) {
    user = User{}
    err = Db.QueryRow("SELECT id, uuid, name, email, password, photo, created_at FROM users WHERE email = ?", email).
    Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.Photo, &user.CreatedAt)
    return
}

// test: get the first user
func GetFirstUser() (user User, err error) {
    user = User{}
    err = Db.QueryRow("SELECT id, uuid, name, email, password, photo, created_at FROM users LIMIT 0, 1").Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.Photo, &user.CreatedAt)
    return
}
