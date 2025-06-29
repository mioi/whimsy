package whimsy

import (
	"strings"
	"testing"
)

func validateNameList(t *testing.T, names []string, nameType string) {
	if len(names) < 200 {
		t.Errorf("Expected at least 200 %s names, got %d", nameType, len(names))
	}

	for _, name := range names {
		if len(name) > 6 {
			t.Errorf("%s name '%s' is longer than 6 characters", nameType, name)
		}
		for _, char := range name {
			if char < 'a' || char > 'z' {
				t.Errorf("%s name '%s' contains invalid character '%c'", nameType, name, char)
			}
		}
	}

	// Check alphabetical order
	for i := 1; i < len(names); i++ {
		if names[i-1] >= names[i] {
			t.Errorf("%s names are not in alphabetical order: '%s' should come after '%s'", nameType, names[i-1], names[i])
		}
	}

	// Check for duplicates
	for i := 1; i < len(names); i++ {
		if names[i-1] == names[i] {
			t.Errorf("%s names have duplicates: '%s' found multiple times", nameType, names[i-1])
		}
	}
}

func TestPlantNames(t *testing.T) {
	validateNameList(t, Plants(), "Plant")
}

func TestAnimalNames(t *testing.T) {
	validateNameList(t, Animals(), "Animal")
}

func TestColorNames(t *testing.T) {
	validateNameList(t, Colors(), "Color")
}

func TestRandomPlant(t *testing.T) {
	plant, err := RandomPlant()
	if err != nil {
		t.Errorf("RandomPlant() returned error: %v", err)
	}
	if plant == "" {
		t.Error("RandomPlant() returned empty string")
	}

	// Verify it's a valid plant name
	found := false
	for _, p := range Plants() {
		if p == plant {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("RandomPlant() returned '%s' which is not in the plants list", plant)
	}
}

func TestRandomAnimal(t *testing.T) {
	animal, err := RandomAnimal()
	if err != nil {
		t.Errorf("RandomAnimal() returned error: %v", err)
	}
	if animal == "" {
		t.Error("RandomAnimal() returned empty string")
	}

	// Verify it's a valid animal name
	found := false
	for _, a := range Animals() {
		if a == animal {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("RandomAnimal() returned '%s' which is not in the animals list", animal)
	}
}

func TestRandomColor(t *testing.T) {
	color, err := RandomColor()
	if err != nil {
		t.Errorf("RandomColor() returned error: %v", err)
	}
	if color == "" {
		t.Error("RandomColor() returned empty string")
	}

	// Verify it's a valid color name
	found := false
	for _, c := range Colors() {
		if c == color {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("RandomColor() returned '%s' which is not in the colors list", color)
	}
}

func TestRandomName(t *testing.T) {
	// Test default (2 parts)
	name, err := RandomName()
	if err != nil {
		t.Errorf("RandomName() returned error: %v", err)
	}
	if name == "" {
		t.Error("RandomName() returned empty string")
	}

	parts := strings.Split(name, "-")
	if len(parts) != 2 {
		t.Errorf("RandomName() default should return 2 parts, got %d: %s", len(parts), name)
	}

	// Test 1 part
	name1, err := RandomName(1)
	if err != nil {
		t.Errorf("RandomName(1) returned error: %v", err)
	}
	parts1 := strings.Split(name1, "-")
	if len(parts1) != 1 {
		t.Errorf("RandomName(1) should return 1 part, got %d: %s", len(parts1), name1)
	}

	// Test 3 parts
	name3, err := RandomName(3)
	if err != nil {
		t.Errorf("RandomName(3) returned error: %v", err)
	}
	parts3 := strings.Split(name3, "-")
	if len(parts3) != 3 {
		t.Errorf("RandomName(3) should return 3 parts, got %d: %s", len(parts3), name3)
	}

	// Test max parts (dynamic based on categories)
	maxParts := len(Categories())
	nameMax, err := RandomName(maxParts)
	if err != nil {
		t.Errorf("RandomName(%d) returned error: %v", maxParts, err)
	}
	partsMax := strings.Split(nameMax, "-")
	if len(partsMax) != maxParts {
		t.Errorf("RandomName(%d) should return %d parts, got %d: %s", maxParts, maxParts, len(partsMax), nameMax)
	}

	// Test all parts are valid and unique
	allWords := make(map[string]bool)
	for _, cat := range Categories() {
		for _, word := range cat.Words {
			allWords[word] = true
		}
	}

	for _, part := range partsMax {
		if !allWords[part] {
			t.Errorf("RandomName(%d) returned invalid part '%s'", maxParts, part)
		}
	}

	// Check for duplicates in max-part name
	partMap := make(map[string]bool)
	for _, part := range partsMax {
		if partMap[part] {
			t.Errorf("RandomName(%d) returned duplicate part '%s' in %s", maxParts, part, nameMax)
		}
		partMap[part] = true
	}
}

func TestRandomNameInvalidCount(t *testing.T) {
	maxParts := len(Categories())

	// Test invalid counts
	_, err := RandomName(0)
	if err == nil {
		t.Error("RandomName(0) should return error")
	}

	_, err = RandomName(maxParts + 1)
	if err == nil {
		t.Errorf("RandomName(%d) should return error (max is %d)", maxParts+1, maxParts)
	}

	_, err = RandomName(-1)
	if err == nil {
		t.Error("RandomName(-1) should return error")
	}
}

func TestRandomNameUniqueness(t *testing.T) {
	// Test that we get different names (not requiring all to be different due to randomness)
	names := make(map[string]bool)
	for i := 0; i < 20; i++ {
		name, err := RandomName(2)
		if err != nil {
			t.Errorf("RandomName(2) returned error: %v", err)
		}
		names[name] = true
	}

	// We should get at least several different names
	if len(names) < 10 {
		t.Errorf("Expected at least 10 different names in 20 attempts, got %d", len(names))
	}
}

func TestCategories(t *testing.T) {
	categories := Categories()

	// Should have at least the 3 default categories
	if len(categories) < 3 {
		t.Errorf("Expected at least 3 categories, got %d", len(categories))
	}

	// Check that we have the expected categories
	expectedCategories := map[string]bool{
		"plants":  false,
		"animals": false,
		"colors":  false,
	}

	for _, cat := range categories {
		if _, exists := expectedCategories[cat.Name]; exists {
			expectedCategories[cat.Name] = true
		}

		// Each category should have words
		if len(cat.Words) == 0 {
			t.Errorf("Category '%s' has no words", cat.Name)
		}
	}

	// Verify all expected categories were found
	for catName, found := range expectedCategories {
		if !found {
			t.Errorf("Expected category '%s' not found", catName)
		}
	}
}

// Benchmarks
func BenchmarkRandomPlant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := RandomPlant()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRandomAnimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := RandomAnimal()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRandomColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := RandomColor()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRandomName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := RandomName(2)
		if err != nil {
			b.Fatal(err)
		}
	}
}
