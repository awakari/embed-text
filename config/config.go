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
		Path           string `envconfig:"MODEL_PATH" default:"/model" required:"true"`
		Name           string `envconfig:"MODEL_NAME" default:"intfloat/multilingual-e5-small"`
		FileName       string `envconfig:"MODEL_FILE_NAME" default:"model.onnx" required:"true"`
		Chunk          ChunkConfig
		IntraOpThreads struct {
			Enabled bool `envconfig:"MODEL_INTRA_OP_THREADS_ENABLED" default:"true" required:"true"`
			Num     int  `envconfig:"MODEL_INTRA_OP_THREADS_NUM" default:"1" required:"true"`
		}
	}
}

type ChunkConfig struct {
	Size    uint32 `envconfig:"MODEL_CHUNK_SIZE" default:"1024" required:"true"`
	Overlap uint32 `envconfig:"MODEL_CHUNK_OVERLAP" default:"256" required:"true"`
}

func NewConfigFromEnv() (cfg Config, err error) {
	err = envconfig.Process("", &cfg)
	return
}
