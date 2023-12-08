package encoder

import (
	"encoding/base64"
	"gopkg.in/yaml.v3"

	"github.com/michaelact/kubernetes-secret-editor/wasm/common"
)

func GetSecretData(input []byte) (*common.SecretData, error) {
	var secretData common.SecretData
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

func parseSecretData(s *common.SecretData) error {
	var err error
	for key, value := range s.Data {
		s.Data[key], err = encodeString(value)
		if err != nil {
			return err
		}
	}

	return nil
}

func encodeString(decoded string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(decoded))
	return encoded, nil
}
