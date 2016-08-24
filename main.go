package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "fmt"
    "strings"
    "log"
)

type DHCPEvent struct {
    Event    string
    MAC      string
    IP       string
    Hostname string
}

func main() {
    var dbh DBHelper
    var config ConfigFile
    path := "dashhandler.conf"
    err := config.ReadConfigFile(path)

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
