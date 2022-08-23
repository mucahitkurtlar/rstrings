package rstrings

import (
	"github.com/mucahitkurtlar/rstrings/internal/argcheck"
	"github.com/mucahitkurtlar/rstrings/pkg/command"
	"github.com/mucahitkurtlar/rstrings/pkg/files"
)

// Rstrings returns the printable strings in the given directory.
func Rstrings(args ...string) (outputs []string, err error) {
	if len(args) == 0 {
		return nil, &ArgumentError{}
	}
	if argcheck.CheckHelp(args) {
		return []string{HelpText()}, nil
	}
	if argcheck.CheckVersion(args) {
		return []string{VersionText()}, nil
	}
	filepaths, err := files.GetFiles(args[0])
	if err != nil {
		return
	}
	argsWithoutDir := args[1:]
	for _, path := range filepaths {
		allArgs := append(argsWithoutDir, path)
		output, err := command.Run("strings", allArgs...)
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, output)
	}
	return
}

// Return the help text for the rstrings command.
func HelpText() string {
	return `Usage: rstrings <directory> [option(s)]
	Display printable strings in directories (stdin by default)
	The options are:
	 -a - --all                Scan the entire file, not just the data section [default]
	 -d --data                 Only scan the data sections in the file
	 -f --print-file-name      Print the name of the file before each string
	 -n <number>               Locate & print any sequence of at least <number>
	   --bytes=<number>         displayable characters.  (The default is 4).
	 -t --radix={o,d,x}        Print the location of the string in base 8, 10 or 16
	 -w --include-all-whitespace Include all whitespace as valid string characters
	 -o                        An alias for --radix=o
	 -T --target=<BFDNAME>     Specify the binary file format
	 -e --encoding={s,S,b,l,B,L} Select character size and endianness:
							   s = 7-bit, S = 8-bit, {b,l} = 16-bit, {B,L} = 32-bit
	 --unicode={default|show|invalid|hex|escape|highlight}
	 -U {d|s|i|x|e|h}          Specify how to treat UTF-8 encoded unicode characters
	 -s --output-separator=<string> String used to separate strings in output.
	 @<file>                   Read options from <file>
	 -h --help                 Display this information
	 -v -V --version           Print the program's version number
   rstrings: supported targets: elf64-x86-64 elf32-i386 elf32-iamcu elf32-x86-64 pei-i386 pe-x86-64 pei-x86-64 elf64-little elf64-big elf32-little elf32-big pe-bigobj-x86-64 pe-i386 elf64-bpfle elf64-bpfbe srec symbolsrec verilog tekhex binary ihex plugin
   rstrings uses strings from GNU binutils
   Report bugs to <https://github.com/mucahitkurtlar/rstrings/issues/>`
}

func VersionText() string {
	return "rstrings 1.0.1"
}
