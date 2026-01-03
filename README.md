# kata-to-go

A CLI tool to quickly bootstrap Go projects for Codewars kata challenges.

## Installation

### Install from source

```bash
go install github.com/johanovskyt/kata-to-go@latest
```

### Build locally

```bash
git clone https://github.com/johanovskyt/kata-to-go.git
cd kata-to-go
go install
```

## Usage

Create a new kata project by providing the kata ID and target directory:

```bash
kata-to-go new <kata-id> <path>
```

**Example:**

```bash
kata-to-go new 5270d0d18625160ada0000e4 ./my-katas
```

## License

MIT
