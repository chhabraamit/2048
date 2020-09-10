package game

import (
	"fmt"
	"math/rand"
)

const rows = 4
const cols = 4

type Board interface {
	Display()
}

type board struct {
	matrix [][]int
}

func (b board) Display() {
	b.matrix = getRandom()
	printHorizontal()
	for i := 0; i < len(b.matrix); i++ {
		printVertical()
		for j := 0; j < len(b.matrix[0]); j++ {
			if b.matrix[i][j] == 0 {
				fmt.Printf("%6s", "")
			}else {
				fmt.Printf("%6d", b.matrix[i][j])
			}
		}
		fmt.Printf("%4s", "")
		printVertical()
		fmt.Println()
	}
	printHorizontal()
}

func printVertical() {
	fmt.Print("|")
}

func printHorizontal() {
	for i := 0; i < 30; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func getRandom() [][]int {
	arr := make([]int, 0)
	val := 2
	arr = append(arr, 0)
	arr = append(arr, val)
	for true {
		val *= 2
		arr = append(arr, val)
		if val == 2048 {
			break
		}
	}
	size := len(arr)
	board := make([][]int, 0)
	for i := 0; i < 4; i++ {
		board = append(board, make([]int, 4))
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			element := arr[rand.Int()%size]
			board[i][j] = element
		}
	}
	return board
}

func New() Board {
	return &board{}
}
