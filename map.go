package main

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Map defines the map of cities and aliens
type Map struct {
	cities map[string]*City
	aliens map[uint]*Alien
}

// Create an empty map
func newEmptyMap() *Map {
	return &Map{
		cities: make(map[string]*City),
		aliens: make(map[uint]*Alien),
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

// SpreadAliens creates n number of new aliens and
// send them to random cities
// NOTE: A city can have more than 2 aliens
func (m *Map) SpreadAliens(numberOfAliens uint) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	aliens := make(map[uint]*Alien)

	for i := uint(0); i < numberOfAliens; i++ {
		// Select a random city to send this alien
		c := m.RandomCity(r)
		a := &Alien{
			name:        i,
			currentCity: c.name,
		}
		c.currentAliens = append(c.currentAliens, a.name)
		// Update aliens map
		aliens[a.name] = a
	}
	m.aliens = aliens
}

// RandomCity will return any pseudo-random city from the map
func (m *Map) RandomCity(r *rand.Rand) *City {
	var cities []string
	for city := range m.cities {
		cities = append(cities, city)
	}

	randomCity := cities[r.Intn(len(cities))]

	return m.cities[randomCity]
}
