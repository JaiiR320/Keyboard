package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/JaiiR320/keyboard"
)

func main() {
	exPath, err := os.Executable()
	if err != nil {
		panic(err)

	}
	exeDir := filepath.Dir(exPath)

	kb := keyboard.New(25)
	kb.NextCommandDelay = 800

	time.Sleep(2 * time.Second)
	path := exeDir

	// uses command prompt to create a file and open it with notepad
	kb.Hit(keyboard.VK_LWIN)
	kb.Write("cmd")
	kb.Hit(keyboard.VK_ENTER)
	kb.Write("cd " + path)
	kb.Hit(keyboard.VK_ENTER)
	kb.Write("echo.> test.txt")
	kb.Hit(keyboard.VK_ENTER)
	kb.Write("notepad test.txt")
	kb.Hit(keyboard.VK_ENTER)
	kb.Write("hello world")
	kb.Hold(keyboard.VK_LCONTROL)
	kb.Hit(keyboard.VK_S)
	kb.Lift(keyboard.VK_LCONTROL)
}
