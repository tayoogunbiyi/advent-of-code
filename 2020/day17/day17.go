// https://adventofcode.com/2020/day/17#part2
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Core ideas from - https://github.com/CodingNagger/advent-of-code-2020/blob/master/pkg/days/day17/computer.go

var (
	OCCUPIED uint8 = '#'
)

type HyperCube struct {
	X, Y, Z, W int
}

type HyperGrid struct {
	Cubes map[HyperCube]bool
}

type Cube struct {
	X, Y, Z int
}

type Grid struct {
	Cubes map[Cube]bool
}

func (c Cube) getNeighbours() []Cube {
	var res []Cube

	for _, v := range c.getNeighbourVectors() {
		res = append(res, Cube{c.X + v.X, c.Y + v.Y, c.Z + v.Z})
	}

	return res
}

func (c *Cube) getNeighbourVectors() []Cube {
	return []Cube{
		{X: -1, Y: -1, Z: -1}, {X: 0, Y: -1, Z: -1}, {X: 1, Y: -1, Z: -1},
		{X: -1, Y: 0, Z: -1}, {X: 0, Y: 0, Z: -1}, {X: 1, Y: 0, Z: -1},
		{X: -1, Y: 1, Z: -1}, {X: 0, Y: 1, Z: -1}, {X: 1, Y: 1, Z: -1},
		{X: -1, Y: -1, Z: 0}, {X: 0, Y: -1, Z: 0}, {X: 1, Y: -1, Z: 0},
		{X: -1, Y: 0, Z: 0}, {X: 1, Y: 0, Z: 0},
		{X: -1, Y: 1, Z: 0}, {X: 0, Y: 1, Z: 0}, {X: 1, Y: 1, Z: 0},
		{X: -1, Y: -1, Z: 1}, {X: 0, Y: -1, Z: 1}, {X: 1, Y: -1, Z: 1},
		{X: -1, Y: 0, Z: 1}, {X: 0, Y: 0, Z: 1}, {X: 1, Y: 0, Z: 1},
		{X: -1, Y: 1, Z: 1}, {X: 0, Y: 1, Z: 1}, {X: 1, Y: 1, Z: 1},
	}
}

func (g *Grid) countActiveNeighbours(c Cube) int {
	count := 0
	for _, neighbourCubes := range c.getNeighbours() {
		if g.Cubes[neighbourCubes] {
			count++
		}
	}
	return count
}

func (g *Grid) Cycle() {
	nextState := map[Cube]bool{}
	g.markInvisibleNeighboursAsUnoccupied()
	for c, isActive := range g.Cubes {
		activeNeighbourCount := g.countActiveNeighbours(c)
		nextStateIsActive := false

		if isActive {
			nextStateIsActive = activeNeighbourCount == 2 || activeNeighbourCount == 3
		} else {
			nextStateIsActive = activeNeighbourCount == 3
		}
		nextState[c] = nextStateIsActive
	}

	g.Cubes = nextState

}
func (g *Grid) markInvisibleNeighboursAsUnoccupied() {
	var invisibleNeighbours []Cube

	for cube := range g.Cubes {
		for _, neighbourCube := range cube.getNeighbours() {
			if !g.Cubes[neighbourCube] {
				invisibleNeighbours = append(invisibleNeighbours, neighbourCube)
			}
		}
	}

	for _, c := range invisibleNeighbours {
		g.Cubes[c] = false
	}
}

func parseInput(input string) *Grid {
	g := Grid{}
	g.Cubes = map[Cube]bool{}

	for row, line := range strings.Split(input, "\n") {
		if len(line) > 0 {
			for col := range line {
				Cube := Cube{
					X: col,
					Y: row,
					Z: 0,
				}
				g.Cubes[Cube] = line[col] == OCCUPIED
			}
		}
	}
	return &g
}

func (g *Grid) countActiveCubes() int {
	count := 0
	for _, isActive := range g.Cubes {
		if isActive {
			count++
		}
	}
	return count
}

