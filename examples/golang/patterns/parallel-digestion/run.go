package parallel_digestion

import (
	stdctx "context"
	"crypto/md5"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
)

func Run(ctx stdctx.Context) error {
	ctx, cancel := stdctx.WithCancel(ctx)
	defer cancel()

	tmp, err := os.MkdirTemp("", "mddocks-parallel-digest-*")
	if err != nil {
		return fmt.Errorf("parallel-digestion: MkdirTemp: %w", err)
	}
	defer os.RemoveAll(tmp)

	contents := map[string][]byte{
		"a.txt": []byte("hello"),
		"b.txt": []byte("world"),
		"c.txt": []byte("!"),
	}

	expected := map[string][md5.Size]byte{}
	for name, data := range contents {
		full := filepath.Join(tmp, name)
		if err := os.WriteFile(full, data, 0o600); err != nil {
			return fmt.Errorf("parallel-digestion: WriteFile: %w", err)
		}
		expected[full] = md5.Sum(data)
	}

	pathsCh := make(chan string, 8)
	resultsCh := make(chan string, 8)
	errc := make(chan error, 1)

	workerCount := 2
	var wg sync.WaitGroup
	wg.Add(workerCount)
	var errOnce sync.Once
	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			for p := range pathsCh {
				select {
				case <-ctx.Done():
					return
				default:
				}
				data, err := os.ReadFile(p)
				if err != nil {
					errOnce.Do(func() { errc <- err })
					cancel()
					return
				}
				sum := md5.Sum(data)
				exp, ok := expected[p]
				if !ok || sum != exp {
					errOnce.Do(func() { errc <- fmt.Errorf("parallel-digestion: mismatch for %s", p) })
					cancel()
					return
				}
				// Mark processed.
				select {
				case <-ctx.Done():
					return
				case resultsCh <- p:
				}
			}
		}()
	}

	// Producer: walk tree and feed paths.
	go func() {
		defer close(pathsCh)
		walkErr := filepath.WalkDir(tmp, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case pathsCh <- path:
				return nil
			}
		})
		if walkErr != nil {
			errOnce.Do(func() { errc <- walkErr })
			cancel()
		}
	}()

	// Close results when workers exit.
	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	got := map[string]bool{}
	for p := range resultsCh {
		got[p] = true
	}

	select {
	case e := <-errc:
		if e != nil {
			return fmt.Errorf("parallel-digestion: %w", e)
		}
	default:
	}

	if len(got) != len(expected) {
		return fmt.Errorf("parallel-digestion: expected %d files got %d", len(expected), len(got))
	}
	for p := range expected {
		if !got[p] {
			return fmt.Errorf("parallel-digestion: missing processed file %s", p)
		}
	}
	return nil
}

