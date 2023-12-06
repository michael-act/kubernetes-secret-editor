package decoder

import (
	"encoding/base64"
	"gopkg.in/yaml.v3"
)

const (
	dataKey = "data"
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

	secret[dataKey] = sd.Data
	return &secret, nil
}

func GetSecretData(input []byte) (*SecretData, error) {
	var secretData SecretData
	var err error

	err = yaml.Unmarshal(input, &secretData)
	if err != nil {
		return nil, err
	}

	err = parseSecretData(&secretData)
	if err != nil {
		return nil, err
	}

	return &secretData, nil
}

func parseSecretData(s *SecretData) error {
	var err error
	for key, value := range s.Data {
		s.Data[key], err = decodeString(value)
		if err != nil {
			return err
		}
	}

	return nil
}

func decodeString(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
