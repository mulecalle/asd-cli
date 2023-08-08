# Asd CLI

A CLI to use when reverse-i-search is not enought.

## Usage

### Get available domains

Run the binary with the command `notes` get all the available domains:

```bash
> asd notes
DOMAIN
aws
docker
...
```

### Get notes in domain

Run the `notes` command using `-d` to specify the name of the domain:

```bash
> asd notes -d docker
dangling:
    - docker volume rm $(docker volume ls -qf dangling=true)
harbor-login:
    - docker login -u {} -p {} artifacts.msap.io
logs:
    - docker logs [OPTIONS] CONTAINER
```

### Get a note in domain

Run the `notes` command using `-d` to specify the name of the domain, and `-n` the name of the note:

```bash
> asd notes -d docker -n dangling
- docker volume rm $(docker volume ls -qf dangling=true)
```

## Domains and Notes lifecycle

`Domains` are represented as yaml files under the [data](./cmd/data/) folder, each entry in the yaml represents a `Note`.
