package main

import (
    "database/sql"
    "fmt"
    _ "github.com/glebarez/go-sqlite"
    "log"
    _ "log"
)

func initDB() {
    db, err := sql.Open("sqlite", "./local.db")
    if err != nil {
        log.Fatal(err)
        return
    }


    defer closeDB(db)

    err = createTables(db)
    if err != nil {
        log.Fatal(err)
        return
    }
}

func createTables(db *sql.DB) error {
    createUser := `create table if not exists user (
            id text primary key,
            name text not null
        );`

    _, err := db.Exec(createUser)
    if err != nil {
        return err
    }

    createEntityType := `create table if not exists entity_type (
            name text primary key,
            description text
        );`

    _, err = db.Exec(createEntityType)
    if err != nil {
        return err
    }

    createEntity := `create table if not exists entity (
            id text not null,
            created_by text not null,
            name text not null unique,
            type text not null,
            details text,
            primary key (id, created_by),
            foreign key (created_by)
                references user (id),
            foreign key (type)
                references entity_type (name)
        );`

    _, err = db.Exec(createEntity)
    if err != nil {
        return err
    }

    createConnection := `create table if not exists connection (
            id text not null,
            created_by text not null,
            from text not null,
            to text not null,
            by text,
            start text,
            end text,
            details text,
            primary key (id, created_by)
        );`

    _, err = db.Exec(createConnection)
    if err != nil {
        return err
    }

    createDuplicates := `create table if not exists duplicates (
            id text primary key,
            created_by text primary key,
            duplicate_ids text not null
        );`

    _, err = db.Exec(createDuplicates)
    if err != nil {
        return err
    }

    return nil
}

func insertTestData() {
    insertUser := "insert into user (id, name) values ()"
}

func verifyTestData() {

}

func closeDB(db *sql.DB) {
    err := db.Close()
    if err != nil {
        fmt.Println(err)
    }
}