package main

import (
    "frenDB/db"
    "frenDB/test"
    _ "github.com/glebarez/go-sqlite"
)


func main() {
    database := db.OpenDB()
    // defer is cool
    defer db.CloseDB(database)

    // create tables if necessary
    db.InitDB(database)
    
    // run some tests
    test.TestDb(database)
}