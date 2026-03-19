package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"mddocks/examples/golang/patterns"
)

func main() {
	var pattern string
	flag.StringVar(&pattern, "pattern", "", "pattern name (e.g. context, cancellation, pipeline)")
	flag.Parse()

	if pattern == "" {
		log.Fatal("missing -pattern")
	}

	// Global deadline so demo cannot hang forever.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := patterns.Run(ctx, pattern); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

