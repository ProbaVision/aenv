package toolcheck

import (
	"os/exec"
	"regexp"
	"strings"
)

type Tool struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

var toolNames = []string{"jq", "yq", "timeout", "grep", "sed", "awk"}

var versionRe = regexp.MustCompile(`\d+\.\d+(\.\d+)?`)

func CheckTools() ([]Tool, error) {
	tools := make([]Tool, 0, len(toolNames))
	for _, name := range toolNames {
		t := Tool{Name: name, Version: "Missing"}
		if path, err := exec.LookPath(name); err == nil {
			if v := getVersion(path); v != "" {
				t.Version = v
			}
		}
		tools = append(tools, t)
	}
	return tools, nil
}

func getVersion(path string) string {
	out, err := exec.Command(path, "--version").Output()
	if err != nil {
		return ""
	}
	firstLine := strings.SplitN(string(out), "\n", 2)[0]
	return versionRe.FindString(firstLine)
}
