package scraper

import "fmt"

// ScraperInteractor represents the business logic (future RIBs Interactor)
type ScraperInteractor struct{}

func NewInteractor() *ScraperInteractor {
	return &ScraperInteractor{}
}

func (i *ScraperInteractor) ExecuteHello(target string) (string, error) {
	if target == "" {
		return "", fmt.Errorf("target is required")
	}
	// Logic isolated from the transport layer
	return fmt.Sprintf("Scrappy successfully reached: %s", target), nil
}
