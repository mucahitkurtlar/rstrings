# rstrings
Display printable strings in directories.

## Usage
```sh
rstrings <path> [options]
```

## Example
```sh
rsrings ./ -n 15 # display every  file's strings in this directory (recursively) with at least 15 printable characters
```

## Installation
Before installing `rstrings`, make sure you have `strings` installed on your system and then you can install already builded binary or build it yourself.
### Binary
Simply install binary from [releases](https://github.com/mucahitkurtlar/rstrings/releases), extract it and put it in your `$PATH` and you are ready to go :)

### Build
Firstly, make sure you have `go` installed on your system and then clone this repository.
```sh
git clone https://github.com/mucahitkurtlar/rstrings.git
```
Change directory to `rstrings/cmd/rstrings/`.
```sh
cd rstrings/cmd/rstrings/
```
Create directory `bin/` and build it.
```sh
mkdir bin/
go build -o bin/rstrings
```
Now you can use `rstrings` command from `bin/` directory. You can also put it in your `$PATH` to use it from anywhere.
