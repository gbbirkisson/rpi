package rpi

var version string = ""
var revision string = "development"

// Returns the version and revision of rpi. These are set during the compilation of the binaries.
func GetVersion() (string, string) {
	return version, revision
}
