package connection

import (
    "net"
    "log"
    "bufio"
    "net/http"
    "github.com/Vallerious/http_server/handlers"
)

func HandleConnection(c net.Conn) error {
    log.Println("received connection")
    defer c.Close()
    reader := bufio.NewReader(c)
    req, err := http.ReadRequest(reader)
    
    if err != nil {
        return err
    }

    log.Printf("Incoming request on path: %s and method: %s\n", req.URL, req.Method)
    
    b, err := handlers.FileHandler(req)
    data := string(b)

    if err != nil {
        data = err.Error()
    }

    log.Println(data)

    writer := bufio.NewWriter(c)
    writer.WriteString(data)
    writer.Flush()

    return nil
}


