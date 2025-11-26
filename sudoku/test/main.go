package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/VishalSInha23/sudoku-solver/sudoku/utils"
)

type sudokuInput [9]string

func main() {
	exSInp := sudokuInput{
		"026000000",
		"000600003",
		"074080000",
		"000003002",
		"080040010",
		"600500000",
		"000010780",
		"500009000",
		"000000040",
	}

	fmt.Println("Example sudoku :")

	exS := convertSudokuInput(exSInp)
	utils.PrintSudoku(exS)

	fmt.Println("\nEnter the unsolved sudoku in the following form :")
	printInputSudoku(exSInp)

	fmt.Println("")

	reader := bufio.NewReader(os.Stdin)

	var sInp sudokuInput
	fmt.Println("Enter the unsolved sudoku (enter s to solve example sudoku) :")

	for i := 0; i < 9; i++ {
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error in reading input!!! Error : %v\n", err)
		}

		if str == "s\n" {
			sInp = exSInp
			break
		}

		sInp[i] = strings.TrimSuffix(str, "\n")
	}

	fmt.Println("\nInput sudoku :")
	// printInputSudoku(sInp)
	s := convertSudokuInput(sInp)
	utils.PrintSudoku(s)

	start := time.Now()

	res, err := utils.SolveSudoku(s, 0, 0)
	if err != nil {
		fmt.Printf("\nError in solving sudoku : %v\n", err)
	}

	end := time.Now()

	fmt.Println("\n Output sudoku :")
	utils.PrintSudoku(res)

	timeTakenNano := end.UnixNano() - start.UnixNano()
	timeTakenSecs := end.Unix() - start.Unix()

	fmt.Printf("\nTime taken to solve the sudoku : %v nano-seconds or %v seconds.\n", timeTakenNano, timeTakenSecs)
}

// convertSudokuInput : Converts input sudoku to sudoku array
func convertSudokuInput(sInp sudokuInput) utils.Sudoku {
	var s utils.Sudoku

	for i, val := range sInp {
		str := fillZeroesAtEnd(val)
		for j := 0; j < 9; j++ {
			var err error
			s[i][j], err = strconv.Atoi(string(str[j]))

			if err != nil {
				log.Fatalf("Error in converting input!!! Error : %v\n", err)
			}
		}
	}

	return s
}

// fillZeroesAtEnd : Fills zeroes at the end of the string so total len of the string becomes 9
func fillZeroesAtEnd(str string) string {
	zeroes := 9 - len(str)

	for i := 0; i < zeroes; i++ {
		str = str + "0"
	}

	return str
}

// printInputSudoku : Prints input sudoku in the input form
func printInputSudoku(s sudokuInput) {
	for _, val := range s {
		str := fillZeroesAtEnd(val)
		fmt.Println(str)
	}
}
