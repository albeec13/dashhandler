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
    config, err := readConfigFile()
    router := gin.Default()

    if err == nil {
        var dbh DBHelper
        err := dbh.Open(config.DBUser, config.DBPass, config.DBName)

        if err == nil {
            rows, err := dbh.Query("SHOW TABLES")
            if err == nil {
                for rows.Next() {
                    var name string
                    if err := rows.Scan(&name); err == nil {
                        fmt.Printf("Table: %s\n", name)
                    } else {
                        fmt.Printf("Row.Scan() error: %s", err)
                    }
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
