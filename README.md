# Keyboard - a simple keyboard simulator for Go

This program allows users to simulate keypresses on a keyboard. This will currently only work for Windows, as it depends on the windows virtual key codes. For more information check out the Microsoft page for keyboard input: [keybd_event function](https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-keybd_event), [virtual key codes](https://learn.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes)

## Installation

`go get github.com/JaiiR320/keyboard`

## Features

Simluate keystrokes using simple syntax. 

 - "Press" down any key on a standard keyboard
 - "Release" a key
 - "Hit" a key (perform a keystroke)
 - Simulate typing a character found on a keyboard (limited to qwerty and basic symbols most keyboards are capable of typing)

## Example

```Go
package main

func main(){
    // create a new keyboard
    kb := keyboard.New(25) // add the inter-key delay time

    // you can also set the inter key delay like this
    kb.InterKeyDelay = 25

    // the between-command delay
    // works for "Hit" and "Write"
    kb.NextCommandDelay = 1000

    // simulates a "tap" of the tab button
    kb.Hit(keyboard.VK_TAB)
    // pauses for 1000ms based on NextCommandDelay
    // simulates typing a specific character
    kb.Simulate('@')

    // simulates writing a string
    kb.Write("hello world")
}
```

## Future

Linux support coming soon
Open to suggestions, additions, tweaks, features, etc.

## Acknowledgements

This code was adapted from [micmonay/keybd_event](https://github.com/micmonay/keybd_event). Credit given for writing the key up and down functions, as well as the idea for adding 0xFFF to virtual keys.
