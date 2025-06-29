# Whimsy

[![Tests](https://github.com/mioi/whimsy/actions/workflows/test.yml/badge.svg)](https://github.com/mioi/whimsy/actions/workflows/test.yml)
[![Coverage](https://codecov.io/gh/mioi/whimsy/branch/main/graph/badge.svg)](https://codecov.io/gh/mioi/whimsy)
[![Go Version](https://img.shields.io/badge/go-1.21-blue.svg)](https://golang.org)

A Go library that generates random memorable names using plants, animals, and colors. Perfect for infrastructure naming.

## Installation

```bash
go get github.com/mioi/whimsy
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/mioi/whimsy"
)

func main() {
    // Individual categories
    plant, _ := whimsy.RandomPlant()   // "oak"
    animal, _ := whimsy.RandomAnimal() // "fox"
    color, _ := whimsy.RandomColor()   // "blue"
    
    // Combined names (1-3 parts)
    name1, _ := whimsy.RandomName(1)   // "fox"
    name2, _ := whimsy.RandomName(2)   // "blue-fox" 
    name3, _ := whimsy.RandomName(3)   // "blue-fox-oak"
    name, _ := whimsy.RandomName()     // "red-pine" (default: 2 parts)
}
```

## API

- `RandomPlant() (string, error)` - Random plant name
- `RandomAnimal() (string, error)` - Random animal name  
- `RandomColor() (string, error)` - Random color name
- `RandomName(count ...int) (string, error)` - 1-N random words from any category
- `Plants() []string` - All plant names
- `Animals() []string` - All animal names
- `Colors() []string` - All color names
- `Categories() []Category` - All categories with metadata

## Constraints

- 200+ names per category
- Max 6 characters, lowercase a-z only
- Alphabetically sorted, no duplicates
- Cryptographically secure random generation

## Extensibility

Easy to add new categories:

```go
// 1. Add word list to names.go
var weather = []string{"cloud", "rain", "snow", "sun", "wind", "fog"}

// 2. Update getAllCategories()
func getAllCategories() []Category {
    return []Category{
        {"plants", plants},
        {"animals", animals}, 
        {"colors", colors},
        {"weather", weather}, // Add this line
    }
}
```

`RandomName()` automatically supports new categories with no code changes needed.

## License

MIT