package area51

import (
	"os"
	"testing"
)

var (
	testFile        = "map_test.txt"
	numberOfAliens  = uint(5)
	maxSteps        = uint(10000)
	totalCities     = 5
	totalDirections = 5
)

func createTestFile() error {
	f, err := os.Create(testFile)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString("Foo north=Bar west=Baz south=Qu-ux\n")
	if err != nil {
		return err
	}
	_, err = f.WriteString("Bar south=Foo west=Bee\n")
	if err != nil {
		return err
	}
	return f.Sync()
}

func deleteTestFile() error {
	return os.Remove(testFile)
}

func TestNewMapFromFile(t *testing.T) {
	_, err := NewMapFromFile("wrongfile")
	if err == nil {
		t.Errorf("Expected to receive error")
	}

	err = createTestFile()
	if err != nil {
		t.Errorf("Unable to create a test file")
	}

	m, err := NewMapFromFile(testFile)
	if err != nil {
		t.Errorf("Unable to load map from the file")
	}

	if len(m.cities) != totalCities {
		t.Errorf("Incorrect total number of cities, Expected: %v, Received: %v", totalCities, len(m.cities))
	}

	n := 0
	for _, c := range m.cities {
		n += len(c.direction)
	}
	if n != totalDirections {
		t.Errorf("Incorrect total number of neighbouring cities, Expected: %v, Received: %v", totalDirections, n)
	}

	err = deleteTestFile()
	if err != nil {
		t.Errorf("Unable to delete test file")
	}
}

func TestSpreadAliens(t *testing.T) {
	err := createTestFile()
	if err != nil {
		t.Errorf("Unable to create a test file")
	}

	m, err := NewMapFromFile(testFile)
	if err != nil {
		t.Errorf("Unable to load map from the file")
	}

	m.SpreadAliens(numberOfAliens)
	if uint(len(m.aliens)) != numberOfAliens {
		t.Errorf("Incorrect total aliens, Expected: %v , Received: %v", numberOfAliens, len(m.aliens))
	}

	totalAliens := uint(0)
	for _, c := range m.cities {
		totalAliens += uint(len(c.currentAliens))
	}
	if totalAliens != numberOfAliens {
		t.Errorf("Incorrect total aliens from all the cities, Expected: %v , Received: %v", numberOfAliens, totalAliens)
	}

	err = deleteTestFile()
	if err != nil {
		t.Errorf("Unable to delete test file")
	}
}

func TestInvade(t *testing.T) {
	err := createTestFile()
	if err != nil {
		t.Errorf("Unable to create a test file")
	}

	m, err := NewMapFromFile(testFile)
	if err != nil {
		t.Errorf("Unable to load map from the file")
	}

	m.SpreadAliens(numberOfAliens)
	m.Invade(maxSteps)

	for _, c := range m.cities {
		if len(c.currentAliens) > 1 {
			t.Errorf("City %v with %v aliens is not yet destroyed", c.name, len(c.currentAliens))
		}
	}

	err = deleteTestFile()
	if err != nil {
		t.Errorf("Unable to delete test file")
	}
}

func TestMoveAlien(t *testing.T) {
	err := createTestFile()
	if err != nil {
		t.Errorf("Unable to create a test file")
	}

	m, err := NewMapFromFile(testFile)
	if err != nil {
		t.Errorf("Unable to load map from the file")
	}

	m.SpreadAliens(numberOfAliens)

	a := m.aliens[0]
	currentCity := a.currentCity
	m.moveAlien(a)

	if currentCity == a.currentCity && len(m.cities[a.currentCity].direction) != 0 {
		t.Errorf("Alien can not move to the same city %v", a.currentCity)
	}

	err = deleteTestFile()
	if err != nil {
		t.Errorf("Unable to delete test file")
	}
}

func TestDeepClean(t *testing.T) {
	err := createTestFile()
	if err != nil {
		t.Errorf("Unable to create a test file")
	}

	m, err := NewMapFromFile(testFile)
	if err != nil {
		t.Errorf("Unable to load map from the file")
	}

	m.SpreadAliens(numberOfAliens)
	m.Invade(maxSteps)

	for _, c := range m.cities {
		for _, dest := range c.direction {
			if len(m.cities[dest].currentAliens) > 1 {
				t.Errorf("The dead city %v still exist as an outward connection to %v", dest, c.name)
			}
		}
	}

	err = deleteTestFile()
	if err != nil {
		t.Errorf("Unable to delete test file")
	}
}

func TestRemoveAliensFromMap(t *testing.T) {
	err := createTestFile()
	if err != nil {
		t.Errorf("Unable to create a test file")
	}

	m, err := NewMapFromFile(testFile)
	if err != nil {
		t.Errorf("Unable to load map from the file")
	}

	m.SpreadAliens(numberOfAliens)
	// Remove alien '0'
	m.removeAliensFromMap([]uint{0})
	if m.aliens[0] != nil {
		t.Errorf("Alien %v is still alive", 0)
	}

	err = deleteTestFile()
	if err != nil {
		t.Errorf("Unable to delete test file")
	}
}
