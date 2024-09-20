package handler

import "github.com/laix0n/satim/pkg/pb"

type Handler struct {
	pb.UnimplementedTestServiceServer
}

func (h Handler) SendTestResults(stream pb.TestService_SendTestResultsServer) error {
	return nil
}
