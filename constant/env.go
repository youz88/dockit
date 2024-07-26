package constant

import "os"

var Home string

func init() {
	Home = os.Getenv("HOME")
}
