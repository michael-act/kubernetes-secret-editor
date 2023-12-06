//go:build wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/michaelact/kubernetes-secret-editor/wasm/decoder"
)

func main() {
	wait := make(chan struct{}, 0)

	// Set up a callback function that will be called from JavaScript
	js.Global().Set("ProcessEncodedSecret", js.FuncOf(ProcessEncodedSecret))

	// Keep the program running
	<-wait
}

func ProcessEncodedSecret(this js.Value, args []js.Value) interface{} {
	input := []byte(args[0].String())

	secretData, err := decoder.GetSecretData(input)
	if err != nil {
		errorStr := fmt.Sprintf("Unable to get secret data. Error %s occurred\n", err)
		return map[string]interface{} {
			"errorInput": errorStr,
		}
	}

	fullSecret, err := decoder.GetFullSecret(input, secretData)
	if err != nil {
		errorStr := fmt.Sprintf("Unable to get full secret file. Error %s occurred\n", err)
		return map[string]interface{} {
			"errorInput": errorStr,
		}
	}

	output, err := decoder.GetStringSecret(fullSecret)
	if err != nil {
		errorStr := fmt.Sprintf("Unable to convert the secret to string. Error %s occurred\n", err)
		return map[string]interface{} {
			"errorInput": errorStr,
		}
	}

	return map[string]interface{} {
		"output": output,
	}
}
