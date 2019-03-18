# ToDots

todots is a very simple CLI that helps you to easily have a copy of all of your dotfiles.

## Installation

You only need to type the following command in your terminal:

```bash
$ go get -u github.com/danielkvist/todots
```

## Usage

```bash
todots --dst .

# Or

todots
```

todots by default looks for a configuration file that should be on your \$HOME directory called .todots.yaml or .todots.yml.

However you can specify another file using the `--config` flag

```bash
todots --config myconfig.yml --dst .
```

> You always have to specify the destination route

## Configuration file

Here's an example of what your configuration file should look like:

```yaml
# Change user for your username

vim: /home/<user>/.vimrc
i3: /home/<user>/.config/i3/
newsboat: /home/<user>/.newsboat/
tmux: /home/<user>/.tmux.conf
```

As you can see you have to specify the absolute path. And you can specify a file or a directory.

You can see an example of the final result [here](https://github.com/danielkvist/dotfiles).
