package carbonize

import (
	"encoding/json"
	"fmt"
	"io"
)

// ParseConfig parses a file with a JSON configuration for Carbon.
//
// Ideally, the configuration file should contain a configuration exported
// directly from the Carbon website.
func ParseConfig(f io.Reader) (Config, error) {
	b, err := io.ReadAll(f)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read configuration file: %v", err)
	}

	var config Config
	err = json.Unmarshal(b, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal configuration file: %v", err)
	}

	return config, nil
}
