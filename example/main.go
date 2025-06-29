package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/mioi/whimsy"
)

func main() {
	fmt.Println("🎲 Whimsy - Random Memorable Names")
	fmt.Println("==================================")

	// Generate individual names
	fmt.Println("\n📝 Individual Names:")

	plant, err := whimsy.RandomPlant()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("🌱 Plant:  %s\n", plant)

	animal, err := whimsy.RandomAnimal()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("🐾 Animal: %s\n", animal)

	color, err := whimsy.RandomColor()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("🎨 Color:  %s\n", color)

	// Generate random names with different lengths
	fmt.Println("\n🎯 Random Names:")

	name1, _ := whimsy.RandomName(1)
	fmt.Printf("1 part:  %s\n", name1)

	name2, _ := whimsy.RandomName(2)
	fmt.Printf("2 parts: %s\n", name2)

	name3, _ := whimsy.RandomName(3)
	fmt.Printf("3 parts: %s\n", name3)

	maxParts := len(whimsy.Categories())
	nameMax, _ := whimsy.RandomName(maxParts)
	fmt.Printf("%d parts: %s\n", maxParts, nameMax)

	defaultName, _ := whimsy.RandomName()
	fmt.Printf("Default: %s\n", defaultName)

	// Show infrastructure naming examples
	fmt.Println("\n🏗️  Infrastructure Examples:")

	serverName, _ := whimsy.RandomName(2)
	fmt.Printf("Server:      %s.example.com\n", serverName)

	dbName, _ := whimsy.RandomName(3)
	fmt.Printf("Database:    %s-db\n", dbName)

	clusterName, _ := whimsy.RandomName(1)
	fmt.Printf("Cluster:     %s-cluster\n", clusterName)

	// Show statistics
	fmt.Println("\n📊 Statistics:")
	categories := whimsy.Categories()
	totalWords := 0

	for _, cat := range categories {
		fmt.Printf("%-8s: %d names\n", strings.Title(cat.Name), len(cat.Words))
		totalWords += len(cat.Words)
	}

	fmt.Printf("Total words: %d\n", totalWords)
	fmt.Printf("Categories:  %d\n", len(categories))

	// Show some sample names from each category
	fmt.Println("\n🔍 Sample Names:")
	for _, cat := range categories {
		if len(cat.Words) >= 5 {
			fmt.Printf("%-8s: %s, %s, %s, %s, %s\n",
				strings.Title(cat.Name),
				cat.Words[0],
				cat.Words[len(cat.Words)/4],
				cat.Words[len(cat.Words)/2],
				cat.Words[3*len(cat.Words)/4],
				cat.Words[len(cat.Words)-1])
		}
	}
}
