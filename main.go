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
    
    if err := config.ReadConfigFile(path); err == nil {
        if err = dbh.Open(&config); err == nil {
            var tables []string
            if tables, err = dbh.GetTables(); err == nil {
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
        if result, err := dbh.LogDevice(&dhcpevent); err != nil {
            fmt.Printf("LOG DHCPEVENT ERROR: %s, result:", err)
            fmt.Println(result)
        }
    })

    // Run server
    router.Run(":4469")
}
