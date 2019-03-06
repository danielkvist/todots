// Package copier implements utilities for copying files
package copier

import (
	"fmt"
	"io"
	"os"
)

// Copy receives a source file and a destination folder and tries
// to make a copy of the source file on a subfolder inside the destination
// folder. If there is any error, it returns it.
func Copy(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	defer f.Close()

	fi, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// TODO: Work with directories too
	// 			Or with multiple entries in the .yaml file
	if !fi.Mode().IsRegular() {
		return fmt.Errorf("%q is not a regular file", fi.Name())
	}

	_, err = os.Stat(dst)
	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("%v", err)
		}

		if err := os.Mkdir(dst, os.ModePerm); err != nil {
			return fmt.Errorf("while making the destination directory for %q on %q: %v", fi.Name(), dst, err)
		}
	}

	df, err := os.Create(dst + fi.Name())
	if err != nil {
		return fmt.Errorf("while creating destination file %q on %q: %v", fi.Name(), dst, err)
	}

	if _, err := io.Copy(df, f); err != nil {
		return fmt.Errorf("while copying from %q to %q: %v", f.Name(), df.Name(), err)
	}

	return nil
}
