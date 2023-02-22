package main

import (
	"encoding/pem"
	"io"
	"os"

	"github.com/mikesmitty/edkey"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/ssh"
)

func main() {
	r := getReader()
	b := getPrivateKey(r)
	s := getOpenSshPem(b)
	os.Stdout.Write(s)
}

// getReader gets a reader for the file specified in the command args,
// or stdin if the program is used in a pipe.
// It panics in case of errors.
func getReader() io.Reader {
	if len(os.Args) == 2 {
		filename := os.Args[1]
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		return file
	} else if info, _ := os.Stdin.Stat(); info.Mode()&os.ModeCharDevice == 0 {
		return os.Stdin
	}
	panic("No file specifid and not used in pipe")
}

// getPrivateKey gets the private key from the bytes in the given reader.
// It panics in case of errors.
func getPrivateKey(r io.Reader) ed25519.PrivateKey {
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	privKey, err := ssh.ParseRawPrivateKey(b)
	if err != nil {
		panic(err)
	}
	if edKey, ok := privKey.(ed25519.PrivateKey); ok {
		return edKey
	}
	panic("Could not parse private key as ed25519")
}

// getOpenSshPem gets the private key in openssh format from the given ed25519 bytes.
func getOpenSshPem(b ed25519.PrivateKey) []byte {
	pemKey := &pem.Block{
		Type:  "OPENSSH PRIVATE KEY",
		Bytes: edkey.MarshalED25519PrivateKey(b),
	}
	return pem.EncodeToMemory(pemKey)
}
