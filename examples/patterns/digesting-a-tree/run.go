package digesting_a_tree

import (
	stdctx "context"
	"crypto/md5"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func Run(ctx stdctx.Context) error {
	tmp, err := os.MkdirTemp("", "mddocks-digest-*")
	if err != nil {
		return fmt.Errorf("digesting-a-tree: MkdirTemp: %w", err)
	}
	defer os.RemoveAll(tmp)

	// Create deterministic small file set.
	contents := map[string][]byte{
		"a.txt": []byte("hello"),
		"b.txt": []byte("world"),
		"c.txt": []byte("!"),
	}

	expected := map[string][md5.Size]byte{}
	for name, data := range contents {
		full := filepath.Join(tmp, name)
		if err := os.WriteFile(full, data, 0o600); err != nil {
			return fmt.Errorf("digesting-a-tree: WriteFile: %w", err)
		}
		expected[full] = md5.Sum(data)
	}

	got := map[string][md5.Size]byte{}

	err = filepath.WalkDir(tmp, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		got[path] = md5.Sum(data)
		return nil
	})
	if err != nil {
		return fmt.Errorf("digesting-a-tree: walk failed: %w", err)
	}

	if len(got) != len(expected) {
		return fmt.Errorf("digesting-a-tree: expected %d files got %d", len(expected), len(got))
	}
	for path, sum := range expected {
		gs, ok := got[path]
		if !ok {
			return fmt.Errorf("digesting-a-tree: missing digest for %s", path)
		}
		if gs != sum {
			return fmt.Errorf("digesting-a-tree: mismatch for %s", path)
		}
	}
	return nil
}

