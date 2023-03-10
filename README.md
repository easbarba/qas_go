# Qas

Easily manage multiple FLOSS repositories.

## API

`qas_api` provides all projects to `qas`, so it needs to be running.

https://github.com/easbarba/qas_api

## Usage

Grab all projects locally:

    $ qas --grab

Archive selected projects at `$HOME/Downloads/archived`:

    $ qas --archive --name awesomewm,nuxt

## Configuration

`qas` looks for configuration files at `$XDG_CONFIG/qas`:

$XDG_CONFIG/qas/misc.json

```json
[
  {
    "name": "awesomewm",
    "branch": "master",
    "url": "https://github.com/awesomeWM/awesome"
  },
  {
    "name": "nuxt",
    "branch": "main",
    "url": "https://github.com/nuxt/framework"
  },
  {
    "name": "swift_format",
    "branch": "main",
    "url": "https://github.com/apple/swift-format"
  }
]
```

## Installation

Get the needed dependencies and install with:

    $ make deps & make install

## GNU Guix

To load all system dependencies, just run `guix shell`

## TODO

### High

- add/overwrite configuration via cli: `$ qas --add 'go,gum,main,https://github.com/charmbracelet/gum'`
- log history of commands

### Low

- more management utilities
- list current configuration.

## History

`qas` initially was a module of a bigger package called `cejo` extracted as
standalone project to follow UNIX main rule: 'do one thing, well'.

## License

[GPL-v3](https://www.gnu.org/licenses/gpl-3.0.en.html)
