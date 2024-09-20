package handler

import "yugioh-decks/pkg/pb"

type Handler struct {
	*pb.UnimplementedTestServiceServer
}
