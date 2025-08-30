# Gamon3

Automatically switch GitHub CLI account on `cd`.

## Requirements

- Go 1.25
- [`gh`](https://cli.github.com/) v2.40.0+

## Installation

## Usage

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
