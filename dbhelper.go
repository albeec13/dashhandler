package main

import (
//    "strings"
//    "errors"
//    "encoding/json"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type DBHelper struct{
    db *sql.DB
}

func (dbh *DBHelper) Open(user string, pass string, dbname string) (err error) {
    dbh.db, err = sql.Open("mysql", user + ":" + pass + "@/" + dbname)
    return err
}

func (dbh *DBHelper) Query(query string) (rows *sql.Rows, err error) {
    rows, err = dbh.db.Query(query)
    return rows, err
}
