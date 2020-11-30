// https://adventofcode.com/2019/day/3
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	X, Y int
}

type wirePoint struct {
	direction string
	length    int
}

type wire struct {
	Path []point
}

func newPoint(x int, y int) *point {
	return &point{X: x, Y: y}
}

func (p1 point) addPoint(p2 point) point {
	return *newPoint(p1.X+p2.X, p1.Y+p2.Y)
}
func (p1 point) equals(p2 point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func (p1 point) stringify() string {
	return strconv.Itoa(p1.X) + "-" + strconv.Itoa(p1.Y)
}

func (wp wirePoint) generateCoveredPoints(referencePoint point) []point {
	var result []point
	pointDelta := newPoint(0, 0)
	if wp.direction == "R" {
		pointDelta.X = 1

	} else if wp.direction == "U" {
		pointDelta.Y = -1

	} else if wp.direction == "D" {
		pointDelta.Y = 1

	} else {
		pointDelta.X = -1
	}

	prevPoint := referencePoint

	for i := 0; i < wp.length; i++ {
		result = append(result, prevPoint.addPoint(*pointDelta))
		prevPoint = result[len(result)-1]
	}

	return result

}

func (w *wire) addPointsToPath(points []point) {
	for _, p := range points {
		w.Path = append(w.Path, p)
	}
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func manhattanDistance(p1 point, p2 point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func newWire() *wire {
	return &wire{}
}

func newWirePoint(direction string, length int) *wirePoint {
	return &wirePoint{direction, length}
}

func constructWireWithPath(wirePath []wirePoint, referencePoint point) wire {
	wire := newWire()
	wire.addPointsToPath([]point{referencePoint})

	for _, currentWirePoint := range wirePath {
		coveredPoints := currentWirePoint.generateCoveredPoints(referencePoint)
		wire.addPointsToPath(coveredPoints)
		referencePoint = coveredPoints[len(coveredPoints)-1]
	}

	return *wire

}

func findStepsToPoint(points []point, targetPoint point) int {
	for i, point := range points {
		if point.equals(targetPoint) {
			return i
		}
	}
	return -1
}

func findMinimumSignalDelay(wire1 wire, wire2 wire, intersectionPoints []point) int {
	minimumSignalDelay := math.MaxInt16

	for _, point := range intersectionPoints {
		currentSignalDelay := findSignalDelay(wire1, wire2, point)
		if currentSignalDelay < minimumSignalDelay {
			minimumSignalDelay = currentSignalDelay
		}
	}
	return minimumSignalDelay

}

func findSignalDelay(wire1 wire, wire2 wire, intersectionPoint point) int {
	return findStepsToPoint(wire1.Path, intersectionPoint) + findStepsToPoint(wire2.Path, intersectionPoint)

}

func findNearestPointToReference(pointsToSearch []point, referencePoint point) (point, int) {
	nearestIntersectionDistance := math.MaxInt16
	var nearestIntersectionPoint point

	for _, point := range pointsToSearch {
		distance := manhattanDistance(point, centralPortPoint)
		if distance < nearestIntersectionDistance {
			nearestIntersectionDistance = distance
			nearestIntersectionPoint = point
		}
	}
	return nearestIntersectionPoint, nearestIntersectionDistance

}

func findIntersectingPoints(points1 []point, points2 []point) []point {
	var result []point
	ht := make(map[string]bool)

	for _, p1 := range points1 {
		key := p1.stringify()
		ht[key] = true
	}
	for _, p2 := range points2 {
		key := p2.stringify()
		_, ok := ht[key]

		if ok && !p2.equals(centralPortPoint) {
			result = append(result, p2)
		}
	}
	return result
}

var (
	centralPortPoint = *newPoint(0, 0)
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var wirePaths [][]wirePoint

	for scanner.Scan() {
		text := scanner.Text()
		wirePath := make([]wirePoint, 0)

		for _, pathPoint := range strings.Split(strings.TrimRight(text, "\n"), ",") {
			direction := string(pathPoint[0])
			distance, _ := strconv.Atoi(pathPoint[1:])
			wirePath = append(wirePath, *newWirePoint(direction, distance))
		}
		wirePaths = append(wirePaths, wirePath)
	}

	wire1 := constructWireWithPath(wirePaths[0], centralPortPoint)
	wire2 := constructWireWithPath(wirePaths[1], centralPortPoint)

	intersectingPoints := findIntersectingPoints(wire1.Path, wire2.Path)

	nearestIntersectionPoint, nearestIntersectionDistance := findNearestPointToReference(intersectingPoints, centralPortPoint)

	fmt.Println("Nearest intersection distance is ", nearestIntersectionDistance)
	fmt.Println("Nearest intersection point is ", nearestIntersectionPoint)
	fmt.Println("Minimum signal delay is ", findMinimumSignalDelay(wire1, wire2, intersectingPoints))

}
