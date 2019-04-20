package main

import (
    "github.com/DGHeroin/yig"
    "log"
)

func sendPing() {
    ping := yig.NewClient("pong", "ping")
    if result, err := ping.Request([]byte("ping")); err != nil {
        log.Printf("error: %v", err)
    } else {
        log.Printf("result: %s", result)
    }
}

func testQPSPing() {
    ping := yig.NewClient("pong", "ping")

    for {
        if _, err := ping.Request([]byte("ping")); err != nil {
            log.Println(err)
            break
        }
    }
}

var (
    qps = true
)

func main()  {
    if qps {
       testQPSPing()
    } else {
        sendPing()
    }

}


