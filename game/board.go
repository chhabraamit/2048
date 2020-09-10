package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const rows = 4
const cols = 4

type Board interface {
	Display()
	AddElement()
	TakeInput()
	IsOver() bool
}

type board struct {
	matrix [][]int
}

func (b *board) IsOver() bool {
	empty := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if b.matrix[i][j] == 0 {
				empty++
			}
		}
	}
	return empty == 0
}

func (b *board) TakeInput() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	switch string([]byte(input)[0]) {
	case "a":
		b.move(LEFT)
	case "d":
		b.move(RIGHT)
	case "w":
		b.move(UP)
	case "s":
		b.move(DOWN)
	}
	fmt.Printf("Input Char Is : %v\n", string([]byte(input)[0]))
}

type Dir int

const (
	UP Dir = iota
	DOWN
	LEFT
	RIGHT
)

func (b *board) move(dir Dir) {
	switch dir {
	case LEFT:
		b.moveLeft()
	case RIGHT:
		b.moveRight()
	case DOWN:
		b.moveDown()
	case UP:
		b.moveUp()
	}

}

func (b *board) moveDown() {
	b.transpose()
	b.moveLeft()
	b.transpose()
	b.transpose()
	b.transpose()
}

func (b *board) moveRight() {
	b.reverse()
	b.moveLeft()
	b.reverse()
}

func (b *board) moveLeft() {
	for i := 0; i < rows; i++ {
		old := b.matrix[i]
		b.matrix[i] = movedRow(old)
	}
}

func movedRow(elems []int) []int {
	nonEmpty := make([]int, 0)
	for i := 0; i < cols; i++ {
		if elems[i] != 0 {
			nonEmpty = append(nonEmpty, elems[i])
		}
	}
	remaining := cols - len(nonEmpty)
	for i := 0; i < remaining; i++ {
		nonEmpty = append(nonEmpty, 0)
	}
	return mergeElements(nonEmpty)
}

func (b *board) AddElement() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	val := r1.Int() % 100
	if val <= 69 {
		val = 2
	} else {
		val = 4
	}

	empty := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if b.matrix[i][j] == 0 {
				empty++
			}
		}
	}
	elementCount := r1.Int()%empty + 1
	index := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if b.matrix[i][j] == 0 {
				index++
				if index == elementCount {
					b.matrix[i][j] = val
					return
				}
			}
		}
	}
	return
}

const _clearScreenSequence = "\033[H\033[2J"

func (b *board) Display() {
	//b.matrix = getRandom()
	//fmt.Println(_clearScreenSequence)
	printHorizontal()
	for i := 0; i < len(b.matrix); i++ {
		printVertical()
		for j := 0; j < len(b.matrix[0]); j++ {
			if b.matrix[i][j] == 0 {
				fmt.Printf("%6s", "")
			} else {
				fmt.Printf("%6d", b.matrix[i][j])
			}
		}
		fmt.Printf("%4s", "")
		printVertical()
		fmt.Println()
	}
	printHorizontal()
}

func (b *board) reverse() {
	for i := 0; i < rows; i++ {
		b.matrix[i] = reversed(b.matrix[i])
	}
}

func (b *board) transpose() {
	ans := make([][]int, 0)
	for i := 0; i < rows; i++ {
		ans = append(ans, make([]int, cols))
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ans[i][j] = b.matrix[cols-j-1][i]
		}
	}
	b.matrix = ans
}

func (b *board) moveUp() {
	b.reverseRows()
	b.moveDown()
	b.reverseRows()
}

func (b *board) reverseRows() {
	ans := make([][]int, 0)
	for i := 0; i < rows; i++ {
		ans = append(ans, make([]int, cols))
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ans[rows-i-1][j] = b.matrix[i][j]
		}
	}
	b.matrix = ans
}

func reversed(arr []int) []int {
	ans := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		ans = append(ans, arr[i])
	}
	return ans
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
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		matrix = append(matrix, make([]int, cols))
	}
	return &board{
		matrix: matrix,
	}
}

func mergeElements(arr []int) []int {
	newArr := make([]int, len(arr))
	newArr[0] = arr[0]
	index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == newArr[index] {
			newArr[index] += arr[i]
		} else {
			index++
			newArr[index] = arr[i]
		}
	}
	return newArr
}
