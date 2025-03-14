package service

import (
	"context"
	"github.com/awakari/embed-text/config"
	"github.com/awakari/embed-text/util"
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
	chunks := util.TextSplitWithOverlaps(text, int(s.cfgChunk.Size), int(s.cfgChunk.Overlap))
	var prefixedChunks []string
	for _, chunk := range chunks {
		prefixedChunks = append(prefixedChunks, prefix+chunk)
	}
	var out *pipelines.FeatureExtractionOutput
	out, err = s.pipeline.RunPipeline(prefixedChunks)
	if err == nil {
		embeddings = out.Embeddings
	}
	return
}
