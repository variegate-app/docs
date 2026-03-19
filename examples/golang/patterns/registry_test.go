package patterns

import (
	"context"
	"os"
	"path/filepath"
	"sort"
	"testing"
	"time"
)

func TestAllExamplesRun(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatalf("ReadDir: %v", err)
	}

	var slugs []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		slug := e.Name()
		// Each runnable example is represented by a subdir containing run.go.
		if _, err := os.Stat(filepath.Join(slug, "run.go")); err == nil {
			slugs = append(slugs, slug)
		}
	}
	sort.Strings(slugs)

	if len(slugs) == 0 {
		t.Fatal("no example packages found (expected subdirs with run.go)")
	}

	for _, slug := range slugs {
		t.Run(slug, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			if err := Run(ctx, slug); err != nil {
				t.Fatalf("Run(%q) returned error: %v", slug, err)
			}
		})
	}
}

