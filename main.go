package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func pingDevice(ip string, wg *sync.WaitGroup){
    defer wg.Done()

    url := fmt.Sprintf("http://%s/ping", ip)

    client := http.Client{
        Timeout: 2 * time.Second,
    }

    resp, err := client.Get(url)

    if err != nil{
        return
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil{
        fmt.Println(err)
    }

    if strings.ToLower(string(body)) == "pong" {
        fmt.Println(ip)
    }
}

func main(){
    var wg sync.WaitGroup
    var subnet string

    if len(os.Args) >= 2{
        subnet = os.Args[1]
    }

    for i := 1; i < 255; i++ {
        ip := fmt.Sprintf("%s%d", subnet, i)
        wg.Add(1)
        go pingDevice(ip, &wg)
    }
    wg.Wait()
}
