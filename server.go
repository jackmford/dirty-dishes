package main

import (
    "fmt"
    "os"
    "encoding/csv"
    "html/template"
    "log"
    "net/http"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func hello(w http.ResponseWriter, r *http.Request) {
    file, err := os.Open("./static/save-id.txt")
    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    reader := csv.NewReader(file)
    records, _ := reader.ReadAll()
    fmt.Println(records[0][0])
    data := struct {
        Id string
        Text string
    }{
        Id: records[0][0],
        Text: records[1][0],
    }
    tpl.Execute(w, data)
}

func writeIdData(w http.ResponseWriter, r *http.Request) {
    fmt.Println("test")
}

func main() {
    http.HandleFunc("/", hello)
    http.HandleFunc("/write", writeIdData)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
    err := http.ListenAndServe(":8787", nil)
    if err != nil {
        log.Fatal(err)
    }
}
