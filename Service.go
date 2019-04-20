package yig

import (
    "context"
    "fmt"
    proto "github.com/DGHeroin/yig/proto"
    "github.com/micro/go-micro"
    _ "github.com/micro/go-plugins/registry/etcd"
)




var (
    errorNotImpl = fmt.Errorf("function OnRequest() not found")
)

type HandleFunc func(request []byte)(response []byte, err error)

type Service interface {
    Run() error
    Request([]byte) ([]byte, error)
}

type service struct {
    proto.Service
    service micro.Service
    client  proto.Service
    OnHandleFunc HandleFunc
}

func NewService(serviceName string, callback HandleFunc) Service {
    s := &service{}
    s.OnHandleFunc = callback
    s.service = micro.NewService(micro.Name(serviceName))
    s.service.Init()
    proto.RegisterServiceHandler(s.service.Server(), s)
    return s
}

func NewClient(serviceName string, clientName string) Service {
    s := &service{}
    s.service = micro.NewService(micro.Name(clientName))
    s.service.Init()
    s.client = proto.NewService(serviceName, s.service.Client()) // remote service name
    return s
}

func (s *service) DoRequest(ctx context.Context, req *proto.Request, res *proto.Response) (err error) {
    if s.OnHandleFunc == nil {
        return errorNotImpl
    }
    res.Body, err = s.OnHandleFunc(req.Body)
    return
}

func (s *service) Run() error {
    return s.service.Run()
}

func (s *service) Request(data []byte) (result []byte, err error) {
    var (
        req proto.Request
        resp *proto.Response
    )
    req.Body = data
    resp, err = s.client.DoRequest(context.TODO(), &req)
    if resp != nil {
        result = resp.Body
    }

    return
}
