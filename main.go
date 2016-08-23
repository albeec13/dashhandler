package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "fmt"
    "strings"
    "errors"
    "io/ioutil"
    "encoding/json"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type ConfigFile struct {
    DBUser  string
    DBPass  string
}

type DHCPEvent struct {
    Event    string
    MAC      string
    IP       string
    Hostname string
}

func main() {
    var config ConfigFile
    err := readConfigFile(&config)
    router := gin.Default()

    if err == nil {
        db, err :=  sql.Open("mysql", config.DBUser + ":" + config.DBPass + "@/dashhandler")

        if err == nil {
            rows, err := db.Query("SHOW TABLES")
            if err == nil {
                fmt.Println(rows)

                for rows.Next() {
                    var name string
                    if err := rows.Scan(&name); err == nil {
                        fmt.Printf("Table: %s\n", name)
                    } else {
                        fmt.Printf("Row.Scan() error: %s", err)
                    }
                }

            } else {
                fmt.Printf("DB Query error: %s\n", err)
            }
        } else {
            fmt.Printf("DB open failed with error: %s", err)
        }
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

func readConfigFile(config *ConfigFile) (err error) {
    file, err := ioutil.ReadFile("dashhandler.conf")
    if file != nil {
        return parseConfigFile(config, file)
    } else {
        return errors.New("Error: config file not found")
    }
}

func parseConfigFile(config *ConfigFile, file []byte) (err error) {
    if file != nil {
        if err := json.Unmarshal(file, config); err == nil {
            return nil
        } else {
            return errors.New("Error: invalid config JSON")
        }
    } else {
        return errors.New("Error: config file not found")
    }
}

