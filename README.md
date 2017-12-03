# freenas-notify

## Description

A command line utility that sends status message to RabbitMQ. Currently the utility is hardcoded to send message to a
topic exchange called `freenas`.

## Usage

For example,

```bash
./freenas-notify --url amqp://<<username>>:<<password>>@<<address>>:<<port>>/ send shutdown
```


## Install

To install, use `go get`:

```bash
$ go get -d github.com/Hxucaa/freenas-notify
```

## Contribution

1. Fork ([https://github.com/Hxucaa/freenas-notify/fork](https://github.com/Hxucaa/freenas-notify/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[Hxucaa](https://github.com/Hxucaa)
