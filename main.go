package main

import "github.com/Vallerious/http_server/socket"

func main() {
    socket.Listen("tcp", 8080)    
}

