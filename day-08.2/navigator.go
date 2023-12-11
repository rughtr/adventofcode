package main

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strings"
)

type hauntedWastelandMap struct {
	directions string
	points     map[string]*waypoint
	paths      []*mapPath
}

type mapPath struct {
	steps int64
	p     int
	wp    *waypoint
}

func createWastelandMapNavigator(directions string, points map[string]*waypoint) *hauntedWastelandMap {
	paths := make([]*mapPath, 0)
	for _, wp := range points {
		if strings.HasSuffix(wp.name, "A") {
			p := &mapPath{steps: 0, p: -1, wp: wp}
			paths = append(paths, p)
		}
	}
	numPaths := len(paths)
	if numPaths > 0 {
		return &hauntedWastelandMap{
			directions: directions,
			points:     points,
			paths:      paths,
		}
	}
	return nil
}

func (m *hauntedWastelandMap) calculateSteps() int64 {

	steps := make([]int64, 0)
	for _, p := range m.paths {
		steps = append(steps, p.steps)
	}

	calculateGreatestCommonDivisor := func(a, b int64) int64 {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	calculateLeastCommonMultiple := func(n ...int64) int64 {
		lcm := func(a, b int64) int64 {
			ab := a * b
			gcd := calculateGreatestCommonDivisor(a, b)
			return int64(math.Abs(float64(ab) / float64(gcd)))
		}
		result := n[0]
		for i := 1; i < len(n); i++ {
			result = lcm(result, n[i])
		}
		return result
	}

	return calculateLeastCommonMultiple(steps...)
}

func (m *hauntedWastelandMap) NavigateNext() bool {
	if m.hasMoreWaypoints() == false {
		return false
	}
	m.updateWaypoints()
	return true
}

func (m *hauntedWastelandMap) hasMoreWaypoints() bool {
	length := len(m.paths)
	if length == 0 {
		return false
	}
	for i := 0; i < length; i++ {
		p := m.paths[i]
		if p.wp.isEnd == false {
			return true
		}
	}
	return false
}

func (m *hauntedWastelandMap) updateWaypoints() {
	length := len(m.paths)
	for i := 0; i < length; i++ {
		currentPath := m.paths[i]
		if currentPath.wp.isEnd {
			continue // this mapPath is complete
		}
		p, d := m.nextDirection(currentPath)
		current := currentPath.wp
		wp := current.leftWaypoint
		if d == 'R' {
			wp = current.rightWaypoint
		}
		currentPath.wp = wp
		currentPath.p = p
		currentPath.steps += 1
	}
}

func (m *hauntedWastelandMap) nextDirection(p *mapPath) (int, rune) {
	n := p.p + 1
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

	m := make(map[string]*waypoint)

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

	for _, next := range m {
		next.link(m)
	}

	return createWastelandMapNavigator(directions, m)
}
