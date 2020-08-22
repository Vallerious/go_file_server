package socket

import (
    "net"
    "github.com/Vallerious/http_server/stringutils"
    "github.com/Vallerious/http_server/connection"
)

func Listen(protocol string, port uint16) error {
    ln, err := net.Listen(protocol, stringutils.ConstructPortString(port))
    
    if err != nil {
        return err
    }

    for {
        conn, err := ln.Accept()

        if err != nil {
            return err
        }

        go connection.HandleConnection(conn)
    }
}

