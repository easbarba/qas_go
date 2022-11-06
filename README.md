# Qas

Easily manage multiple FLOSS repositories.

[hub.docker](https://hub.docker.com/r/easbarbosa/qas)

## Usage

Grab all projects locally: 

    `$ qas --grab`

Archive selected projects at `$HOME/Downloads/archived`: 

    $ qas --archive --name awesomewm,nuxt

## Configuration

`qas` looks for configuration files at `$XDG_CONFIG/qas`:


$XDG_CONFIG/qas/misc.csv
```csv
name,branch,url
awesomewm,master,https://github.com/awesomeWM/awesome
nuxt,main,https://github.com/nuxt/framework
swift_format,main,https://github.com/apple/swift-format
```

## Installation

`Qas` supports both `npm` and `bun`, and then to get the needed dependencies go with:

    $ make deps & make install

## Container

## TODO

- more management utilities

## History

`qas` initially was a module of a bigger package called `cejo` extracted as
standalone project to follow UNIX main rule: 'do one thing, well'.


## License

[GPL-v3](https://www.gnu.org/licenses/gpl-3.0.en.html)

