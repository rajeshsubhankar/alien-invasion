package main

import (
	"bufio"
	"fmt"
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

// Invade will simulate the alien invasion and delete
// the cities and aliens when they are destroyed/dead
func (m *Map) Invade(maxMoves uint) {
	// move each alien at max of 'maxMoves' times
	for i := uint(0); i < maxMoves; i++ {
		// If all the cities are destroyed, stop
		if m.cities == nil {
			break
		}

		// If all the aliens are dead, stop
		if m.aliens == nil {
			break
		}

		// Move each alien by one step
		for _, a := range m.aliens {
			m.MoveAlien(a)
		}

		// Delete the destroyed cities and dead aliens
		m.CleanUp()
	}
}

// MoveAlien will attempt to move the alien from the currently
// occupied city to the next available neighbouring city
func (m *Map) MoveAlien(a *Alien) {
	if c, ok := m.cities[a.currentCity]; ok {
		// Move in the first available direction
		// NOTE: It will be counted as a valid move even if there
		// is no valid path available (since the move was attempted)
		for _, dest := range c.direction {
			// Update neighbouring city's aliens list
			destCity := m.cities[dest]
			destCity.currentAliens = append(destCity.currentAliens, a.name)

			// Remove the alien from the current city
			c.RemoveAlienFromCity(a.name)

			// Update alien's current city
			a.currentCity = dest

			break
		}
	}
}

// CleanUp will try to clean the map based on
// the alien invasion rule set
func (m *Map) CleanUp() {
	// If a city is currently occupied by >= 2 aliens, destroy it
	for cityName, c := range m.cities {
		if len(c.currentAliens) >= 2 {
			// Remove all the links which are pointing to this city
			m.DeepClean(cityName)

			// Print fight message
			// @dev Update formatting
			fmt.Println(cityName, "has been destroyed by aliens ", c.currentAliens, "!")

			// Remove the dead aliens
			m.RemoveAliensFromMap(c.currentAliens)

			// Finally remove the dead city
			delete(m.cities, cityName)
		}
	}
}

// DeepClean will erase the city from all other neighbouring
// city if they have any outward connection towards this city
func (m *Map) DeepClean(cityName string) {
	for _, c := range m.cities {
		for dir, dest := range c.direction {
			if dest == cityName {
				delete(c.direction, dir)
			}
		}
	}
}

// RemoveAliensFromMap will erase all the dead aliens from the map
func (m *Map) RemoveAliensFromMap(a []uint) {
	for _, alien := range a {
		delete(m.aliens, alien)
	}
}
