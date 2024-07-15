package common

import "os/exec"

func ExecCommand(command string) (string, error) {
	output, err := exec.Command(command).CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

type ResponseTemplate struct {
	Error int `json:"error"`
	Data  any
}

func Response(error int, data interface{}) ResponseTemplate {
	return ResponseTemplate{
		Error: error,
		Data:  data,
	}
}
