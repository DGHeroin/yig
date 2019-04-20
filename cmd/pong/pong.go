package main

import (
    "github.com/DGHeroin/yig"
    "log"
    "sync/atomic"
    "time"
)

type Pong struct {
    yig.Service
}

var (
    qps = int64(0)
)

func (p *Pong) OnRequest(request []byte) (response []byte, err error) {
    response = []byte("pongongongongong")
    atomic.AddInt64(&qps, 1)
    return
}

func main()  {

    go func() {
        ticker := time.NewTicker(time.Second)
        for range ticker.C {
            if qps != 0 {
                log.Printf("qps:%v\n", qps)
            }

            atomic.StoreInt64(&qps, 0)
        }
    }()


    pong := &Pong{}
    yig.NewService("pong", pong.OnRequest).Run()
}

