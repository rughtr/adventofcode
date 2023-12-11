package main

type waypoint struct {
	name  string
	left  string
	right string
}

func createWaypoint(name string, left string, right string) waypoint {
	return waypoint{
		name:  name,
		left:  left,
		right: right,
	}
}
