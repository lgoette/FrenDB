package main

import (
    "frenDB/db"
    "frenDB/test"
    _ "github.com/glebarez/go-sqlite"
    "os"
)


func main() {
    err := os.Remove("./local.db")
    if err != nil {
        return 
    }

    database := db.OpenDB("./local.db")
    // defer is cool
    defer db.CloseDB(database)

    // create tables if necessary
    db.InitDB(database)
    
    // run some tests
    test.TestDb(database)
}