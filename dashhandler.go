package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "fmt"
)

type DHCPEvent struct {
    Event    string
    MAC      string
    IP       string
    Hostname string
}

func main() {
    router := gin.Default()

    router.GET("/handle", func(c *gin.Context) {
        dhcpevent := DHCPEvent{
            c.DefaultQuery("event", "null"),
            c.DefaultQuery("mac", "null"),
            c.DefaultQuery("ip", "null"),
            c.DefaultQuery("hostname", "null"),
        }
        c.String(http.StatusOK, dhcpevent.Event)
        fmt.Printf("%+v\n",dhcpevent)
    })

    router.Run(":4469")
}

