package files

import (
	"os"
	"path/filepath"
)

// Returns recursive file paths in given directory
func GetFiles(dir string) (paths []string, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	return
}
