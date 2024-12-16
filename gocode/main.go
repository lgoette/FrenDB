package main

import (
    _ "github.com/glebarez/go-sqlite"
)


func main() {
    db := openDB()
    defer closeDB(db)

    // create tables if necessary
    initDB(db)
    
    // run some tests
    testDb(db)
}