func (hc *HyperCube) getNeighbours() []HyperCube {
	var res []HyperCube

	for _, v := range hc.getNeighbourVectors() {
		res = append(res, HyperCube{X: hc.X + v.X, Y: hc.Y + v.Y, Z: hc.Z + v.Z, W: hc.W + v.W})
	}
	return res
}

func (hc *HyperCube) getNeighbourVectors() []HyperCube {
	return []HyperCube{
		{X: -1, Y: -1, Z: -1, W: -1}, {X: 0, Y: -1, Z: -1, W: -1}, {X: 1, Y: -1, Z: -1, W: -1},
		{X: -1, Y: 0, Z: -1, W: -1}, {X: 0, Y: 0, Z: -1, W: -1}, {X: 1, Y: 0, Z: -1, W: -1},
		{X: -1, Y: 1, Z: -1, W: -1}, {X: 0, Y: 1, Z: -1, W: -1}, {X: 1, Y: 1, Z: -1, W: -1},
		{X: -1, Y: -1, Z: 0, W: -1}, {X: 0, Y: -1, Z: 0, W: -1}, {X: 1, Y: -1, Z: 0, W: -1},
		{X: -1, Y: 0, Z: 0, W: -1}, {X: 0, Y: 0, Z: 0, W: -1}, {X: 1, Y: 0, Z: 0, W: -1},
		{X: -1, Y: 1, Z: 0, W: -1}, {X: 0, Y: 1, Z: 0, W: -1}, {X: 1, Y: 1, Z: 0, W: -1},
		{X: -1, Y: -1, Z: 1, W: -1}, {X: 0, Y: -1, Z: 1, W: -1}, {X: 1, Y: -1, Z: 1, W: -1},
		{X: -1, Y: 0, Z: 1, W: -1}, {X: 0, Y: 0, Z: 1, W: -1}, {X: 1, Y: 0, Z: 1, W: -1},
		{X: -1, Y: 1, Z: 1, W: -1}, {X: 0, Y: 1, Z: 1, W: -1}, {X: 1, Y: 1, Z: 1, W: -1},

		{X: -1, Y: -1, Z: -1, W: 0}, {X: 0, Y: -1, Z: -1, W: 0}, {X: 1, Y: -1, Z: -1, W: 0},
		{X: -1, Y: 0, Z: -1, W: 0}, {X: 0, Y: 0, Z: -1, W: 0}, {X: 1, Y: 0, Z: -1, W: 0},
		{X: -1, Y: 1, Z: -1, W: 0}, {X: 0, Y: 1, Z: -1, W: 0}, {X: 1, Y: 1, Z: -1, W: 0},
		{X: -1, Y: -1, Z: 0, W: 0}, {X: 0, Y: -1, Z: 0, W: 0}, {X: 1, Y: -1, Z: 0, W: 0},
		{X: -1, Y: 0, Z: 0, W: 0} /*, {X: 0, Y: 0, Z: 0, W: 0} */, {X: 1, Y: 0, Z: 0, W: 0},
		{X: -1, Y: 1, Z: 0, W: 0}, {X: 0, Y: 1, Z: 0, W: 0}, {X: 1, Y: 1, Z: 0, W: 0},
		{X: -1, Y: -1, Z: 1, W: 0}, {X: 0, Y: -1, Z: 1, W: 0}, {X: 1, Y: -1, Z: 1, W: 0},
		{X: -1, Y: 0, Z: 1, W: 0}, {X: 0, Y: 0, Z: 1, W: 0}, {X: 1, Y: 0, Z: 1, W: 0},
		{X: -1, Y: 1, Z: 1, W: 0}, {X: 0, Y: 1, Z: 1, W: 0}, {X: 1, Y: 1, Z: 1, W: 0},

		{X: -1, Y: -1, Z: -1, W: 1}, {X: 0, Y: -1, Z: -1, W: 1}, {X: 1, Y: -1, Z: -1, W: 1},
		{X: -1, Y: 0, Z: -1, W: 1}, {X: 0, Y: 0, Z: -1, W: 1}, {X: 1, Y: 0, Z: -1, W: 1},
		{X: -1, Y: 1, Z: -1, W: 1}, {X: 0, Y: 1, Z: -1, W: 1}, {X: 1, Y: 1, Z: -1, W: 1},
		{X: -1, Y: -1, Z: 0, W: 1}, {X: 0, Y: -1, Z: 0, W: 1}, {X: 1, Y: -1, Z: 0, W: 1},
		{X: -1, Y: 0, Z: 0, W: 1}, {X: 0, Y: 0, Z: 0, W: 1}, {X: 1, Y: 0, Z: 0, W: 1},
		{X: -1, Y: 1, Z: 0, W: 1}, {X: 0, Y: 1, Z: 0, W: 1}, {X: 1, Y: 1, Z: 0, W: 1},
		{X: -1, Y: -1, Z: 1, W: 1}, {X: 0, Y: -1, Z: 1, W: 1}, {X: 1, Y: -1, Z: 1, W: 1},
		{X: -1, Y: 0, Z: 1, W: 1}, {X: 0, Y: 0, Z: 1, W: 1}, {X: 1, Y: 0, Z: 1, W: 1},
		{X: -1, Y: 1, Z: 1, W: 1}, {X: 0, Y: 1, Z: 1, W: 1}, {X: 1, Y: 1, Z: 1, W: 1},
	}

}

