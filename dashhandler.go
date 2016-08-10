package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type DHCPEvent struct {
    Event    string
    MAC      string
    IP       string
    Hostname string
}

func main() {
    router := gin.Default()

    router.GET("/:event/:mac/:ip/*hostname", func(c *gin.Context) {
        dhcpevent := DHCPEvent{
            c.Param("event"),
            c.Param("mac"),
            c.Param("ip"),
            c.Param("hostname"),
        }
        c.String(http.StatusOK, dhcpevent.Event)
    })

    router.Run(":4469")
}

