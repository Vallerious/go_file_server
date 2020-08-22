package handlers

import (
    "net/http"
    "errors"
    "io/ioutil"
    "os"
    "log"
    "path/filepath"
)


func FileHandler(req* http.Request) ([]byte, error) {
    if req.Method != "GET" {
        return nil, errors.New("File handling works only for GET requests")
    }

    wd, _ := os.Getwd()

    url := req.URL.String()
    path := filepath.Join(wd, "public", url)
    log.Println(path)
    byteFileData, _ := ioutil.ReadFile(path)
    
    return byteFileData, nil
}
