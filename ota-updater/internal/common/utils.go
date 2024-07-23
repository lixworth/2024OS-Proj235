package common

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ExecCommand(command string) (string, error) {
	output, err := exec.Command(command).CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

type ResponseTemplate struct {
	Error int `json:"error"`
	Data  any `json:"data"`
}

func Response(error int, data interface{}) ResponseTemplate {
	return ResponseTemplate{
		Error: error,
		Data:  data,
	}
}
func IsRelease() bool {
	arg1 := strings.ToLower(os.Args[0])
	name := filepath.Base(arg1)

	return strings.Index(name, "__") != 0 && strings.Index(arg1, "go-build") < 0
}
