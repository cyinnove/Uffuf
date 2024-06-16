# Uffuf

Uffuf is a command-line tool written in Go that utilizes the `ffuf` (Fuzz Faster U Fool) tool to fuzz URLs with a given wordlist and capture unique results. It supports both GET and POST HTTP methods and writes the unique results to an output file.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Options](#options)
- [Example](#example)
- [License](#license)

## Installation

### Prerequisites

- Go 1.15 or later
- `ffuf` (Fuzz Faster U Fool) installed and available in your PATH. You can install `ffuf` from its [GitHub repository](https://github.com/ffuf/ffuf).

### Steps

1. Clone the repository or download the source code.
2. Build the Go program:

```bash
go build -o uffuf uffuf.go
```

This will create an executable named `uffuf`.

## Usage

```bash
./uffuf -l <list_of_urls> [-w <wordlist>] [-o <output_file>]
```

### Options

- `-w <wordlist>`: Path to the wordlist file (default: `sallam/Wordlists/elisa_fuzz.txt`)
- `-l <list_of_urls>`: File containing a list of URLs (required)
- `-o <output_file>`: File to write unique results (default: `output.txt`)

## Example

1. Prepare a file `urls.txt` with the list of URLs you want to fuzz:

```
http://example.com
http://test.com
```

2. Run the tool:

```bash
./uffuf -l urls.txt -w /path/to/wordlist.txt -o results.txt
```

3. The unique results will be written to `results.txt`.

### Sample Output

When running the tool, you will see debug output showing the `ffuf` commands being executed and the unique results found:

```plaintext
[DEBUG] Running command: ffuf -u http://example.com/FUZZ -w /path/to/wordlist.txt -X GET
[DEBUG] Running command: ffuf -u http://example.com/FUZZ -w /path/to/wordlist.txt -X POST
[DEBUG] Found unique result: http://example.com/test
[DEBUG] Running command: ffuf -u http://test.com/FUZZ -w /path/to/wordlist.txt -X GET
[DEBUG] Running command: ffuf -u http://test.com/FUZZ -w /path/to/wordlist.txt -X POST
```

## License

This project is licensed under the MIT License.

---

## Contributing

Feel free to submit issues, fork the repository, and send pull requests. Contributions are always welcome.

## Acknowledgments

- [ffuf](https://github.com/ffuf/ffuf) - The fuzzing tool used by Uffuf.