func (hg *HyperGrid) countActiveNeighbours(hc HyperCube) int {
	count := 0
	for _, neighbourCubes := range hc.getNeighbours() {
		if hg.Cubes[neighbourCubes] {
			count++
		}
	}
	return count
}

func (hg *HyperGrid) Cycle() {
	nextState := map[HyperCube]bool{}
	hg.markInvisibleNeighboursAsUnoccupied()
	for c, isActive := range hg.Cubes {
		activeNeighbourCount := hg.countActiveNeighbours(c)
		nextStateIsActive := false

		if isActive {
			nextStateIsActive = activeNeighbourCount == 2 || activeNeighbourCount == 3
		} else {
			nextStateIsActive = activeNeighbourCount == 3
		}
		nextState[c] = nextStateIsActive
	}

	hg.Cubes = nextState

}
func (hg *HyperGrid) markInvisibleNeighboursAsUnoccupied() {
	var invisibleNeighbours []HyperCube

	for cube := range hg.Cubes {
		for _, neighbourCube := range cube.getNeighbours() {
			if !hg.Cubes[neighbourCube] {
				invisibleNeighbours = append(invisibleNeighbours, neighbourCube)
			}
		}
	}

	for _, c := range invisibleNeighbours {
		hg.Cubes[c] = false
	}
}

func parseInputFor2(input string) *HyperGrid {
	hg := HyperGrid{}
	hg.Cubes = map[HyperCube]bool{}

	for row, line := range strings.Split(input, "\n") {
		if len(line) > 0 {
			for col := range line {
				hc := HyperCube{
					X: col,
					Y: row,
					Z: 0,
					W: 0,
				}
				hg.Cubes[hc] = line[col] == OCCUPIED
			}
		}
	}
	return &hg
}

func (hg *HyperGrid) countActiveHyperCubes() int {
	count := 0
	for _, isActive := range hg.Cubes {
		if isActive {
			count++
		}
	}
	return count
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	grid := parseInput(string(data))
	for i := 0; i < 6; i++ {
		grid.Cycle()
	}
	fmt.Println(grid.countActiveCubes())

	grid2 := parseInputFor2(string(data))

	for i := 0; i < 6; i++ {
		grid2.Cycle()
	}
	fmt.Println(grid2.countActiveHyperCubes())

}
