package asciigenerator

import (
	"os"
	"path/filepath"
	"strings"
)

// ----------- PARSER ----------------
func ParseBanner(bannerType string) (map[rune][]string, error) {

	charMap := make(map[rune][]string, 95)

	path := filepath.Join("banners", bannerType+".txt")
	bannerFile, err := os.ReadFile(path)
	if err != nil {
		return charMap, err
	}
	banner := strings.ReplaceAll(string(bannerFile), "\r\n", "\n")
	lines := strings.Split(banner, "\n")

	var block []string
	code := 32 //ASCII space

	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		if line == "" {
			if len(block) > 0 {
				charMap[rune(code)] = block
				block = []string{}
				code++
			}
		} else {
			block = append(block, line)
		}
	}
	if len(block) > 0 {
		charMap[rune(code)] = block
	}
	return charMap, nil
}

func PrintBannertoArt(text string, charMap map[rune][]string) string {

	text = strings.ReplaceAll(text, "\\n", "\n")
	lines := strings.Split(text, "\n")

	var result []string

	for _, line := range lines {
		if line == "" {
			result = append(result, "")
			continue
		}

		for row := 0; row < 8; row++ {

			var builder strings.Builder

			for _, char := range line {
				block, ok := charMap[char]

				if ok {
					if row < len(block) {
						builder.WriteString(block[row])
					}
				} else {
					continue
				}
			}
			result = append(result, strings.TrimRight(builder.String(), " "))
		}
	}

	return strings.Join(result, "\n")
}
