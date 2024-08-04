package constant

import "os"

const DefaultFilePerm = 0755

var Home string

func init() {
	Home = os.Getenv("HOME")
}
