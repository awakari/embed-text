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

func (c controller) EmbedText(ctx context.Context, req *EmbedTextRequest) (resp *EmbedTextResponse, err error) {
	resp = &EmbedTextResponse{}
	var embeddings [][]float32
	embeddings, err = c.svc.EmbedText(ctx, req.Prefix, req.Text)
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
