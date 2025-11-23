# Gamon3

Automatically switch GitHub CLI account on `cd`.

## Requirements

- Linux or MacOS
- Go 1.25+
- [`gh`](https://cli.github.com/) v2.40.0+

## Install

### Pre-built Binaries

Pre-built binaries can be found under [GitHub Releases](https://github.com/peter-bread/gamon3/releases/latest).

These can be downloaded and extracted manually.

Alternatively you can use the provided [installation script](./scripts/install). The
commands below download the script and pipe it into Bash.

> [!IMPORTANT]
>
> Always read through scripts before running them to make sure they aren't malicious.

```bash
curl -fsSL https://raw.githubusercontent.com/peter-bread/gamon3/refs/heads/main/scripts/install |
  bash
```

```bash
wget -qO- https://raw.githubusercontent.com/peter-bread/gamon3/refs/heads/main/scripts/install |
  bash
```

#### Customise Installation

The installation script can be configured with some optional flags.

The command below downloads the script and pipes it into Bash, providing all
options with their default values.

```bash
curl -fsSL https://raw.githubusercontent.com/peter-bread/gamon3/refs/heads/main/scripts/install |
  bash -s -- --version latest --ext tar.gz --prefix /usr/local
```

Options:

| Option       | Description        | Default      | Allowed Value          |
| ------------ | ------------------ | ------------ | ---------------------- |
| `--version`  | Version to install | `latest`     | `latest` or `[v]X.Y.Z` |
| `--ext`      | Archive extension  | `tar.gz`     | `tar.gz` or `zip`      |
| `--prefix`   | Install location   | `/usr/local` | Any filepath           |
| `-h, --help` | Prints help        | N/A          | N/A                    |

### Homebrew

```bash
brew install peter-bread/tap/gamon3
```

### Go Install

```bash
go install github.com/peter-bread/gamon3@latest
```

### Build From Source

To build and install Gamon3 under the default prefix (`/usr/local`), run:

```bash
git clone https://github.com/peter-bread/gamon3
cd gamon3
make
sudo make install
```

To install under a custom prefix, e.g. `~/.local`, run:

```bash
make install PREFIX=~/.local
```

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

#### Other Shells

If you use another shell, consult its documentation to see how to hook into the
`cd` command or 'change `PWD`' event. I probably won't spend any time researching
more niche shells that I don't personally use, but feel free to open a pull
request to add support for your favourite shell.

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

> [!NOTE]
>
> You do not need to specify the `default` account in `accounts`, as it will
> always be a fallback option.

### Overrides

The default configuration file can be overridden in two ways:

- a local `.gamon.yaml` or `.gamon3.yaml` config file (`.yaml` or `.yml`), or
- the `GAMON3_ACCOUNT` environment variable.

These overrides are useful if projects are not always organised by GitHub
account.

#### Local Config File

Gamon3 will search (inclusively) upward from `$PWD` to `$HOME` for a file
called `.gamon.yaml` or `.gamon3.yaml`. This file should contain a single
`account` key with a `value` being the GitHub account to use. If `$PWD` is
not a descendant of `$HOME`, the search will continue until the filesystem
root (on Linux and MacOS this is `/`).

E.g.

```yaml
---
account: some-account
```

#### Environment Variable

Gamon3 will check to see if the `GAMON3_ACCOUNT` environment variable has been
set to a valid GitHub account. If it has, this will override both a local
`.gamon3.yaml` and the main `config.yaml`.

### Account Resolution and Errors

Currently, Gamon3 will only report configuration errors if they affect the
account you are trying to switch to. For example, if you have a completely
invalid local config file, but the account is selected via an envionrment
variable, the local config file will never be checked and thus no erros will be
found.

> [!NOTE]
>
> I plan to enventually add a `gamon3 doctor` command that will check all
> discovered config files and report all problems. See [this
> issue](https://github.com/peter-bread/gamon3/issues/22).

## See Also

- [Homebrew Tap](https://github.com/peter-bread/homebrew-tap)
