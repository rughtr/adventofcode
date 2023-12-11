package main

import (
	"bufio"
	"os"
	"regexp"
)

type hauntedWastelandMap struct {
	directions string
	p          int
	points     map[string]waypoint
	current    waypoint
}

const start = "AAA"
const end = "ZZZ"

func createWastelandMapNavigator(directions string, points map[string]waypoint) *hauntedWastelandMap {
	startingPoint, ok := points[start]
	if ok {
		return &hauntedWastelandMap{
			p:          -1,
			directions: directions,
			points:     points,
			current:    startingPoint,
		}
	}
	return nil
}

func (m *hauntedWastelandMap) NavigateNext() bool {
	if m.current.name == end {
		return false
	}
	p, nextWaypointName := m.nextWaypointName()
	nextWaypoint, ok := m.points[nextWaypointName]
	if ok {
		m.update(p, nextWaypoint)
		return true
	}
	return false
}

func (m *hauntedWastelandMap) update(p int, wp waypoint) {
	m.p = p
	m.current = wp
}

func (m *hauntedWastelandMap) nextWaypointName() (int, string) {
	p, d := m.nextDirection()
	current := m.current
	if d == 'R' {
		return p, current.right
	}
	return p, current.left
}

func (m *hauntedWastelandMap) nextDirection() (int, rune) {
	n := m.p + 1
	if n >= len(m.directions) {
		n = 0
	}
	return n, rune(m.directions[n])
}

func createNavigator(filename string) *hauntedWastelandMap {

	file, _ := os.Open(filename)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reader := bufio.NewScanner(file)
	reader.Scan()
	directions := reader.Text()

	r := regexp.MustCompile("^(?P<pt>\\w+)\\s*=\\s*\\((?P<left>\\w+),\\s*(?P<right>\\w+)\\)$")
	waypointNameIndex := r.SubexpIndex("pt")
	leftIndex := r.SubexpIndex("left")
	rightIndex := r.SubexpIndex("right")

	m := make(map[string]waypoint)

	for reader.Scan() {
		line := reader.Text()
		if len(line) == 0 {
			continue
		}
		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			name := match[waypointNameIndex]
			m[name] = createWaypoint(name, match[leftIndex], match[rightIndex])
		}
	}

	return createWastelandMapNavigator(directions, m)
}
