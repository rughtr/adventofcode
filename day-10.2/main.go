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
	id     int
	r      rune
	marker rune
	from   *pipe
	to     *pipe
	d      direction
	pt     point
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

const (
	enclosedTileMarker rune = 'I'
	groundMarker       rune = '.'
	openMarker         rune = '0'
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

type queryable[TValue any] []TValue

func (q queryable[TValue]) firstOrDefault(predicate func(TValue) bool) (TValue, bool) {
	for _, next := range q {
		if predicate(next) {
			return next, true
		}
	}
	var d TValue
	return d, false
}

func (m *pipesMap) getConnectionPoints(p *pipe, pos point) queryable[point] {
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
		neighbor, ok := neighbors.firstOrDefault(func(neighbor point) bool {
			p := m.grid[neighbor.x][neighbor.y]
			return p.from == nil && p.r != 'S'
		})
		if ok {
			p := m.grid[neighbor.x][neighbor.y]
			getDirection := func() direction {
				if pos.x < neighbor.x {
					return south
				}
				if pos.x > neighbor.x {
					return north
				}
				if pos.y < neighbor.y {
					return west
				}
				return east
			}
			p.d = getDirection()
			current.to = p
			if p.from != nil {
				continue
			}
			p.from = current
			stack.Enqueue(neighbor)
		}
	}
}

func (m *pipesMap) countEnclosedTiles() int {

	result := 0

	q := MakeQueue[*pipe]()

	entryPipe := m.grid[m.start.x][m.start.y]
	q.Enqueue(entryPipe)

	set := make(map[int]struct{})

	taint := func(current *pipe, d direction, marker rune, r rune) {
		x := current.pt.x
		y := current.pt.y
		dx := 0
		dy := 0
		switch d {
		case north:
			dx = -1
		case east:
			dy = 1
		case south:
			dx = 1
		case west:
			dy = -1
		}
		if dx != 0 {
			for a := x + dx; dx > 0 && a < len(m.grid) || dx < 0 && a >= 0; a += dx {
				p := m.grid[a][y]
				if p.from != nil && p.d != d {
					return
				} else if p.r == groundMarker {
					p.marker = marker
					p.r = r
				}
			}
		} else if dy != 0 {
			for a := y + dy; dy > 0 && a < len(m.grid[x]) || dy < 0 && a >= 0; a += dy {
				p := m.grid[x][a]
				if p.from != nil && p.d != d {
					return
				} else if p.r == groundMarker {
					p.marker = marker
				}
			}
		}
	}

	for q.Any() {
		next := q.Dequeue()
		_, seen := set[next.id]
		if seen {
			break
		}
		p := next.to
		if p != nil {
			/* traceOuter := func(c *pipe, d direction) {
				switch d {
				case north:
					taint(c, west, openMarker)
				case east:
					taint(c, north, openMarker)
				case south:
					taint(c, east, openMarker)
				case west:
					taint(c, south, openMarker)
				}
			} */
			traceInner := func(c *pipe, d direction) {
				switch d {
				case north:
					taint(c, east, enclosedTileMarker, 'E')
				case east:
					taint(c, south, enclosedTileMarker, 'Z')
				case south:
					taint(c, west, enclosedTileMarker, 'W')
				case west:
					taint(c, north, enclosedTileMarker, 'N')
				}
			}
			enteredFrom := next.d
			dir := p.d
			// traceOuter(next, enteredFrom)
			// traceOuter(p, dir)
			traceInner(next, enteredFrom)
			traceInner(p, dir)
			q.Enqueue(p)
		}
		set[next.id] = struct{}{}
	}

	m.write()

	return result
}

func readPipesMap(filename string) pipesMap {

	file, _ := os.Open(filename)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	row, column := 0, 0
	grid := make([][]*pipe, 0)

	id := 0

	lineNumber := 0
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		runes := []rune(reader.Text())
		numPipes := len(runes)
		pipes := make([]*pipe, numPipes)
		for i, r := range runes {
			id++
			pipes[i] = &pipe{r: r, id: id, pt: point{x: lineNumber, y: i}}
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
	numEnclosedTiles := m.countEnclosedTiles()

	fmt.Printf("%d\n", numEnclosedTiles)
}
