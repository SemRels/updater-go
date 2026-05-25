// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

package main

import (
	"log"

	plugin "github.com/SemRels/updater-go/internal/plugin"
)

func main() {
	builder := plugin.NewBuilder(plugin.BuildConfig{})
	log.Printf("updater-go plugin ready: builds Go release archives (%T)", builder)
}
