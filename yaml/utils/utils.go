package utils

import (
	"gopkg.in/yaml.v3"
	"fmt"
)

func SerializeToYaml(v any)(out []byte, err error) {
	out, err = yaml.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling to YAML: %w",err)
	}
	return out, nil
	}