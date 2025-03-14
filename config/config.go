package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Api struct {
		Port uint16 `envconfig:"API_PORT" default:"50051" required:"true"`
	}
	Log struct {
		Level int `envconfig:"LOG_LEVEL" default:"-4" required:"true"`
	}
	Model struct {
		Path     string `envconfig:"MODEL_PATH" default:"/model" required:"true"`
		Name     string `envconfig:"MODEL_NAME" default:"intfloat/multilingual-e5-small"`
		FileName string `envconfig:"MODEL_FILE_NAME" default:"model.onnx" required:"true"`
		Chunk    ChunkConfig
	}
}

type ChunkConfig struct {
	Size    uint32 `envconfig:"MODEL_CHUNK_SIZE" default:"1024" required:"true"`
	Overlap uint32 `envconfig:"MODEL_CHUNK_OVERLAP" default:"128" required:"true"`
}

func NewConfigFromEnv() (cfg Config, err error) {
	err = envconfig.Process("", &cfg)
	return
}
