package main

import(
    "fmt"
    "log"
    "net/http"
    "os"
    "sync"
)

var wg sync.WaitGroup
var loc sync.Mutex


func sendRequest(url string){
    defer wg.Done()
    res, err := http.Get(url)
    if err != nil{
        panic(err)
    }
    loc.Lock()
    defer loc.Unlock()
    fmt.Printf("[%d] %s\n", res.StatusCode, url)
}


func main(){
    if len(os.Args) < 2{
        log.Fatalln("Usage: go run multipleUrl.go <url1> <url2> .. <urln>")
    }

    for _, url := range os.Args[1:] {
        go sendRequest("https://" + url)
        wg.Add(1)
    }
    wg.Wait()
}
