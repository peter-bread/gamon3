# How should this all work

## ResolveAccount

This function decides which account should currently be active.

Only error if it involves the account that should be active.

- Check `GAMON3_ACCOUNT`
  - if set
    - return its value
- Check for a local config file
  - if it exists
    - parse it
    - if it has an `account` field
      - if it is non-empty
        - return the value
      - else
        - error: invalid YAML; `account` cannot be empty
    - else
      - error: invalid YAML; there must be an `account` key
- Check for a main config file
  - first in `$GAMON3_CONFIG_DIR`
  - second in `$XDG_CONFIG_HOME/gamon3`
  - finally in `$HOME/.config/gamon3`
  - if file does not exist, either
    - error
    - ignore and exit
  - if file does exist...
