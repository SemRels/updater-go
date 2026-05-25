// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

// Package plugin updates Go version files in-place.
package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var declarationTemplate = `(?m)^(?P<prefix>\s*(?:const|var)\s+%s\s*=\s*)"[^"]*"`

// Updater updates Go version declarations or simple version files.
type Updater struct{}

// NewUpdater creates an updater.
func NewUpdater() *Updater {
	return &Updater{}
}

// Update rewrites the configured version file.
func (u *Updater) Update(path, varName, version string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	if updated, ok := replaceDeclaration(data, varName, version); ok {
		return writeFile(path, updated)
	}

	if isPlainVersionFile(path, data) {
		return writeFile(path, []byte(version+"\n"))
	}

	return fmt.Errorf("version variable %q not found in %s", varName, path)
}

func replaceDeclaration(data []byte, varName, version string) ([]byte, bool) {
	re := regexp.MustCompile(fmt.Sprintf(declarationTemplate, regexp.QuoteMeta(varName)))
	if !re.Match(data) {
		return nil, false
	}

	updated := re.ReplaceAllString(string(data), `${prefix}"`+version+`"`)
	return []byte(updated), true
}

func isPlainVersionFile(path string, data []byte) bool {
	base := strings.ToLower(filepath.Base(path))
	if base != "version" {
		return false
	}

	content := strings.TrimSpace(string(data))
	return content != "" && !strings.Contains(content, "\n")
}

func writeFile(path string, data []byte) error {
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}
