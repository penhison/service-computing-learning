package main

import (
    "os"
    flag "github.com/spf13/pflag"
    "./service"
)

const (
    PORT string = "8080" //默认端口
)

func main() {
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = PORT
    }

    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }

    server := service.NewServer()
    server.Run(":" + port)
}
