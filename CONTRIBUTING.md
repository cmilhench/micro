# Contributing

Thank you for considering contributing to the project! Here are some guidelines
to help you get started:

# Getting started

**Tl;Dr**: So you want to run this project or even contribute to it and don't
have time to get up to speed on everything. This is the section to read.

1. Create a new branch for your feature or bugfix.
1. Write your code.
1. Write tests for your code.
1. Run tests locally to ensure everything works.
1. Commit your changes and push your branch.
1. Create a pull request.

Please follow the coding style used in the repository. If you encounter a bug, 
please create an issue and include details on how to reproduce it. Include the
version of the project you are using and any relevant logs.

## Source control

There are no `pre-commit` hooks. Imagine Microsoft Word, Google Docs or any
other application preventing you from saving until all spelling mistakes have 
been fixed.  We're all about removing as much friction as possible, and we
trust you to be a good participant.

The project uses a feature-branching strategy. Commits that are merged into 
main/master are considered shippable. Use a branch for development.

Tagging is used to mark notable milestones, typically releases e.g. `v1.0.0`.
The build system will embed the latest git tag and commit SHA1 within the
binary and docker image.

Version numbers follow a semantic versioning (SemVer) scheme conveying meaning
about the underlying changes in a release. It uses a three-part version number 
format: `MAJOR.MINOR.PATCH`, where each part signifies the nature and scope of
changes.

There is no enforced naming convention for git branches or commit messages. 
However using a slash ('/') in the name helps organize and categorize branches
in a hierarchical manner.  

A good commit message should be concise yet informative. It is recommended to
keep it within 50 characters. If more detail is needed, you can add a body to
the commit message.  Each line in the body should be wrapped to enhance
readability.

> [!TIP]
> Your code, commit messages and branch names will be available for all to
> read. A simple rule of thumb is to stay in keeping with the codebase. If
> significant drift occurs your contribution could be lost or omitted.

## Approach

### Do not use frameworks, build what you want with the best practices.

We believe in the philosophy of using the smallest, sharpest tools to build the
things we need. Go matches this philosophy, it doesn't provide any opinionated 
frameworks, instead providing a better standard lib that can be used to write
customized code which can be read easily and is very specific to the task being
performed. Out with one line of magic, and in with 10 lines of readable code 
that does a very specific task. In other words, write customized coding over 
configurations and conventions.

Where possible we make use of the standard library and make decisions that add
as little cognitive load to the project as possible.  At each point the focus 
is on creating loosely coupled application components that can be easily
connected to the software environment maintaining agility.

## Development

This project makes use of the gnu make task runner. Type `make` for available
commands defined as targets described in Makefile.

The Makefile should work on a Windows laptop, provided you're using an
appropriate environment like Git Bash or WSL that supports the required shell
commands and syntax.

### Running the project locally

```sh
make start
```

> [!TIP]
> You can use `make watch` which will monitor for changes and restart.

Some commands–such as `watch`–require dependencies, including:
- entr

***If entr is unavailable on Windows, replace with watchexec or similar***

There is a hierarchy to some commands which depend upon others, as follows:
`make watch` > `make start` > `make test` > `make lint` > `make gen` > `clean`.

The `make watch` command will run `make start` if any file changes in any
subdirectory excluding vendor, docs and hidden subdirectories.

The `make start` command will run `make test` before assigning current
environment variables from the `.env` file and run the main source file

The `make test` command will run `make lint` before executing all test cases
of the current package. These test will be included in files ending with 
`_test.go`

The `make lint` command will run `make gen` before executing the Go vet command

The `make gen` command will execute any code generation logic identified in Go
source files with the `//go:generate` comment.

In this project "go generate" is used to produce a `.version` file which will
be embedded within the binary for versioning purposes.  You can see this with
the `//go:embed` comment.

For development you may simply use `make watch`.
