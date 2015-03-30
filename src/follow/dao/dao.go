package dao

import (
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
    "follow/config"
)

func Exec(q string, params []interface{}) (sql.Result, error) {
    db, err := sql.Open("sqlite3", config.Get("dbf"))
    if err != nil {
        return nil, err
    }
    defer db.Close()

    stmt, err := db.Prepare(q)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    return stmt.Exec(params...)
}

func Query(q string, params []interface{}, callback func(*sql.Rows) error) error {
    db, err := sql.Open("sqlite3", config.Get("dbf"))
    if err != nil {
        return err
    }
    defer db.Close()

    stmt, err := db.Prepare(q)
    if err != nil {
        return err
    }
    defer stmt.Close()

    rows, err := stmt.Query(params...)
    if err != nil {
        return err
    }
    defer rows.Close()

    for rows.Next() {
        err = callback(rows)
        if err != nil {
            return err
        }
    }    

    return nil
}

func Row(q string, params []interface{}, bind []interface{}) error {
    db, err := sql.Open("sqlite3", config.Get("dbf"))
    if err != nil {
        return err
    }
    defer db.Close()

    stmt, err := db.Prepare(q)
    if err != nil {
        return err
    }
    defer stmt.Close()

    row := stmt.QueryRow(params...)
    if err != nil {
        return err
    }

    err = row.Scan(bind...)
    if err != nil {
        return err
    }

    return nil
}

