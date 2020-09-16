package game

import (
	"github.com/eiannone/keyboard"
	"log"
)

// GetCharKeystroke returns a key without pressing enter/return key
// supported keys are [W A S D] and Arrow keys
// below magical numbers are their key codes
func GetCharKeystroke() (Dir, error) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	char, key, err := keyboard.GetKey()
	ans := int(char)
	if ans == 0 {
		ans = int(key)
	}
	if DebugLogLevel {
		log.Printf("the key is: %v \n", ans)
	}
	if err != nil {
		return NO_DIR, err
	}
	switch ans {
	case 119, 65517, 107:
		return UP, nil
	case 97, 65515, 104:
		return LEFT, nil
	case 115, 65516, 106:
		return DOWN, nil
	case 100, 65514, 108:
		return RIGHT, nil
	case 3:
		return NO_DIR, errEndGame
	}
	return NO_DIR, nil
}
