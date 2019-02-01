// Package testing the package fixes issue when runnign unittest from non-root folders
package testing

import (
	"os"
	"path"
	"runtime"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	path.Join(path.Dir(filename), "..")
	err := os.Chdir("../../../../")
	if err != nil {
		panic(err)
	}
}
