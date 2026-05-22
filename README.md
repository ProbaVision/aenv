# aenv
cockpit view of agent env

## install

```bash
go install github.com/ProbaVision/aenv@latest
```

## supported agent cli

- claude code
- copilot

## usage

- `aenv -h`
- `aenv check`
- `aenv completion bash`
- `aenv completion zsh`

`aenv check` prints YAML entries for detected agent-related paths in `$HOME`, including `.claude`, `.copilot`, `.agent`, `claude.md`, and `agents.md`.

`aenv completion <shell>` prints shell completion for `bash`, `zsh`, `fish`, and `powershell`.

## development

```bash
make test
make build
make install
```

- `make test` runs `go test ./...`
- `make build` runs `go build -o bin/aenv .`
- `make install` runs `go install .`