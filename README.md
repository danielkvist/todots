# ToDots

todots is a simple CLI writen in Go to make a copy of your dotfiles.

## Installation

```bash
$ go get -u github.com/danielkvist/todots
```

Or

```bash
$ go install github.com/danielkvist/todots
```

## Usage

```bash
todots --dst ./backup
```

Or

```bash
todots
```

todots by default looks for a configuration file that should be on your \$HOME directory called ```.todots.yaml```.
However you can specify another file using the `--config` flag

```bash
todots --config myconfig.yml
```

## Configuration file

Here's an example of what your configuration file should look like:

> The paths need to be relative to your $HOME.

```yaml
vim: .vimrc
fish: .config/fish/config.fish
```

You can see an example of the final result [here](https://github.com/danielkvist/dotfiles).
