package main

// City defines the name, list of currently
// occupied aliens and the direction map for
// the outbound cities
type City struct {
	name          string
	direction     map[string]string
	currentAliens []uint
}

// RemoveAlienFromCity will remove the alien from the city's
// currently occupied aliens list
func (c *City) RemoveAlienFromCity(alien uint) {
	for i, currentAlien := range c.currentAliens {
		if currentAlien == alien {
			// Delete the alien
			c.currentAliens[i] = c.currentAliens[len(c.currentAliens)-1]
			c.currentAliens = c.currentAliens[:len(c.currentAliens)-1]

			break
		}
	}
}
