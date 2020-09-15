package game

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

const (
	_rows = 4
	_cols = 4

	// this is the sequence which is used to clear the screen :magic
	// _clearScreenSequence = "\033[H\033[2J" // this works in mac. Might need other string for other OS
	_clearScreenSequence =	"\033c" // this clears the screen and scroll buffer

	probabilitySpace = 100
	probabilityOfTwo = 80 // probabilityOfTwo times 2 will come as new element out of  probabilitySpace1
)

type Board interface {
	Display()
	AddElement()
	TakeInput()
	IsOver() bool
	CountScore() (int, int)
}

type board struct {
	matrix [][]int
	over   bool
	newRow int
	newCol int
}

func (b *board) CountScore() (int, int) {
	total := 0
	maximum := 0
	matrix := b.matrix
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			total += matrix[i][j]
			maximum = max(maximum, matrix[i][j])
		}
	}
	return maximum, total
}

func max(one int, two int) int {
	if one > two {
		return one
	}
	return two
}

func (b *board) IsOver() bool {
	empty := 0
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			if b.matrix[i][j] == 0 {
				empty++
			}
		}
	}
	return empty == 0 || b.over
}

func (b *board) TakeInput() {
	var dir Dir
	dir, err := GetCharKeystroke()
	if err != nil {
		if errors.Is(err, errEndGame) {
			b.over = true
			return
		} else {
			log.Fatal("error while taking input for game: %v", err)
			return
		}
	}
	log.Debugf("the dir is: %v \n", dir)
	if dir == NO_DIR {
		// this makes pressing any keys other than move-set doesn't make any change in the game
		b.TakeInput() // retry to get a valid direction
	}
	b.move(dir)
}

// AddElement : it first finds the empty slots in the board. They are the one with 0 value
// The it places a new cell randomly in one of those empty places
// The new value to put is also calculated randomly
func (b *board) AddElement() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	val := r1.Int() % probabilitySpace
	if val <= probabilityOfTwo {
		val = 2
	} else {
		val = 4
	}

	empty := 0
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			if b.matrix[i][j] == 0 {
				empty++
			}
		}
	}
	elementCount := r1.Int()%empty + 1
	index := 0

	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			if b.matrix[i][j] == 0 {
				index++
				if index == elementCount {
					b.newRow = i
					b.newCol = j
					b.matrix[i][j] = val
					return
				}
			}
		}
	}
	return
}

// Display this is the method which draws the board
// board contains a matrix which has cells. Each cell is a number.
// A Cell with 0 is considered empty
// to display number pretty, we make use of left and right padding
// Grid is formed using Ascii characters and some amount of test-&-see
func (b *board) Display() {
	d := color.New(color.FgBlue, color.Bold)
	//b.matrix = getRandom()
	fmt.Println(_clearScreenSequence)
	for i := 0; i < len(b.matrix); i++ {
		printHorizontal()
		for j := 0; j < len(b.matrix[0]); j++ {
			fmt.Printf("%3s", "")
			if b.matrix[i][j] == 0 {
				fmt.Printf("%-6s|", "")
			} else {
				if i == b.newRow && j == b.newCol {
					d.Printf("%-6d|", b.matrix[i][j])
				} else {
					fmt.Printf("%-6d|", b.matrix[i][j])
				}
			}
		}
		fmt.Printf("%4s", "")
		fmt.Println()
	}
	printHorizontal()
}

// printVertical for printing a vertical grid element
func printVertical() {
	log.Debug("|")
}

// printHorizontal prints a grid row
func printHorizontal() {
	for i := 0; i < 40; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func New() Board {
	matrix := make([][]int, 0)
	for i := 0; i < _rows; i++ {
		matrix = append(matrix, make([]int, _cols))
	}
	return &board{
		matrix: matrix,
	}
}

var errEndGame = errors.New("GameOverError")
