# Gamon3

Automatically switch GitHub CLI account on `cd`.

## Requirements

- Go 1.25
- [`gh`](https://cli.github.com/) v2.40.0+

## Get Started

### 1. Install

TODO: Add installation instructions.

### 2. Setup shell to use Gamon3

Bash:

```bash
eval "$(gamon3 hook bash)"
```

Zsh:

```bash
eval "$(gamon3 hook zsh)"
```

### Configure Gamon3

See [Configuration](#configuration)

## Configuration

Create a config file.

```bash
mkdir -p "$HOME/.config/gamon3" && touch "$HOME/.config/gamon3/config.yaml"
```

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
    - $HOME/work/github/
    - $WORK
  some-other-account:
    - $HOME/other-stuff/
```

A minimal `config.yaml` would just be:

```yaml
---
default: primary-account
```

This config file is especially useful if projects are organised by GitHub
account.

> [!IMPORTANT]
> You **CANNOT** use `~` for your home directory in `config.yaml`.
>
> Use `$HOME` instead.
>
> ---
>
> _This should be supported eventually. See [this issue](https://github.com/peter-bread/gamon3/issues/5)_.

### Overrides

The default configuration file can be overridden in two ways:

- a local `.gamon.yaml` config file, or
- the `GAMON3_ACCOUNT` environment variable

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
set to a valid GitHub account.
