package service

import (
	"context"
	"fmt"
	"github.com/awakari/embed-text/util"
	"log/slog"
)

type logging struct {
	svc Service
	log *slog.Logger
}

func NewLogging(svc Service, log *slog.Logger) Service {
	return logging{
		svc: svc,
		log: log,
	}
}

func (l logging) EmbedTexts(ctx context.Context, texts []string) (embeddings [][]float32, err error) {
	embeddings, err = l.svc.EmbedTexts(ctx, texts)
	ll := util.LogLevel(err)
	l.log.Log(ctx, ll, fmt.Sprintf("EmbedTexts(%d): %s", len(texts), err))
	return
}
