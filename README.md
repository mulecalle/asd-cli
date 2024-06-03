# Asd CLI

A CLI to use when reverse-i-search is not enought.

## Usage

### Get available domains

Run the binary with the command `notes` get all the available domains:

```bash
> asd notes
DOMAIN
example
```

### Get notes in domain

Run the `notes` command using `-d` to specify the name of the domain:

```bash
> asd notes -d example
hello:
    - world
foo:
    - bar1
    - bar1
```

### Get a note in domain

Run the `notes` command using `-d` to specify the name of the domain, and `-n` the name of the note:

```bash
> asd notes -d example -n hello
    - world
```

### Domains and Notes lifecycle

`Domains` are represented as yaml files under the [data](./cmd/data/) folder, each entry in the yaml represents a `Note`.

## Contributors

Big thanks to all the people who have already contributed! (lol)



[![contributors](https://contrib.rocks/image?repo=corentinth/it-tools)](https://github.com/mulecalle/asd-cli/graphs/contributors)

