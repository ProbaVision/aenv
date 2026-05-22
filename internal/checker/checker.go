package checker

import (
	"os"
	"path/filepath"
)

type Entry struct {
	Path  string `yaml:"path"`
	Agent string `yaml:"agent"`
}

type candidate struct {
	relativePath string
	agent        string
}

var candidates = []candidate{
	{relativePath: ".claude", agent: "claude code"},
	{relativePath: ".copilot", agent: "copilot"},
	{relativePath: ".agent", agent: "open source"},
	{relativePath: "claude.md", agent: "claude code"},
	{relativePath: "agents.md", agent: "open source"},
}

func ScanHome(home string) ([]Entry, error) {
	entries := make([]Entry, 0, len(candidates))

	for _, item := range candidates {
		path := filepath.Join(home, item.relativePath)
		if _, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}

		entries = append(entries, Entry{
			Path:  path,
			Agent: item.agent,
		})
	}

	return entries, nil
}
