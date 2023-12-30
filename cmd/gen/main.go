package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/delaneyj/gostar/cfg"
	"github.com/delaneyj/gostar/generator"
	"github.com/delaneyj/gostar/generator/iconify"
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

	if err := generator.GenerateAll(ctx, "./elements", cfg.Default); err != nil {
		return fmt.Errorf("failed to generate all: %w", err)
	}

	if err := iconify.GenerateIconify(ctx, "./tmp", "./elements"); err != nil {
		return fmt.Errorf("failed to generate iconify: %w", err)
	}

	// pkgs, err := mdn.ScrapePackages(ctx, "./tmp")
	// if err != nil {
	// 	return fmt.Errorf("failed to scrape packages: %w", err)
	// }

	// pkgs, err := idls.GetWebCoreIDLs(ctx, "./tmp")
	// if err != nil {
	// 	return fmt.Errorf("failed to get WebCore IDLs: %w", err)
	// }

	// if err := generator.GenerateAll(ctx, "./elements", pkgs); err != nil {
	// 	return fmt.Errorf("failed to generate all: %w", err)
	// }

	return nil
}
