package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	root := os.Getenv("GHQ_ROOT")
	if len(root) < 1 {
		log.Fatal("env $GHQ_ROOT must be set")
	}

	const maxDepth = 6
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return errors.Wrap(err, "failed to walk dir")
		}

		if d.IsDir() && strings.Count(path, string(os.PathSeparator)) > maxDepth {
			return fs.SkipDir
		}

		if d.IsDir() && strings.Count(path, string(os.PathSeparator)) == maxDepth {
			fmt.Println(path)
		}

		return nil
	})
}
