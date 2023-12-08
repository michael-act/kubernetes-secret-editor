package common

import (
	"gopkg.in/yaml.v3"
)

const (
	DataKey = "data"
)

// Secret allows us to read and return the full Kubernetes secret
type Secret map[string]interface{}

// SecretData extracts out the data portion of a Kubernetes secret
type SecretData struct {
	Data map[string]string `json:"data" yaml:"data"`
}

func GetStringSecret(s *Secret) (string, error) {
	secret, err := yaml.Marshal(s)
	if err != nil {
		return "", err
	}
	
	return string(secret), nil
}

func GetFullSecret(output []byte, sd *SecretData) (*Secret, error) {
	var secret Secret
	var err error

	err = yaml.Unmarshal(output, &secret)
	if err != nil {
		return nil, err
	}

	secret[DataKey] = sd.Data
	return &secret, nil
}
