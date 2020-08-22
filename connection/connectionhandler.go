package connection

import (
    "net"
    "log"
    "bufio"
    "net/http"
)

func HandleConnection(c net.Conn) error {
    log.Println("received connection")
    reader := bufio.NewReader(c)
    req, err := http.ReadRequest(reader)
    
    if err != nil {
        return err
    }

    log.Printf("Incoming request on path: %s and method: %s", req.URL, req.Method)
    
    c.Close()

    return nil
}


