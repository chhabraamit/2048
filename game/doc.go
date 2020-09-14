package game

/**
This is an implementation of popular 2048 game.

The game has a board which is made of 2D array. Each array represents a row.
An empty cell is represented with a 0 value.

Refreshing of screen:
* The screen refresh illusion is coming from sending the clear screen command and printing the new board.
* This is screen-clear string could be machine dependent and might not work on non-mac machines. It may require a
  different value.
* You can see all the printed boards by scrolling back in terminal.

Accepting input without pressing return/enter key:
* Another point of interest is accepting key pressing without return keys.
* This is done using "github.com/eiannone/keyboard" library

Movement of the board:
* At fundamental level: only moving the board to left is implemented.
* Move right is implemented by first reversing the list, moving it left and then reversing it back.

* Move down is little tricky. For this the matrix is rotated clockwise first, so _cols become _rows and then it's moved
  left. Then, to bring it back, it's rotated three times more clockwise. It could have been just rotated anti-clock
  wise but I wanted to use the existing method.

* Move right is achieved by rearranging the board in reverse order-- bottom row becomes the top one and so on. Then
  move down method is applied and the _rows are reversed back again.

*/
