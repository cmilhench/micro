# Contributing

Before contributing you'll need a few things:

- docker
- golint `go install golang.org/x/lint/golint@latest`

Type `make` for available development commands, for example;

## Development

`make watch`

The watch command requires `entr` which can be installed via brew.

Use `make dev-db` to start a development amqp service as well as example 
services with `make add` and `make log`.

## Testing

`make test`

`make run`

