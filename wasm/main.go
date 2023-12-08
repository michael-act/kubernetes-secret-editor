//go:build wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/michaelact/kubernetes-secret-editor/wasm/common"
	"github.com/michaelact/kubernetes-secret-editor/wasm/decoder"
	"github.com/michaelact/kubernetes-secret-editor/wasm/encoder"
)

func main() {
	wait := make(chan struct{}, 0)

	// Set up a callback function that will be called from JavaScript
	js.Global().Set("ProcessDecodeSecret", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		return ProcessSecret(decoder.GetSecretData, p)
	}))
	js.Global().Set("ProcessEncodeSecret", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		return ProcessSecret(encoder.GetSecretData, p)
	}))

	// Keep the program running
	<-wait
}

func ProcessSecret(getSecretData func([]byte) (*common.SecretData, error), p []js.Value) interface{} {
	input := []byte(p[0].String())

	secretData, err := getSecretData(input)
	if err != nil {
		errorStr := fmt.Sprintf("Unable to get secret data. Error %s occurred\n", err)
		return map[string]interface{}{
			"errorInput": errorStr,
		}
	}

	fullSecret, err := common.GetFullSecret(input, secretData)
	if err != nil {
		errorStr := fmt.Sprintf("Unable to get full secret file. Error %s occurred\n", err)
		return map[string]interface{}{
			"errorInput": errorStr,
		}
	}

	output, err := common.GetStringSecret(fullSecret)
	if err != nil {
		errorStr := fmt.Sprintf("Unable to convert the secret to string. Error %s occurred\n", err)
		return map[string]interface{}{
			"errorInput": errorStr,
		}
	}

	return map[string]interface{}{
		"output": output,
	}
}
