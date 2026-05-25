package plugin

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUpdaterUpdateGoFile(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "updater-go-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file := filepath.Join(dir, "version.go")
	original := "package version\n\nconst Version = \"1.2.3\"\n"
	if err := os.WriteFile(file, []byte(original), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := NewUpdater().Update(file, "Version", "1.3.0"); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	got, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), `const Version = "1.3.0"`) {
		t.Fatalf("updated file = %s", got)
	}
}

func TestUpdaterUpdatePlainVersionFile(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "updater-go-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file := filepath.Join(dir, "version")
	if err := os.WriteFile(file, []byte("1.2.3\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := NewUpdater().Update(file, "Version", "2.0.0"); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	got, err := os.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "2.0.0\n" {
		t.Fatalf("updated file = %q", string(got))
	}
}

func TestUpdaterMissingFile(t *testing.T) {
	t.Parallel()

	err := NewUpdater().Update(filepath.Join(t.TempDir(), "missing.go"), "Version", "1.3.0")
	if err == nil || !strings.Contains(err.Error(), "read") {
		t.Fatalf("expected read error, got %v", err)
	}
}

func TestUpdaterVariableNotFound(t *testing.T) {
	t.Parallel()

	dir, err := os.MkdirTemp("", "updater-go-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	file := filepath.Join(dir, "version.go")
	if err := os.WriteFile(file, []byte("package version\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	err = NewUpdater().Update(file, "Version", "1.3.0")
	if err == nil || !strings.Contains(err.Error(), `version variable "Version" not found`) {
		t.Fatalf("expected variable error, got %v", err)
	}
}
