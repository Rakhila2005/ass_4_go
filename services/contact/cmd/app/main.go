package main

import (
    "fmt"
    "day2/pkg/store/postgres"
)

func main() {
    host := "localhost"
    port := "5432"
    user := "postgres"
    password := "1908"
    dbname := "ass_3_go"

    db, err := postgres.ConnectDB(host, port, user, password, dbname)
    if err != nil {
        fmt.Println("Ошибка подключения к базе данных:", err)
        return
    }
    defer db.Close()
}
