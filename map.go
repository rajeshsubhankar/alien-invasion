package main

import (
	"bufio"
	"os"
	"strings"
)

// Map defines the map of cities and aliens
type Map struct {
	cities map[string]*City
	aliens map[int]*Alien
}

// Create an empty map
func newEmptyMap() *Map {
	return &Map{
		cities: make(map[string]*City),
		aliens: make(map[int]*Alien),
	}
}

// Create a new map from a file
func newMapFromFile(fileName string) (*Map, error) {
	m := newEmptyMap()

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		// Create a new city
		c := &City{
			name: tokens[0],
		}
		c.direction = make(map[string]string)

		// If the city has at least one neighbouring city
		if len(tokens) > 1 {
			for _, path := range tokens[1:] {
				dirAndDest := strings.Split(path, "=")

				// Update the city's direction map with the neighbouring city
				c.direction[dirAndDest[0]] = dirAndDest[1]

				// If the neighbouring city doesn't exist in the map, create one
				if m.cities[dirAndDest[1]] == nil {
					m.cities[dirAndDest[1]] = &City{
						name: dirAndDest[1],
					}
				}
			}
		}

		// Add the brand new city to the map
		m.cities[c.name] = c
	}

	return m, nil
}
