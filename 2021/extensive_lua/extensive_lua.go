package main

import "fmt"

var (
	cols = []int{11, 9, 13, 0, 11, 9, 13, 0, 6, 15, 6}
	rows = []int{54, 16, 34, 20, 54}
)

func main() {
	rowSet := makeSet(rows)
	colSet := makeSet(cols)

	// get row permutations that match given rows totals
	rowPerms := map[int][]int{}
	// rowInts := map[int]bool{}
	for rowInt := 0; rowInt < 1<<11; rowInt++ {
		row := makeRow(rowInt)
		var total int
		for i, b := range row {
			if b {
				total += i + 1
			}
		}
		if rowSet[total] {
			rowPerms[total] = append(rowPerms[total], rowInt)
			// rowInts[total] = true
		}
	}
	// get col permutations that match given cols totals
	colPerms := map[int][]int{}
	// colInts := map[int]bool{}
	for colInt := 0; colInt < 1<<11; colInt++ {
		col := makeCol(colInt)
		var total int
		for i, b := range col {
			if b {
				total += i + 1
			}
		}
		if colSet[total] {
			colPerms[total] = append(colPerms[total], colInt)
			// colInts[total] = true
		}
	}

	fmt.Println("rows:", rowPerms)
	fmt.Println("cols:", rowPerms)
	arr := make([][]bool, 5)
	for i := 0; i < 5; i++ {
		arr[i] = make([]bool, 11)
	}
	count := 0
	for row0 := range rowPerms[rows[0]] {
		for i, b := range makeRow(row0) {
			arr[0][i] = b
		}
		for row1 := range rowPerms[rows[1]] {
			for i, b := range makeRow(row1) {
				arr[1][i] = b
			}
			for row2 := range rowPerms[rows[2]] {
				for i, b := range makeRow(row2) {
					arr[2][i] = b
				}
				for row3 := range rowPerms[rows[3]] {
					for i, b := range makeRow(row3) {
						arr[3][i] = b
					}
				row4loop:
					for row4 := range rowPerms[rows[4]] {
						for i, b := range makeRow(row4) {
							arr[4][i] = b
						}
						count++
						for col := 0; col < 11; col++ {
							total := 0
							for r := 0; r < 5; r++ {
								if arr[r][col] {
									total += r + 1
								}
							}
							if total != cols[col] {
								break row4loop
							}
						}
						fmt.Println(arr)
					}
				}
			}
		}
	}
	fmt.Println("tested", count, "iterations")
}

func makeSet(arr []int) map[int]bool {
	r := map[int]bool{}
	for _, v := range arr {
		r[v] = true
	}
	return r
}

func makeCol(colInt int) []bool {
	var rc []bool
	for i := 0; i < 5; i++ {
		rc = append(rc, colInt&1 == 1)
		colInt = colInt >> 1
	}
	return rc
}

func makeRow(rowInt int) []bool {
	var rc []bool
	for i := 0; i < 11; i++ {
		rc = append(rc, rowInt&1 == 1)
		rowInt = rowInt >> 1
	}
	return rc
}
