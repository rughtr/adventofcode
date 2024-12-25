package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

type point struct {
	x, y int
}

type pipe struct {
	r        rune
	distance int
	from     *pipe
	to       []*pipe
}

type pipesMap struct {
	grid  [][]*pipe
	start point
}

type direction int

const (
	north direction = iota
	east
	south
	west
)

func (p *pipe) getDirections() []direction {
	switch p.r {
	case '|':
		return []direction{north, south}
	case '-':
		return []direction{west, east}
	case 'L':
		return []direction{north, east}
	case 'J':
		return []direction{north, west}
	case '7':
		return []direction{south, west}
	case 'F':
		return []direction{south, east}
	case 'S':
		return []direction{north, east, south, west}
	}
	return []direction{}
}

func (p *pipe) matchesDirections(neighbor *pipe, d direction) bool {
	neighborDirections := neighbor.getDirections()
	for _, dir := range neighborDirections {
		if d == north && dir == south {
			return true
		} else if d == east && dir == west {
			return true
		} else if d == south && dir == north {
			return true
		} else if d == west && dir == east {
			return true
		}
	}
	return false
}

func (m *pipesMap) getConnectionPoints(p *pipe, pos point) []point {
	neighbors := make([]point, 0)
	collectNeighbor := func(x, y int, d direction) {
		if x >= 0 && y >= 0 && x < len(m.grid) && y < len(m.grid[x]) {
			other := m.grid[x][y]
			if p.matchesDirections(other, d) {
				neighbors = append(neighbors, point{x: x, y: y})
			}
		}
	}
	dirList := p.getDirections()
	for _, d := range dirList {
		switch d {
		case north:
			collectNeighbor(pos.x-1, pos.y, d)
		case east:
			collectNeighbor(pos.x, pos.y+1, d)
		case south:
			collectNeighbor(pos.x+1, pos.y, d)
		case west:
			collectNeighbor(pos.x, pos.y-1, d)
		}
	}
	return neighbors
}

func (m *pipesMap) wirePipes() {
	stack := MakeQueue[point]()
	stack.Enqueue(m.start)
	for stack.Any() {
		pos := stack.Dequeue()
		current := m.grid[pos.x][pos.y]
		neighbors := m.getConnectionPoints(current, pos)
		for _, neighbor := range neighbors {
			p := m.grid[neighbor.x][neighbor.y]
			if p.from != nil || p.r == 'S' {
				continue
			}
			current.to = append(current.to, p)
			distance := current.distance + 1
			if distance < p.distance || p.distance == 0 { /// ???
				p.distance = distance
				p.from = current
			}
			stack.Enqueue(neighbor)
		}
	}
}

func (m *pipesMap) farthest() *pipe {
	var farthest = m.grid[m.start.x][m.start.y]
	stack := MakeStack[*pipe]()
	stack.Push(farthest)
	for stack.Any() {
		next := stack.Pop()
		for _, neighbor := range next.to {
			if farthest.distance < neighbor.distance {
				farthest = neighbor
			}
			stack.Push(neighbor)
		}
	}
	return farthest
}

func readPipesMap(filename string) pipesMap {

	file, _ := os.Open(filename)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	row, column := 0, 0
	grid := make([][]*pipe, 0)

	lineNumber := 0
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		runes := []rune(reader.Text())
		numPipes := len(runes)
		pipes := make([]*pipe, numPipes)
		for i, r := range runes {
			pipes[i] = &pipe{r: r, distance: 0, to: make([]*pipe, 0)}
			if r == 'S' {
				row = lineNumber
				column = i
			}
		}
		grid = append(grid, pipes)
		lineNumber++
	}

	m := pipesMap{
		grid: grid,
		start: point{
			x: row,
			y: column,
		},
	}

	m.wirePipes()

	return m
}

func main() {

	cwd, _ := os.Getwd()
	inputPath := path.Join(cwd, "input.txt")
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		panic("the input file does not exist")
	}

	m := readPipesMap(inputPath)
	farthest := m.farthest()

	fmt.Printf("%d\n", farthest.distance)
}
