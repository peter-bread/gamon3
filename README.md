# Gamon3

Automatically switch GitHub CLI account on `cd`.

## Requirements

- Go 1.25
- [`gh`](https://cli.github.com/) v2.40.0+

## Installation

## Usage

### GH CLI

Authenticate any GitHub accounts with the GH CLI tool.

### Config File

First create `config.yaml` one of:

1. `$GAMON3_CONFIG_DIR`
1. `$XDG_CONFIG_HOME/gamon3`
1. `$HOME/.config/gamon3`

**NOTE:** `config.yaml` takes precedence over `config.yml` if both exist.

There are two top-level fields:

- `default`: (required) This should be your primary, likely personal, GitHub
  account.
- `accounts`: (optional) This is a mapping of GitHub accounts to filepaths in
  which they should be used. Environment variables can be used.

> ![IMPORTANT]
> You **CANNOT** use `~` for your home directory. Use `$HOME` instead.

Example:

```yaml
---
default: peter-bread        # (required)
accounts:
  ak22112:
    - $DEVELOPER/ak22112/
```

### Setup

Run this after editing your config.

Alternatively, put this in your shell rc and restart your shell after
editing your config file.

```bash
gamon3 setup
```

### Hook

Create a hook for shell `cd` command in your shell rc file.

bash:

```bash
cd() {
  builtin cd "$@" || exit
  gamon3 run
}
```

zsh:

```bash
_gamon3_gh_switch() {
  gamon3 run
}

autoload -U add-zsh-hook
add-zsh-hook chpwd _gamon3_gh_switch
```
