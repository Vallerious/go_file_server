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
    byteFileData, err := ioutil.ReadFile(path)

    if err != nil {
        return nil, errors.New("page not found")
    }
    
    return byteFileData, nil
}
