package main

import (
    "fmt"
    "log"
    "net/http"
    "io"
)

func main() {
    http.HandleFunc("/get_from_second", getDataFromSecondService) 
    http.HandleFunc("/message", message) 
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func getDataFromSecondService(w http.ResponseWriter, r *http.Request){
    resp, err := http.Get("http://localhost:8081/message")
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
    w.Write([]byte("hello from first service"))
}

