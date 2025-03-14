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

func (l logging) EmbedText(ctx context.Context, prefix, text string) (embeddings [][]float32, err error) {
	embeddings, err = l.svc.EmbedText(ctx, prefix, text)
	ll := util.LogLevel(err)
	l.log.Log(ctx, ll, fmt.Sprintf("EmbedText(prefix=%s, bytes=%d): %d, %s", prefix, len(text), len(embeddings), err))
	return
}
