# tlccryptotool

tlccryptotool is a command line tool for encrypting and decrypting files using the Golang native crypto library.

## Usage

```bash
./tlccryptotool aes encrypt -k "this_must_be_of_32_byte_length__" -v "this is the sentence that needs to be encrypted"
```

```bash
./tlccryptotool aes decrypt -k "this_must_be_of_32_byte_length__" -v "71ctABqAXNuhprLn_Cl6V6DRd0fxxe6RCUhvuTQiiYVQ6aB3d_xquNqIivoqZMox18J2UmfvQtNRUzL7C4BVDDR_i0XQC20kVfSQ"
```

## Compiling the Tool

### MacOS

64-bit (M1 chip)

```bash
GOOS=darwin GOARCH=arm64 go build -o tlccryptotool main.go
```

64-bit (Intel chip)

```bash
GOOS=darwin GOARCH=amd64 go build -o tlccryptotool main.go
```

32-bit (Intel chip)

```bash
GOOS=darwin GOARCH=386 go build -o tlccryptotool main.go
```

### Windows

64-bit (Intel chip)

```bash 
GOOS=windows GOARCH=amd64 go build -o tlccryptotool.exe main.go
```

32-bit (Intel chip)

```bash
GOOS=darwin GOARCH=amd64 go build -o tlccryptotool main.go
```

### Linux

64-bit (Intel chip)

```bash 
GOOS=linux GOARCH=amd64 go build -o tlccryptotool.exe main.go
```

32-bit (Intel chip)

```bash
GOOS=linux GOARCH=386 go build -o tlccryptotool main.go
```