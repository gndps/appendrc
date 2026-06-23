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
eval "$(appendrc source)"     # source all files
eval "$(appendrc source --fast)"  # source only files that load in <1s
```

## Commands

| Command | Description |
|---------|-------------|
| `add <name> [content]` | Create a sourcable file (opens `$EDITOR` if no content) |
| `edit <name>` | Open file in `$EDITOR` |
| `remove <name>` | Delete a sourcable file |
| `list` | List all files with timing info |
| `time` | Benchmark all files and rebuild `fast.sh` |
| `source [--fast]` | Print source commands for eval |
| `path add/remove/list` | Manage a PATH list |

## Shell integration

Add to `~/.bash_profile` or `~/.zshrc`:

```sh
eval "$(appendrc source)"
# or for fast startup:
eval "$(appendrc source --fast)"
```
