package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		return
	}

	emos := map[string]string{
		"gb": "🇬🇧",
		"je": "🇯🇪",
		"de": "🇩🇪",
		"mx": "🇲🇽",
		"lv": "🇱🇻",
		"fr": "🇫🇷",
		"hu": "🇭🇺",
		"zw": "🇿🇼",

		"burger":    "🍔",
		"taco":      "🌮",
		"banana":    "🍌",
		"croissant": "🥐",
		"fries":     "🍟",

		"unicorn": "🦄",
		"monkey":  "🐒",
		"cat":     "😾",
		"goat":    "🐐",
	}

	emoji := emos[strings.ToLower(strings.TrimSpace(os.Args[1]))]
	message := fmt.Sprintf("%s %s %s", emoji, os.Args[2], emoji)

	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.Run()
}
