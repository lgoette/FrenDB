package db

import (
    "database/sql"
    "fmt"
    _ "github.com/glebarez/go-sqlite"
    _ "github.com/google/uuid"
    "log"
    _ "log"
)

func OpenDB() *sql.DB {
    db, err := sql.Open("sqlite", "./local.db")
    if err != nil {
        log.Fatal(err)
        return nil
    }
    return db
}

func InitDB(db *sql.DB) {
    err := createTables(db)
    if err != nil {
        log.Fatal(err)
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
            creator_id text not null,
            name text not null,
            type text not null,
            details text,
            primary key (id, creator_id),
            foreign key (creator_id)
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
            creator_id text not null,
            entity_a text not null,
            entity_b text not null,
            connecting_entity text,
            start_date text,
            end_date text,
            details text,
            primary key (id, creator_id)
        );`

    _, err = db.Exec(createConnection)
    if err != nil {
        return err
    }

    createDuplicates := `create table if not exists duplicates (
            id text not null,
            creator_id text not null,
            duplicate_ids text not null,
            primary key (id, creator_id)
        );`

    _, err = db.Exec(createDuplicates)
    if err != nil {
        return err
    }

    return nil
}


func CloseDB(db *sql.DB) {
    err := db.Close()
    if err != nil {
        fmt.Println(err)
    }
}