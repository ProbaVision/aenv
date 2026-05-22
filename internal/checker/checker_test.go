package checker

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanHomeReturnsKnownEntriesInSpecOrder(t *testing.T) {
	t.Parallel()

	home := t.TempDir()
	for _, name := range []string{".claude", ".copilot", "agents.md"} {
		path := filepath.Join(home, name)
		if filepath.Ext(name) == ".md" {
			if err := os.WriteFile(path, []byte("x"), 0o644); err != nil {
				t.Fatalf("write %s: %v", name, err)
			}
			continue
		}
		if err := os.Mkdir(path, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", name, err)
		}
	}

	got, err := ScanHome(home)
	if err != nil {
		t.Fatalf("ScanHome returned error: %v", err)
	}

	want := []Entry{
		{Path: filepath.Join(home, ".claude"), Agent: "claude code"},
		{Path: filepath.Join(home, ".copilot"), Agent: "copilot"},
		{Path: filepath.Join(home, "agents.md"), Agent: "open source"},
	}

	if len(got) != len(want) {
		t.Fatalf("len(got) = %d, want %d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d] = %+v, want %+v", i, got[i], want[i])
		}
	}
}
