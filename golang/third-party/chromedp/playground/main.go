package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// Create a context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Create a timeout context
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second) // Increased timeout for navigation and screenshot
	defer cancel()

	// Connect to the remote browser
	allocatorContext, cancel := chromedp.NewRemoteAllocator(ctx, "ws://127.0.0.1:9222")
	defer cancel()

	taskCtx, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	// Run tasks
	var buf []byte
	url := "https://jp.mercari.com" // Target URL for screenshot

	err := chromedp.Run(taskCtx,
		chromedp.Navigate(url),
		chromedp.FullScreenshot(&buf, 90), // 90% quality
	)
	if err != nil {
		log.Fatalf("Failed to run tasks: %v", err)
	}

	// Save the screenshot to a file
	if err := os.WriteFile("screenshot.png", buf, 0644); err != nil {
		log.Fatalf("Failed to write screenshot: %v", err)
	}

	log.Printf("Successfully navigated to %s and saved screenshot.png", url)
}
