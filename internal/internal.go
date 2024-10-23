package internal

import (
	"strings"
)

func Ascii(args ...string) (asciiArt, err, state string) {
	var (
		validUserInput string
		bannerMap      map[rune][]string
	)
	validUserInput, err = UserInputChecker(args[0])
	if err != "" {
		return "", err, "405"
	}
	bannerMap, err = MapBuilder(args[1])
	if err != "" {
		return "", err, "404"
	}
	asciiArt = BuildAsciiArt(strings.Split(validUserInput, "\\n"), bannerMap)
	return asciiArt, "", "200"
}
