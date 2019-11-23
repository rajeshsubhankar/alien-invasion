package main

import (
	"flag"
)

var (
	fileName    string
	numOfAliens uint
	maxMoves    uint
)

func parseArguments() {
	flag.StringVar(&fileName, "f", "map.txt", "File containing world map")
	flag.UintVar(&numOfAliens, "n", 5, "Total number of aliens")
	flag.UintVar(&maxMoves, "max", 10000, "Maximum number of moves per alien")

	flag.Parse()
}

func main() {
	// Parse command line arguments
	parseArguments()

	// Create a new map from the file
	// Create and spread aliens
	// Let the aliens invade the map
	// Print result
}
