# Go ed25519 to OpenSSH format converter

This little tool converts ed25519 private keys from the format used in Golang to a format that OpenSSH understands. The Go private key (input) is expected to be in PEM format.

## Installation

```
go install github.com/ericselin/ed25519-openssh
```

## Usage

You can either specify a filename as an argument, or pipe in the key from another command.

Specify filename:
```
ed25519-openssh private_key_pem_file
```

Pipe (e.g. from a paste)
```
wl-paste | ed25519-openssh
```

The resulting key is written to stdout. What you do with it is up to you, but one possibility is to redirect the output to a file (`... > ed25519_file`).
