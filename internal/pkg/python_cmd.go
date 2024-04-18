package pkg

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
)

func StreamPythonScriptOutput(wg *sync.WaitGroup, scriptPath ...string) {
	defer wg.Done()

	cmd := exec.Command("python3", scriptPath...)
	fmt.Println(cmd.String())
	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()
	multiReader := io.MultiReader(stdoutPipe, stderrPipe)

	err := cmd.Start()
	if err != nil {
		fmt.Println("run python failed, error:", err)
		return
	}

	var wgOutput sync.WaitGroup
	wgOutput.Add(1)
	go func() {
		defer wgOutput.Done()
		reader := bufio.NewReader(multiReader)
		for {
			line, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				break
			}
			fmt.Print(line)
		}
	}()

	err = cmd.Wait()
	if err != nil {
		fmt.Println("wait python finished failed, error:", err)
	}

	wgOutput.Wait()
}

func RunPythonScript(scriptPath ...string) (string, error) {
	cmd := exec.Command("python3", scriptPath...)
	fmt.Println(cmd.String())
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	result := strings.TrimSpace(string(output))
	return result, nil
}
