package main

import "fmt"

// write Renders a visually more appealing and recognizable map version for debugging purposes.
func (m *pipesMap) write() {
	for x := 0; x < len(m.grid); x++ {
		for y := 0; y < len(m.grid[x]); y++ {
			p := m.grid[x][y]
			if p.from != nil {
				switch p.r {
				case '|':
					fmt.Printf("%c", '\u2551')
				case '-':
					fmt.Printf("%c", '\u2550')
				case 'L':
					fmt.Printf("%c", '\u255a')
				case 'J':
					fmt.Printf("%c", '\u255d')
				case '7':
					fmt.Printf("%c", '\u2557')
				case 'F':
					fmt.Printf("%c", '\u2554')
				default:
					fmt.Printf("%s", string(p.marker))
				}
				continue
			}
			switch p.r {
			case groundMarker:

				switch p.marker {
				case enclosedTileMarker:
					fmt.Printf("I")
				case openMarker:
					fmt.Printf("0")
				default:
					fmt.Printf(string(p.r))
				}

			case '|', '-', 'L', 'J', '7', 'F':
				fmt.Printf(" ") // do not print waste tiles

			default:

				// fmt.Printf(" ")
				fmt.Printf("%s", string(p.r))
			}
		}
		fmt.Printf("\n")
	}
}
