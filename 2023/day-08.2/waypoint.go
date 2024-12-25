package main

import "strings"

type waypoint struct {
	name          string
	left          string
	leftWaypoint  *waypoint
	right         string
	rightWaypoint *waypoint
	isEnd         bool
}

func (p *waypoint) link(m map[string]*waypoint) {
	leftWp := m[p.left]
	p.leftWaypoint = leftWp
	rightWp := m[p.right]
	p.rightWaypoint = rightWp

}

func createWaypoint(name string, left string, right string) *waypoint {
	z := strings.HasSuffix(name, "Z")
	return &waypoint{
		name:  name,
		left:  left,
		right: right,
		isEnd: z,
	}
}
