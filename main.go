package main

import (
	"flag"
	"log"
	"os"

	"github.com/rajeshsubhankar/alien-invasion/area51"
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
	m, err := area51.NewMapFromFile(fileName)
	if err != nil {
		log.Fatalln("Error: Unable to create a map from the file.", err)
		os.Exit(1)
	}

	// Create and spread aliens
	m.SpreadAliens(numOfAliens)

	// Let the aliens invade the map
	m.Invade(maxMoves)

	// Print result
	m.Print()
}
