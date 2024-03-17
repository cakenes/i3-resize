package main

import (
	"encoding/json"
	"fmt"
	// "log"
	"os"
	"os/exec"
)

type Msg struct {
	Success bool `json:"success"`
}

func main() {
	args := os.Args[1:5]
	cmd := exec.Command("i3-msg", "resize", "grow", args[0], args[2], args[3])
	out, _ := cmd.Output()

	var msg []Msg

	err := json.Unmarshal(out, &msg)
	if err != nil {
		fmt.Println("error:", err)
	}

	if !msg[0].Success {
		cmd := exec.Command("i3-msg", "resize", "shrink", args[1], args[2], args[3])
		cmd.Run()
	}
}
