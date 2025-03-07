package service

import (
	"context"
	"github.com/knights-analytics/hugot/pipelines"
)

type Service interface {
	EmbedTexts(ctx context.Context, texts []string) (embeddings [][]float32, err error)
}

type svc struct {
	pipeline *pipelines.FeatureExtractionPipeline
}

func New(pipeline *pipelines.FeatureExtractionPipeline) Service {
	return svc{
		pipeline: pipeline,
	}
}

func (s svc) EmbedTexts(ctx context.Context, texts []string) (embeddings [][]float32, err error) {
	var out *pipelines.FeatureExtractionOutput
	out, err = s.pipeline.RunPipeline(texts)
	if err == nil {
		embeddings = out.Embeddings
	}
	return
}
