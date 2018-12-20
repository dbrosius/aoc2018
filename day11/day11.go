package main

import (
	"fmt"
)

func main() {
	serial := 5034
	grid := [300][300]int{}
	fillWithPower(&grid, serial)
	x, y, maxPower := findBest(&grid, 3)
	fmt.Printf("Size = 3, X: %d, Y: %d\n", x+1, y+1)
	maxSz, maxX, maxY := 0, 0, 0
	for sz := 1; sz <= 300; sz++ {
		xTest, yTest, pwrTest := findBest(&grid, sz)
		if pwrTest > maxPower {
			maxSz = sz
			maxX = xTest
			maxY = yTest
		}
	}
	fmt.Printf("Size = %d, X: %d, Y: %d, Max Power: %d\n", maxSz, maxX, maxY, maxPower)
}

func fillWithPower(grid *[300][300]int, serial int) {
	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			grid[x][y] = cellPower(x+1, y+1, serial)
		}
	}
}

func cellPower(x, y, serial int) int {
	rackID := x + 10
	power := rackID * y
	power = power + serial
	power = power * rackID
	power = (power / 100) % 10
	return power - 5
}

func findBest(grid *[300][300]int, sz int) (int, int, int) {
	maxPower := -46
	maxX := -1
	maxY := -1
	for x := 0; x < 300-sz; x++ {
		for y := 0; y < 300-sz; y++ {
			power := sumSubGrid(grid, x, y, sz)
			//fmt.Printf("X: %d, Y:%d, POWER: %d\n", x+1, y+1, power)
			if power > maxPower {
				maxPower = power
				maxX = x
				maxY = y
			}
		}
	}
	return maxX, maxY, maxPower
}

func sumSubGrid(grid *[300][300]int, x, y, sz int) int {
	sum := 0
	for xi := x; xi < x+sz; xi++ {
		for yi := y; yi < y+sz; yi++ {
			sum += grid[xi][yi]
		}
	}
	return sum
}
