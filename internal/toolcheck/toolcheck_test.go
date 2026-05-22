package toolcheck

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckToolsInOrder(t *testing.T) {
	t.Parallel()

	tools, err := CheckTools()
	if err != nil {
		t.Fatalf("CheckTools returned error: %v", err)
	}

	if len(tools) != len(toolNames) {
		t.Fatalf("len(tools) = %d, want %d", len(tools), len(toolNames))
	}

	for i, name := range toolNames {
		if tools[i].Name != name {
			t.Fatalf("tools[%d].Name = %q, want %q", i, tools[i].Name, name)
		}
		if tools[i].Version == "" {
			t.Fatalf("tools[%d].Version is empty", i)
		}
	}
}

func TestGetVersion(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	script := filepath.Join(dir, "faketool")
	if err := os.WriteFile(script, []byte("#!/bin/sh\necho 'faketool version 3.2.1'\n"), 0o755); err != nil {
		t.Fatalf("write script: %v", err)
	}

	got := getVersion(script)
	if got != "3.2.1" {
		t.Fatalf("getVersion = %q, want %q", got, "3.2.1")
	}
}

func TestGetVersionNoMatch(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	script := filepath.Join(dir, "faketool")
	if err := os.WriteFile(script, []byte("#!/bin/sh\necho 'no version here'\n"), 0o755); err != nil {
		t.Fatalf("write script: %v", err)
	}

	got := getVersion(script)
	if got != "" {
		t.Fatalf("getVersion = %q, want empty", got)
	}
}
