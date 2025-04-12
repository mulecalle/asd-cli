# Asd CLI

A CLI tool to quickly access and manage your command notes and snippets. Perfect for when reverse-i-search isn't enough!

## Installation

```bash
# Clone the repository
git clone https://github.com/mulecalle/asd-cli.git
cd asd-cli

# Build and install
make install
```

## Usage

### List Available Domains

List all available domains (note categories):

```bash
asd notes
```

Example output:
```
DOMAIN
k8s
git
aws
```

### List Notes in a Domain

View all notes within a specific domain:

```bash
asd notes -d <domain>
```

Example:
```bash
asd notes -d k8s
```

### Get a Specific Note

Retrieve a specific note from a domain:

```bash
asd notes -d <domain> -n <note>
```

Example:
```bash
asd notes -d k8s -n get-jobs-by-namespace
```

## Development

### Project Structure

- `cmd/`: Contains the CLI commands
  - `data/`: YAML files containing notes for each domain
- `utils/`: Utility functions for the CLI
- `main.go`: Application's entrypoint

### Adding New Notes

1. Create or edit a YAML file in `cmd/data/` with your domain name (e.g., `k8s.yaml`)
2. Add your notes in YAML format:
   ```yaml
   note-name:
     - command or note content
     - additional content if needed
   ```

### Testing

Run the test suite with coverage:

```bash
make coverage
```

This will:
- Run all tests
- Generate a coverage report
- Create an HTML coverage report (coverage.html)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Contributors

Big thanks to all the people who have already contributed!

[![contributors](https://contrib.rocks/image?repo=mulecalle/asd-cli)](https://github.com/mulecalle/asd-cli/graphs/contributors)
