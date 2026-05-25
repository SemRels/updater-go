# updater-go

Go updater plugin for Semantic Release.

Updates Go module metadata and versions during Semantic Release.

## Documentation

- Docs (coming soon): <https://github.com/SemRels/semrel/tree/main/docs/plugins/updater-go>
- Template source: <https://github.com/SemRels/plugin-template>

## Repository Layout

`	ext
cmd/plugin/              Plugin entry point
internal/plugin/         Business logic scaffold
internal/grpc/           gRPC transport scaffold
proto/v1                 Symlink to the SemRel protobuf contract
.github/workflows/       CI, release, and security automation
`

## Development

`ash
go build ./cmd/plugin
go test ./...
`

## Configuration Example

`yaml
plugins:
  - name: updater-go
    type: updater
    config:
      go_mod_file: go.mod
      version_package: internal/version
      update_ldflags: true
`

## Status

This repository is bootstrapped from SemRels/plugin-template and is ready for implementation.