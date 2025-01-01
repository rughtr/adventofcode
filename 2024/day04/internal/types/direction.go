package types

type Direction struct {
	dx int
	dy int
}

func (d Direction) X() int {
	return d.dx
}

func (d Direction) Y() int {
	return d.dy
}

var (
	HorizontalForward  = Direction{1, 0}
	HorizontalReverse  = Direction{-1, 0}
	VerticalDown       = Direction{0, 1}
	VerticalUp         = Direction{0, -1}
	DiagonalLtr        = Direction{1, 1}
	DiagonalLtrReverse = Direction{1, -1}
	DiagonalRtl        = Direction{-1, 1}
	DiagonalRtlReverse = Direction{-1, -1}
)
