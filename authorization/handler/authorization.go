package handler

import (
	"context"
	"strings"

	log "github.com/micro/micro/v3/service/logger"

	authorization "authorization/proto"
)

// AuthorizationHandler is the handler containing the methods for the service endpoint
type AuthorizationHandler struct{}

var _ authorization.AuthorizationHandler = (*AuthorizationHandler)(nil)

// New returns a new AuthorizationHandler
func New() *AuthorizationHandler {
	return &AuthorizationHandler{}
}

// Call is a single request handler called via client.Call or the generated client code
func (h *AuthorizationHandler) Call(_ context.Context, req *authorization.Request, rsp *authorization.Response) error {
	log.Info("Received AuthorizationHandler.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (h *AuthorizationHandler) Stream(_ context.Context, req *authorization.StreamingRequest, stream authorization.Authorization_StreamStream) error {
	log.Infof("Received AuthorizationHandler.Stream request with count: %d", req.Count)

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
func (h *AuthorizationHandler) PingPong(_ context.Context, stream authorization.Authorization_PingPongStream) error {
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

func (h *AuthorizationHandler) Health(_ context.Context, req *authorization.HealthRequest, rsp *authorization.HealthResponse) error {
	rsp.CanReachMicroservice = true
	rsp.AccessTokenIsValid = validateAccessToken(req.GetAccessToken())

	return nil
}

/*
validateAccessToken check if the supplied token is valid & not outdated

at the moment, it only checks whether the given token follow the structure of a JWT `${Header}.${Payload}.${Signature}`

TODO: implement proper validation as below
https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-verifying-a-jwt.html
*/
func validateAccessToken(accessToken string) bool {
	return len(strings.Split(accessToken, ".")) == 3
}
