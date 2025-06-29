// Package whimsy provides random memorable names using plants, animals, and colors.
// Perfect for creating human-friendly names for infrastructure resources while
// maintaining the "pets vs. cattle" principle.
package whimsy

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Plants returns a slice of all available plant names
func Plants() []string {
	return plants
}

// Animals returns a slice of all available animal names
func Animals() []string {
	return animals
}

// Colors returns a slice of all available color names
func Colors() []string {
	return colors
}

// Categories returns information about all available word categories
func Categories() []Category {
	return getAllCategories()
}

// RandomPlant returns a random plant name
func RandomPlant() (string, error) {
	return randomFromSlice(plants)
}

// RandomAnimal returns a random animal name
func RandomAnimal() (string, error) {
	return randomFromSlice(animals)
}

// RandomColor returns a random color name
func RandomColor() (string, error) {
	return randomFromSlice(colors)
}

// Category represents a category of words with its name and word list
type Category struct {
	Name  string
	Words []string
}

// getAllCategories returns all available word categories
func getAllCategories() []Category {
	return []Category{
		{"plants", plants},
		{"animals", animals},
		{"colors", colors},
	}
}

// getAllWords returns all words from all categories combined
func getAllWords() []string {
	categories := getAllCategories()
	totalWords := 0
	for _, cat := range categories {
		totalWords += len(cat.Words)
	}

	allWords := make([]string, 0, totalWords)
	for _, cat := range categories {
		allWords = append(allWords, cat.Words...)
	}
	return allWords
}

// RandomName returns a random name with 1-N parts from any combination of available word categories,
// where N is the number of available categories. If count is not provided, defaults to 2.
// Examples: "blue", "fox-oak", "red-wolf-pine"
func RandomName(count ...int) (string, error) {
	maxParts := len(getAllCategories())

	// Default to 2 parts if not specified
	numParts := 2
	if len(count) > 0 {
		numParts = count[0]
		if numParts < 1 || numParts > maxParts {
			return "", fmt.Errorf("count must be between 1 and %d, got %d", maxParts, numParts)
		}
	}

	allWords := getAllWords()
	if len(allWords) == 0 {
		return "", fmt.Errorf("no words available")
	}

	// Generate unique random parts
	parts := make([]string, 0, numParts)
	used := make(map[string]bool)

	for len(parts) < numParts {
		word, err := randomFromSlice(allWords)
		if err != nil {
			return "", fmt.Errorf("failed to generate random word: %w", err)
		}

		// Ensure we don't repeat words
		if !used[word] {
			parts = append(parts, word)
			used[word] = true
		}
	}

	// Join with hyphens
	name := ""
	for i, part := range parts {
		if i > 0 {
			name += "-"
		}
		name += part
	}

	return name, nil
}

// randomFromSlice returns a random element from the given slice using crypto/rand
func randomFromSlice(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", fmt.Errorf("slice is empty")
	}

	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(slice))))
	if err != nil {
		return "", fmt.Errorf("failed to generate random number: %w", err)
	}

	return slice[index.Int64()], nil
}
