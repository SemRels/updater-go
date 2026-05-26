# updater-go

Updates a Go source file that stores the application version.

This plugin is distributed as the standalone Go binary `semrel-plugin-updater-go`. Semrel executes the binary as a subprocess, provides plugin configuration through `SEMREL_PLUGIN_*` environment variables, provides release context through `SEMREL_*` environment variables, reads standard output, and treats exit code `0` as success and any non-zero exit code as failure. Install the binary in `~/.semrel/plugins/` or anywhere on your `$PATH`.

## Installation

```bash
go install github.com/SemRels/updater-go/cmd/plugin@latest
```

## Configuration

```yaml
plugins:
  - name: updater-go
    path: ~/.semrel/plugins/semrel-plugin-updater-go
    env:
      SEMREL_PLUGIN_FILE: "internal/version/version.go"
      SEMREL_PLUGIN_VARIABLE: "Version"
```

## `SEMREL_PLUGIN_*` variables

| Name | Required | Description | Default |
| --- | --- | --- | --- |
| `SEMREL_PLUGIN_FILE` | Optional | Path to the Go source file containing the version variable. | version.go |
| `SEMREL_PLUGIN_VARIABLE` | Optional | Go variable name to update. | Version |

## `SEMREL_*` release context used

| Variable | Description |
| --- | --- |
| `SEMREL_VERSION` | Resolved release version for the current run. |
| `SEMREL_NEXT_VERSION` | Next version computed by semrel for the release. |
| `SEMREL_DRY_RUN` | Whether semrel is running in dry-run mode. |

## Example behavior

The plugin rewrites the selected Go version variable with the next release version. In dry-run mode it reports the edit only.

## License

Apache-2.0
