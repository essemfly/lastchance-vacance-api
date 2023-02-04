package crawler

import (
	"strconv"
	"strings"
	"unicode"
)

func ParsePriceString(letter string) int {
	ret := ""
	for _, val := range letter {
		if unicode.IsDigit(val) {
			ret += string(val)
		}
	}
	retInt, _ := strconv.Atoi(ret)
	return retInt
}

// "관심 8 ∙채팅 2∙조회 158"
func ParseViewCountsString(letter string) (int, int) {
	likeCount, viewCount := 0, 0
	tokens := strings.Split(letter, "")
	for _, token := range tokens {
		tokenKeyPair := strings.Split(token, " ")
		if len(tokenKeyPair) < 2 {
			continue
		}
		if tokenKeyPair[0] == "관심" {
			likeCount, _ = strconv.Atoi(tokenKeyPair[1])
		} else if tokenKeyPair[0] == "조회" {
			viewCount, _ = strconv.Atoi(tokenKeyPair[1])
		}
	}
	return likeCount, viewCount
}
