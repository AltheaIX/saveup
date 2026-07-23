package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	switch cfg.Compression.Format {
	case "zip":
	default:
		return nil, fmt.Errorf(
			"unsupported compression format: %s\ncurrently supported compression are: zip",
			cfg.Compression.Format,
		)
	}

	return &cfg, nil
}
