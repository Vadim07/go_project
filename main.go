package main

import (
    "errors"

    "go_project/log" 
)

func main() {
    log.Info("123Some information")
    log.Warn("123Warn about something")
    log.Err(errors.New("123Some error"))
}
