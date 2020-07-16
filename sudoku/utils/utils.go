package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

type Sudoku [9][9]int

// PrintSudoku : Prints sudoku in a fancy way
func PrintSudoku(s Sudoku) {
	fmt.Println("_________________________")

	for i, arr := range s {
		for index, value := range arr {
			if index == 0 {
				fmt.Print("|")
			}

			if value == 0 {
				fmt.Print(" .")
			} else {
				fmt.Printf(" %v", value)
			}

			if index%3 == 2 {
				fmt.Print(" |")
			}
		}

		fmt.Println()
		if i%3 == 2 && i != 8 {
			fmt.Println("|-------|-------|-------|")
		}
	}

	fmt.Println("\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e\u203e")
}

// SolveSudoku : Solves sudoku starting from the given position
func SolveSudoku(s Sudoku, r, c int) (Sudoku, error) {
	if r == -1 && c == -1 {
		return s, nil
	}

	if s[r][c] != 0 {
		nr, nc := getNextSquare(s, r, c)
		return SolveSudoku(s, nr, nc)
	}

	possibleValues := getPossibleValues(s, r, c)
	if len(possibleValues) == 0 {
		return s, errors.New("No possible solution.")
	}

	possibleValues = rearangeArray(possibleValues)

	for _, v := range possibleValues {
		s[r][c] = v
		nr, nc := getNextSquare(s, r, c)
		ns, err := SolveSudoku(s, nr, nc)
		if err == nil {
			return ns, nil
		}
	}

	return s, errors.New("Time to back-track.")
}

// getNextSquare : Returns the position of next empty square (row wise) or -1, -1 if all are filled
func getNextSquare(s Sudoku, r, c int) (int, int) {
	for i := r; i < 9; i++ {
		var start int
		if i == r {
			start = c
		} else {
			start = 0
		}

		for j := start; j < 9; j++ {
			if s[i][j] == 0 {
				return i, j
			}
		}
	}

	return -1, -1
}

// getPossibleValues : Returns an array of possible values at a position
func getPossibleValues(s Sudoku, r, c int) []int {
	if s[r][c] != 0 {
		arr := make([]int, 1)
		arr[0] = s[r][c]
		return arr
	}

	notPossibleValues := make([]int, 0)
	for i := 0; i < 9; i++ {
		if s[i][c] != 0 {
			notPossibleValues = append(notPossibleValues, s[i][c])
		}

		if s[r][i] != 0 {
			notPossibleValues = append(notPossibleValues, s[r][i])
		}
	}

	notPossibleValues = append(notPossibleValues, getSmallSquareElements(s, r, c)...)

	possibleValues := make([]int, 0)

	for i := 1; i <= 9; i++ {
		if !includes(notPossibleValues, i) {
			possibleValues = append(possibleValues, i)
		}
	}

	return possibleValues
}

// getSmallSquareElements : Returns the 3x3 square elements at a position
func getSmallSquareElements(s Sudoku, r, c int) []int {
	var rmin, rmax, cmin, cmax int

	if r >= 0 && r <= 2 {
		rmin = 0
		rmax = 2
	} else if r >= 3 && r <= 5 {
		rmin = 3
		rmax = 5
	} else {
		rmin = 6
		rmax = 8
	}

	if c >= 0 && c <= 2 {
		cmin = 0
		cmax = 2
	} else if c >= 3 && c <= 5 {
		cmin = 3
		cmax = 5
	} else {
		cmin = 6
		cmax = 8
	}

	elements := make([]int, 0)

	for i := rmin; i <= rmax; i++ {
		for j := cmin; j <= cmax; j++ {
			if s[i][j] != 0 {
				elements = append(elements, s[i][j])
			}
		}
	}

	return elements
}

// rearangeArray : Randomly re-aranges the array elements
func rearangeArray(arr []int) []int {
	nArr := make([]int, len(arr))
	index := 0

	for i := len(arr) - 1; i >= 0; i-- {
		rando := getRandomInt(0, i)
		nArr[index] = arr[rando]
		index++
		arr = splice(arr, rando, 1)
	}

	return nArr
}

// getRandomInt : Gives random integer between min and max (both inclusive)
func getRandomInt(min, max int) int {
	max = max + 1
	m := max - min
	mb := big.NewInt(int64(m))

	rb, _ := rand.Int(rand.Reader, mb)
	r := int(rb.Int64())

	return r + min
}

/*------------------------------------------------------------------------------
	 ARRAY UTILITY FUNCTIONS
------------------------------------------------------------------------------*/

// includes : Checks if an element is present in an array
func includes(arr []int, n int) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}

	return false
}

// splice : Deletes n elements from a given position
func splice(arr []int, pos, n int) []int {
	return append(arr[:pos], arr[pos+n:]...)
}
