# gwc

`gwc` is a simple implementation of the Unix `wc` (word count) tool written in Go. This project is a learning exercise inspired by Coding Challenges, aiming to deepen understanding of Go and Unix-like command-line utilities.

## Features

- Count lines, words, and bytes in files or standard input.
- Mimics basic functionality of the classic `wc` command.

## Usage

```bash
./gwc [options] [file...]
```

### Options

- `-l`: Count lines.
- `-w`: Count words.
- `-c`: Count bytes.
- `-m`: Count characters.

### Examples

#### Count lines, words, and bytes in a file

```bash
./gwc file.txt
```

#### Count lines only

```bash
./gwc -l file.txt
```

#### Count words in multiple files

```bash
./gwc -w file1.txt file2.txt
```

#### Count characters in multiple files

```bash
./gwc -m file1.txt file2.txt
```

#### Use standard input

```bash
echo "Hello, World!" | ./gwc
```

## Installation

<!--- TODO: update when released to use `go install` --->

1. Make sure you have Go installed on your system (version 1.20 or later).
2. Clone the repository:
   ```bash
   git clone https://github.com/flan6/gwc.git
   cd gwc
   ```
3. Build the project:
   ```bash
   go build -o gwc
   ```

## Running Tests

Run tests using:

```bash
go test ./...
```

## Contributing

This project is a personal learning exercise, but feedback and suggestions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

---

### Acknowledgments

- Inspired by challenges from [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc/).
- Special thanks to the Go community for their resources and tools.

### Disclaimer

This is a learning project and may not handle edge cases or performance optimizations as rigorously as the standard `wc` tool.
