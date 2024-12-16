package test

import (
    "database/sql"
    "fmt"
    "frenDB/data"
    _ "github.com/glebarez/go-sqlite"
    "github.com/google/uuid"
    "log"
    _ "log"
)

func TestDb(db *sql.DB)  {
	_, err := insertTestData(db)
    if err != nil {
        log.Fatal(err)
        return
    }

    _, err = verifyTestData(db)
    if err != nil {
        log.Fatal(err)
        return
    }

    _, err = deleteTestData(db)
    if err != nil {
        log.Fatal(err)
        return
    }

    _, err = verifyTestData(db)
    if err != nil {
        log.Fatal(err)
        return
    }
}


func insertTestData(db *sql.DB) (int64, error) {
    fmt.Println("inserting user with name = testName into users")
    u := data.User{}
    u.Id = uuid.New()
    u.Name = "testName"
    insertUser :=  `insert into user (id, name) values (?, ?);`
    res, err := db.Exec(insertUser, u.Id, u.Name)
    if err != nil {
        fmt.Println("no user inserted")
        return 0, err
    }

    num, err := res.LastInsertId()
    fmt.Printf("%d users inserted\n", num)
    return num, err
}

func verifyTestData(db *sql.DB) (*data.User, error) {
    fmt.Println("selecting user with name = testName from users, result:")
    selectUser := "select * from user where name = ?;"

    row := db.QueryRow(selectUser, "testName")
    u := &data.User{}
    err := row.Scan(&u.Id, &u.Name)
    if err != nil {
        fmt.Println("no user found")
        return nil, nil
    }
    fmt.Printf("user found, name = %s, id = %s\n", u.Name, u.Id)
    return u, nil
}


func deleteTestData(db *sql.DB) (int64, error) {
    fmt.Println("deleting user with name = testName from users, result:")
    deleteUser := `delete from user where name = ?;`
    result, err := db.Exec(deleteUser, "testName")
    if err != nil {
        fmt.Println("no user deleted")
        return 0, err
    }
    num, err := result.RowsAffected()
    fmt.Printf("%d users deleted\n", num)
    return num, err
}
