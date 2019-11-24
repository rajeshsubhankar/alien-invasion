package main

// City defines the name, list of currently
// occupied aliens and the direction map for
// the outbound cities
type City struct {
	name          string
	direction     map[string]string
	currentAliens []uint
}
