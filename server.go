package main
import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "encoding/csv"
    "html/template"
    "log"
    "net/http"
)

type Color struct {
    BtnColor string
}

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
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatalln(err)
    }
    var color Color
    json.Unmarshal(b, &color)
    fmt.Println(string(b))
    fmt.Println(color)
    fmt.Println(color)
    resp:=make(map[string]string)
    resp["message"] = "Success"
    jsonResp, err := json.Marshal(resp)
    if err != nil {
        log.Fatalf("Error %s", err)
    }
    w.Write(jsonResp)
    return
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
