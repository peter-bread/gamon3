# Gamon3

Automatically switch GitHub CLI account on `cd`.

## Requirements

- Go 1.25+
- [`gh`](https://cli.github.com/) v2.40.0+

## Install

```bash
curl -fsSL https://raw.githubusercontent.com/peter-bread/gamon3/refs/heads/main/scripts/install | bash
```

```bash
wget -qO- https://raw.githubusercontent.com/peter-bread/gamon3/refs/heads/main/scripts/install | bash
```

**TODO** To provide arguments:

```bash
curl -fsSL https://raw.githubusercontent.com/peter-bread/gamon3/refs/heads/main/scripts/install | bash -s -- --version latest --ext tar.gz --prefix /usr/local
```

### Homebrew

```bash
brew install peter-bread/tap/gamon3
```

### Go Install

```bash
go install github.com/peter-bread/gamon3@latest
```

### Build From Source

```bash
git clone https://github.com/peter-bread/gamon3
cd gamon3
go build
cp ./gamon3 ~/.local/bin
```

Assuming `~/.local/bin` is in `PATH`.

## Usage

### Authenticate with GH CLI

Before using Gamon3, you will need to [authenticate your GitHub
account(s)](https://cli.github.com/manual/gh_auth_login) with the GH CLI.

### Setup shell to use Gamon3

#### Bash

Add the following to `~/.bashrc`:

```bash
eval "$(gamon3 hook bash)"
```

#### Zsh

Add the following to `~/.zshrc`:

```bash
eval "$(gamon3 hook zsh)"
```

#### Fish

Add the following to `~/.config/fish/config.fish`:

```fish
gamon3 hook fish | source
```

### Configure Gamon3

As a minimum, create a config file:

```bash
mkdir -p "$HOME/.config/gamon3" && touch "$HOME/.config/gamon3/config.yaml"
```

Then put this inside:

```yaml
---
default: <your-primary-github-account>
```

For more detail, see [Configuration](#configuration).

## Configuration

### Config File Location

Gamon3 will check 3 locations for a config file:

1. `$GAMON3_CONFIG_DIR/config.yaml`
1. `$XDG_CONFIG_HOME/gamon3/config.yaml`
1. `$HOME/.config/gamon3/config.yaml`

Using `config.yml` is also supported.

### Config File Structure

| Field      | Required | Type                   | Description                                                            |
| ---------- | -------- | ---------------------- | ---------------------------------------------------------------------- |
| `default`  | yes      | `string`               | Primary GitHub account for the user                                    |
| `accounts` | no       | `string` -> `string[]` | Mapping of GitHub accounts to directories in which they should be used |

E.g.

```yaml
---
default: primary-account
accounts:
  work-account:
    - $HOME/repos/work
    - ~/work/github/
    - $WORK
  some-other-account:
    - $HOME/other-stuff/
```

This config file is especially useful if projects are organised by GitHub
account.

### Overrides

The default configuration file can be overridden in two ways:

- a local `.gamon.yaml` config file, or
- the `GAMON3_ACCOUNT` environment variable.

These overrides are useful if projects are not always organised by GitHub
account.

#### Local Config File

Gamon3 will search (inclusively) upward from `$PWD` to `$HOME` for a file
called `.gamon.yaml` or `.gamon.yml`. This file should contain a single
`account` key with a `value` being the GitHub account to use.

E.g.

```yaml
---
account: some-account
```

#### Environment Variable

Gamon3 will check to see if the `GAMON3_ACCOUNT` environment variable has been
set to a valid GitHub account. If it has, this will override both `.gamon.yaml`
and `config.yaml`.

## See Also

- [Homebrew Tap](https://github.com/peter-bread/homebrew-tap)
