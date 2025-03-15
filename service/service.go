package service

import (
	"context"
	"github.com/awakari/embed-text/config"
	"github.com/knights-analytics/hugot/pipelines"
)

type Service interface {
	EmbedText(ctx context.Context, prefix, text string) (embeddings [][]float32, err error)
}

type svc struct {
	pipeline *pipelines.FeatureExtractionPipeline
	cfgChunk config.ChunkConfig
}

func New(pipeline *pipelines.FeatureExtractionPipeline, cfgChunk config.ChunkConfig) Service {
	return svc{
		pipeline: pipeline,
		cfgChunk: cfgChunk,
	}
}

func (s svc) EmbedText(ctx context.Context, prefix, text string) (embeddings [][]float32, err error) {
	snippet := prefix + text
	if len(text) > int(s.cfgChunk.Size) {
		snippet = snippet[:s.cfgChunk.Size]
	}
	var out *pipelines.FeatureExtractionOutput
	for {
		out, err = s.pipeline.RunPipeline([]string{
			snippet,
		})
		if err == nil {
			break
		}
		if len(snippet) < int(s.cfgChunk.Overlap) {
			break
		}
		snippet = snippet[:len(snippet)-int(s.cfgChunk.Overlap)]
	}
	if err == nil {
		embeddings = out.Embeddings
	}
	return
}
