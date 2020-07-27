package main

import (
    "log"
    "./server"
)

func main()  { 
    err := server.Start("7070")
    if err != nil {
        log.Fatal(err)
    }
}
