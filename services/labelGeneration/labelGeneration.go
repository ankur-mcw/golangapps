package main

import (
    "os"
    "github.com/ankur-mcw/golangapps/services/labelGeneration/server"
    "github.com/narvar/NarvarGolangApps/nlog"
)

func main()  {
    nlog.InitStdoutLogging("DEBUG")
    appPort := 7070
    err := server.Start(appPort)
    if err != nil {
        nlog.Errorf("Failed fasthttp.ListenAndServe at port %d: %s.", appPort, err)
    } else {
        nlog.Errorf("Failed fasthttp.ListenAndServe at port %d.", appPort)
    }
    os.Exit(1)
}
