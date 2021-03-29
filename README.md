# Tempi

Tempi is a minimalist template engine.

## Installation

```bash
go get github.com/yozel/tempi/cmd
```

## Usage

```bash
# text format, output to stdout
tempi -f example/values.yaml -o txt:-

# text format, output to file
tempi -f example/values.yaml -o txt:letter.txt

# pdf format, output to file
tempi -f example/values.yaml -o pdf:letter.pdf
```
