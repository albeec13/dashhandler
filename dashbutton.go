package main

import (
    "fmt"
)

type DashButton struct {
    MAC string
    Name string
}

func ProcessDashButton(config *ConfigFile, MAC string) {
    for _, dash := range config.DashButtons {
        if dash.MAC == MAC {
            fmt.Printf("Found %s Dash Button (%s)!\n", dash.Name, MAC)
        }
    }
}

