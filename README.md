# appendrc

Manage sourcable shell files in `~/.appendrc/sourcables/`. Tracks source timing so you can skip slow files with `--fast`.

## Install

```sh
brew install gndps/tap/appendrc
```

## Usage

```sh
appendrc add aliases          # create/edit a sourcable file
appendrc list                 # list all files with timing
appendrc time                 # benchmark all files, rebuild fast.sh
eval "$(appendrc source)"     # source all files (quiet)
eval "$(appendrc source --fast)"  # source only files that load in <1s
eval "$(appendrc source -v)"  # source all files, announcing each one
```

## Commands

| Command | Description |
|---------|-------------|
| `add <name> [content]` | Create a sourcable file (opens `$EDITOR` if no content) |
| `edit <name>` | Open file in `$EDITOR` |
| `remove <name>` | Delete a sourcable file |
| `list` | List all files with timing info |
| `time` | Benchmark all files and rebuild `fast.sh` |
| `source [--fast] [-v]` | Print source commands for eval. Quiet by default; `-v`/`--verbose` announces each file as it is sourced |
| `path add/remove/list` | Manage a PATH list |

## Shell integration

Add to `~/.bash_profile` or `~/.zshrc`:

```sh
eval "$(appendrc source)"
# or for fast startup:
eval "$(appendrc source --fast)"
```

Sourcing is quiet by default. Add `-v`/`--verbose` to have each file
announced on the terminal as it is sourced, e.g. `eval "$(appendrc source --fast -v)"`.
