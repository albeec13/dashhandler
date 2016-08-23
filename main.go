package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "fmt"
    "strings"
    "log"
    "io/ioutil"
    "encoding/json"
)

type ConfigFile struct {
    DBUser  string
    DBPass  string
    DBName  string
}

type DHCPEvent struct {
    Event    string
    MAC      string
    IP       string
    Hostname string
}

func main() {
    var dbh DBHelper
    config, err := readConfigFile()

    if err == nil {
        err := dbh.Open(config.DBUser, config.DBPass, config.DBName)

        if err == nil {
            tables, err := dbh.GetTables()
            if err == nil {
                for _, table := range tables {
                    fmt.Printf("Table: %s\n", table)
                }
            } else {
                fmt.Printf("DATABASE ERROR: %s\n", err)
            }
        } else {
            log.Fatalf("FATAL ERROR: %s\n", err)
        }
    } else {
        log.Fatalf("FATAL ERROR: %s\n", err)
    }

    // Configure routes
    router := gin.Default()
    router.GET("/handle", func(c *gin.Context) {
        dhcpevent := DHCPEvent {
            c.DefaultQuery("event", "null"),
            strings.ToLower(c.DefaultQuery("mac", "null")),
            c.DefaultQuery("ip", "null"),
            c.DefaultQuery("hostname", "null"),
        }
        c.String(http.StatusOK, dhcpevent.Event)
        fmt.Println(dhcpevent)
    })

    // Run server
    router.Run(":4469")
}

func readConfigFile() (*ConfigFile, error) {
    config := &ConfigFile{}

    file, err  := ioutil.ReadFile("dashhandler.conf")
    if file != nil {
        err = json.Unmarshal(file, config)
    }
    return config, err
}
