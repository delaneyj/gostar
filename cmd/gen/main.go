package main

import (
	"context"
	"log"
	"time"

	"github.com/delaneyj/gostar/generator"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	ctx := context.Background()

	start := time.Now()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}

	log.Printf("took %s", time.Since(start))
}

func run(ctx context.Context) error {

	if err := generator.GenerateAllElements(ctx, &generator.GenerateElementArgs{
		OutputPath: "./html",
		TempDir:    "./tmp",
	}); err != nil {
		return err
	}
	return nil
}
