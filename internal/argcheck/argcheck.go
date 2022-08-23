package argcheck

// Returns true if the given arguments contains "help" argument
func CheckHelp(args []string) (isThereHelp bool) {
	isThereHelp = false
	for _, arg := range args {
		if isThereHelp = arg == "-h" || arg == "--help" || arg == "help"; isThereHelp {
			break
		}
	}
	return
}

func CheckVersion(args []string) (isThereVersion bool) {
	isThereVersion = false
	for _, arg := range args {
		if isThereVersion = arg == "-v" || arg == "-V" || arg == "--version" || arg == "version"; isThereVersion {
			break
		}
	}
	return
}
