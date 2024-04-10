package variables

import "strconv"

// NOTE: read only variables
const hello = "Hello, "

// NOTE: read and write variables
var text string = "sample text"

type Os int

const (
	Windows Os = iota + 1
	Linux
	Mac
)

func Hello(name string) string {
	// NOTE: local variable
	title := "Mr."
	return hello + title + name
}

func GetOsName(os Os) string {
	switch os {
	case Windows:
		return "Windows" + strconv.Itoa(int(os))
	case Linux:
		return "Linux" + strconv.Itoa(int(os))
	case Mac:
		return "Mac" + strconv.Itoa(int(os))
	default:
		return "Unknown"
	}
}
