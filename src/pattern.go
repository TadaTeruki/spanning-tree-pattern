package main

import (
	"image/color"
	"math/rand"

	pq "github.com/TadaTeruki/PriorityQueueGo/PriorityQueue"
)

var direction = [][]int{
	{0, -1}, {0, 1}, {-1, 0}, {1, 0},
}

type pixel struct {
	parent_x, parent_y, x, y int
}

func clamp(t float64, min float64, max float64) float64 {
	if t < min {
		return min
	} else if t > max {
		return max
	} else {
		return t
	}
}

func modify(k int, k_range int, prop float64) int {
	return (k + int(float64(k_range)*prop)) % k_range
}

func mapColorUnit(m float64, a, b uint8) uint8 {
	return uint8(clamp(float64(a)*m+float64(b)*(1.-m), 0., 255.))
}

func mapColor(m float64, ca, cb *color.RGBA) *color.RGBA {
	nc := new(color.RGBA)
	nc.R = mapColorUnit(m, ca.R, cb.R)
	nc.G = mapColorUnit(m, ca.G, cb.G)
	nc.B = mapColorUnit(m, ca.B, cb.B)
	nc.A = mapColorUnit(m, ca.A, cb.A)
	return nc
}

type Pattern struct {
	width, height  int
	startx, starty int
	depth          [][]int
	nextpixel      *pq.PriorityQueue
}

func NewPattern(width, height int) *Pattern {
	depth := make([][]int, height)

	for iy := 0; iy < height; iy++ {
		depth[iy] = make([]int, width)
		for ix := 0; ix < width; ix++ {
			depth[iy][ix] = -1
		}
	}

	startx, starty := 0, 0

	var nextpixel pq.PriorityQueue
	nextpixel.Push(pq.MakeObject(pixel{startx, starty, startx, starty}, rand.Float64()))

	return &Pattern{
		width:     width,
		height:    height,
		startx:    startx,
		starty:    starty,
		depth:     depth,
		nextpixel: &nextpixel,
	}
}

type Iteration struct {
	hasValue    bool
	x, y, depth int
}

func (pattern *Pattern) Iterate() Iteration {

	var p pixel
	for {
		if pattern.nextpixel.GetSize() == 0 {
			return Iteration{
				hasValue: false,
			}
		}
		p = pattern.nextpixel.GetFront().Value.(pixel)
		pattern.nextpixel.PopFront()

		if pattern.depth[p.y][p.x] < 0 {
			break
		}
	}

	for i := 0; i < len(direction); i++ {

		nx := p.x + direction[i][0]
		ny := p.y + direction[i][1]

		if nx < 0 {
			nx += pattern.width
		}
		if ny < 0 {
			ny += pattern.height
		}
		if nx >= pattern.width {
			nx -= pattern.width
		}
		if ny >= pattern.height {
			ny -= pattern.height
		}

		if pattern.depth[ny][nx] < 0 {
			pattern.nextpixel.Push(pq.MakeObject(pixel{p.x, p.y, nx, ny}, rand.Float64()))
		}
	}

	pattern.depth[p.y][p.x] = pattern.depth[p.parent_y][p.parent_x] + 1

	return Iteration{
		hasValue: true,
		x:        p.x,
		y:        p.y,
		depth:    pattern.depth[p.y][p.x],
	}
}
