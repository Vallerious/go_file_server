package connection

import (
    "net"
    "io"
    "log"
)

func HandleConnection(c net.Conn) {
    io.Copy(c, c)
    log.Println("received connection")
    c.Close()
}


