package osvdb

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
)

// An Input provides access to JSON-encoded OSV data
type Input func(context.Context) (data io.Reader, cleanup func(), err error)

// NewFileInput returns an Input that reads from a file.
func NewFileInput(path string) Input {
	return func(ctx context.Context) (io.Reader, func(), error) {
		f, err := os.Open(path)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open file: %w", err)
		}

		cleanup := func() {
			_ = f.Close()
		}

		return f, cleanup, nil
	}
}

// NewFSInput returns an Input that reads from JSON files in a file system.
func NewFSInput(fsys fs.FS, recursive bool) Input {
	const fileSuffix = ".json"

	return func(ctx context.Context) (io.Reader, func(), error) {
		var readers []io.Reader

		walkFn := func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				if !recursive {
					return fs.SkipDir
				}
				return nil
			}

			if strings.HasSuffix(d.Name(), fileSuffix) {
				file, err := fsys.Open(path)
				if err != nil {
					return fmt.Errorf("failed to open file: %w", err)
				}
				defer file.Close()

				buf := new(bytes.Buffer)
				_, err = io.Copy(buf, file)
				if err != nil {
					return fmt.Errorf("failed to read file: %w", err)
				}

				readers = append(readers, buf)
			}

			return nil
		}

		if err := fs.WalkDir(fsys, ".", walkFn); err != nil {
			return nil, nil, fmt.Errorf("failed to walk directory: %w", err)
		}

		return io.MultiReader(readers...), nil, nil
	}
}
