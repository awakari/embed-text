package grpc

import (
	"context"
	"github.com/awakari/embed-text/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type controller struct {
	svc service.Service
}

func NewController(svc service.Service) ServiceServer {
	return controller{
		svc: svc,
	}
}

func (c controller) EmbedTexts(ctx context.Context, req *EmbedTextsRequest) (resp *EmbedTextsResponse, err error) {
	resp = &EmbedTextsResponse{}
	var embeddings [][]float32
	embeddings, err = c.svc.EmbedTexts(ctx, req.Texts)
	switch {
	case err == nil:
		for _, embedding := range embeddings {
			resp.Embeddings = append(resp.Embeddings, &Embedding{
				Features: embedding,
			})
		}
	default:
		err = status.Error(codes.Internal, err.Error())
	}
	return
}
