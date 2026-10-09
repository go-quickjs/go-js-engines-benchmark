package engines

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

//go:embed node_worker.js
var nodeWorker string

type Node struct {
	Jitless bool
	cmd     *exec.Cmd
	stdin   io.WriteCloser
	encoder *json.Encoder
	decoder *json.Decoder
	stderr  bytes.Buffer
	printed int
}

type nodeResponse struct {
	Ready   bool       `json:"ready"`
	Jitless bool       `json:"jitless"`
	Output  [][]string `json:"output"`
	Error   string     `json:"error"`
}

func (n *Node) Name() string {
	if n.Jitless {
		return "NodeJitless"
	}
	return "Node"
}

func (n *Node) Init() error {
	if n.cmd != nil {
		return fmt.Errorf("%s is already initialized", n.Name())
	}
	flag := "--no-jitless"
	if n.Jitless {
		flag = "--jitless"
	}
	n.stderr.Reset()
	cmd := exec.Command("node", flag, "-e", nodeWorker)
	cmd.Stderr = &n.stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		stdin.Close()
		return err
	}
	if err := cmd.Start(); err != nil {
		stdin.Close()
		stdout.Close()
		return fmt.Errorf("could not start %s (node must be on PATH): %w", n.Name(), err)
	}
	n.cmd, n.stdin = cmd, stdin
	n.encoder, n.decoder = json.NewEncoder(stdin), json.NewDecoder(stdout)
	// Wait for the VM to be ready before benchmark timing starts.
	var ready nodeResponse
	if err := n.decoder.Decode(&ready); err != nil {
		return errors.Join(fmt.Errorf("%s startup: %w", n.Name(), err), n.Close())
	}
	if !ready.Ready || ready.Jitless != n.Jitless {
		return errors.Join(fmt.Errorf("%s worker started in the wrong mode", n.Name()), n.Close())
	}
	return nil
}

func (n *Node) Run(inputFile string) ([][]string, error) {
	source, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %w", inputFile, err)
	}
	return n.eval(string(source), inputFile)
}

func (n *Node) eval(source, filename string) ([][]string, error) {
	if n.cmd == nil {
		return nil, fmt.Errorf("%s is not initialized", n.Name())
	}
	request := struct {
		Source   string `json:"source"`
		Filename string `json:"filename"`
	}{source, filename}
	if err := n.encoder.Encode(request); err != nil {
		return nil, errors.Join(fmt.Errorf("%s request: %w", n.Name(), err), n.Close())
	}
	var response nodeResponse
	if err := n.decoder.Decode(&response); err != nil {
		return nil, errors.Join(fmt.Errorf("%s response: %w", n.Name(), err), n.Close())
	}
	for _, line := range response.Output[n.printed:] {
		fmt.Println(strings.Join(line, ""))
	}
	n.printed = len(response.Output)
	if response.Error != "" {
		return response.Output, fmt.Errorf("%s: %s", n.Name(), response.Error)
	}
	return response.Output, nil
}

func (n *Node) Close() error {
	if n.cmd == nil {
		return nil
	}
	n.stdin.Close()
	err := n.cmd.Wait()
	n.cmd, n.stdin, n.encoder, n.decoder = nil, nil, nil, nil
	n.printed = 0
	if err != nil {
		return fmt.Errorf("%s process: %w: %s", n.Name(), err, strings.TrimSpace(n.stderr.String()))
	}
	return nil
}

var _ JSEngine = (*Node)(nil)
