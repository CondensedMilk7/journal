# Journal

Super minimal command-line journal manager. It simply creates markdown files with a given date,
then opens it with an editor of your choice. This was developed for personal use.

### Usage

```sh
journal # Creates and opens entry for current date
journal 2025-03-14 # Creates and opens entry for March 3, 2025
```

### Installation

Requires [go](https://go.dev/).

```sh
go install github.com/CondensedMilk7/journal
```

### Configuration

Configuration is read from `$HOME/.config/journal/config.json` (if it exits).
Example Configuration:

```json
{
  "journalDir": "/Users/pridon/Documents/journal",
  "editor": "nvim",
  "filemode": 600
}
```

This is also the default configuration, except for `journalDir`, which is `$HOME/Documents/journal`.
