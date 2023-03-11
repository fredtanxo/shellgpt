package main

import "fmt"

// ShowBanner prints the banner of ShellGPT every time the program starts
func ShowBanner() {
	fmt.Print(`
  ____  _          _ _  ____ ____ _____ 
 / ___|| |__   ___| | |/ ___|  _ \_   _|
 \___ \| '_ \ / _ \ | | |  _| |_) || |
  ___) | | | |  __/ | | |_| |  __/ | |
 |____/|_| |_|\___|_|_|\____|_|    |_|

Type anything to start a conversation.

`)
}
