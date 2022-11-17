package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	authorization "authorization/proto"
)

type Authorization struct{}

// Return a new handler
func New() *Authorization {
	return &Authorization{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Authorization) Call(ctx context.Context, req *authorization.Request, rsp *authorization.Response) error {
	log.Info("Received Authorization.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Authorization) Stream(ctx context.Context, req *authorization.StreamingRequest, stream authorization.Authorization_StreamStream) error {
	log.Infof("Received Authorization.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&authorization.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Authorization) PingPong(ctx context.Context, stream authorization.Authorization_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&authorization.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
