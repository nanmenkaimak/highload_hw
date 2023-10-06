package main

import (
    "io"
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/get_from_first", getDataFromFirstService) 
    http.HandleFunc("/message", message) 
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func getDataFromFirstService(w http.ResponseWriter, r *http.Request){
    resp, err := http.Get("http://localhost:8080/message")
    if err!=nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err!=nil {
        fmt.Println(err)
        return
    }
    w.Write(body)
}
func message(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("hello from second service"))
}

