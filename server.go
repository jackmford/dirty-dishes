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
        Dishes_Id string
        Run_Id string
        Dishes_Text string
        Run_Text string
    }{
        Dishes_Id: records[0][0],
        Run_Id: records[1][0],
        Dishes_Text: records[2][0],
        Run_Text: records[3][0],
    }
    tpl.Execute(w, data)
}

func writeIdData(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")
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
