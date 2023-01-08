# CRYTTIAN

<br />

A small tool for choose and set a color scheme
for your alacritty terminal.

<br />
<br />

## CONFIGURATION

First of all please **make a backup copy of your configuration file**.
Because `cryttian` will edit the original one.

The configuration is pretty simple. Make sure to have
the configuration file at the following path:

```bash
~/.config/alacritty/alacritty.yml
```

Then inside `~/.config/alacritty/` create a folder called `themes`

```bash
mkdir `~/.config/alacritty/themes`
```

Inside `themes` you must place your color schemes as yml files.

## USAGE

You can perform two actions:

- list
- apply

`cryttian list`

Show a list of your available color's themes

`cryttian apply <theme>`

Apply a specific theme.

If the `theme` argument is not provided. The tool list through
[fzf](https://github.com/junegunn/fzf) all the available themes.

The tool simply merge the two files. No strange magics.
