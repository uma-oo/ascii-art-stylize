package internal

import (
	"os"
	"regexp"
	"strings"
)

// Check if the user input if valid
func UserInputChecker(input string) (string, string) {
	if input == "" {
		return "", "You need to provide a text"
	}

	for _, i := range input {
		if i < 32 || i > 126 {
			return "", "Found a character outside the range of printable ascii characters"
		}
	}
	return input, ""
}

// printable ASCII characters and the values are the corresponding ASCII art strings.
func BuildAsciiArt(input []string, asciiMap map[rune][]string) string {
	result := ""
	for _, i := range input {
		if i == "" {
			result += "\n"
			continue
		}
		for j := 0; j < 8; j++ {
			for _, k := range i {
				result += asciiMap[k][j]
			}
			result += "\n"
		}
	}
	return result
}

// Check if banner exists and format its data (\r\n to \n)
func BannerDataChecker(banner string) ([]string, string) {
	bannerContent, err := os.ReadFile("Banners/" + banner + ".txt")
	if err != nil {
		return []string{}, "Banner not Found!!"
	}
	bannerContent = []byte(regexp.MustCompile(`\r\n`).ReplaceAllString(string(bannerContent), "\n"))
	return strings.Split(string(bannerContent[1:len(bannerContent)-1]), "\n\n"), ""
}

// MapBuilder takes a slice of strings and returns a map where the keys are
func MapBuilder(banner string) (map[rune][]string, string) {
	data, err := BannerDataChecker(banner)
	if err != "" {
		return map[rune][]string{}, err
	}
	mapHolder := map[rune][]string{}
	indexCounter := 0
	for i := ' '; i <= '~'; i++ {
		mapHolder[i] = strings.Split(data[indexCounter], "\n")
		indexCounter++
	}
	return mapHolder, ""
}
