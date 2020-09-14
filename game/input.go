package game

import (
	"github.com/eiannone/keyboard"
	log "github.com/sirupsen/logrus"
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
	log.Debugf("the key is: %v \n", ans)
	if err != nil {
		return NO_DIR, err
	}
	switch ans {
	case 119, 65517:
		return UP, nil
	case 97, 65515:
		return LEFT, nil
	case 115, 65516:
		return DOWN, nil
	case 100, 65514:
		return RIGHT, nil
	case 3:
		return NO_DIR, errEndGame
	}
	return NO_DIR, nil
}
