package main_test

import (
    "os"
    "testing"

    "log"
    "rest-api"
)

var a main.App

func TestMain(m *testing.M) {
    a = main.App{}
    a.Initialize(
        os.Getenv("TEST_DB_USERNAME"),
        os.Getenv("TEST_DB_PASSWORD"),
        os.Getenv("TEST_DB_NAME"))

    ensureTableExists()

    code := m.Run()

    clearTable()

    os.Exit(code)
}

func ensureTableExists() {
    if _, err := a.DB.Exec(tableCreationQuery); err != nil {
        log.Fatal(err)
    }
}

func clearTable() {
    a.DB.Exec("DELETE FROM track")
    a.DB.Exec("ALTER SEQUENCE track_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS track 
(
id MEDIUMINT  NOT NULL AUTO_INCREMENT,
user MEDIUMINT,
client MEDIUMINT,
description TEXT,
PRIMARY KEY (id)
)`
