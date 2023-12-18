package mmphonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var myanmarNumbers = map[rune]int{
	'၀': 0,
	'၁': 1,
	'၂': 2,
	'၃': 3,
	'၄': 4,
	'၅': 5,
	'၆': 6,
	'၇': 7,
	'၈': 8,
	'၉': 9,
}

func SanitizeInput(phone string) (string, error) {
	var sanitized string
	if len(phone) == 0 {
		return "", errors.New("phone number should not be empty")
	}
	re := regexp.MustCompile(`[- ()]`)
	sanitized = re.ReplaceAllString(phone, "")
	phoneRe := regexp.MustCompile(`^\+?950?9\d+$`)
	startWithNine := phoneRe.MatchString(sanitized)

	if startWithNine {
		doubleCountryCodeRe := regexp.MustCompile(`/^\+?95950?9\d{7,9}$/`)
		if doubleCountryCodeRe.MatchString(sanitized) {
			sanitized = regexp.MustCompile(`9595`).ReplaceAllString(sanitized, "95")
		}
		zeroBeforeAreaCodeRe := regexp.MustCompile(`/^\+?9509\d{7,9}$/`)
		if zeroBeforeAreaCodeRe.MatchString(sanitized) {
			sanitized = regexp.MustCompile(`9509`).ReplaceAllString(sanitized, "959")
		}
	}
	return sanitized, nil
}

func NormalizeInput(phone string) (string, error) {
	var normalized string
	sanitizedNumber, err := SanitizeInput(phone)
	if err != nil {
		return "", err
	}
	normalized = sanitizedNumber
	possibleRe := regexp.MustCompile(`/^((09-)|(\+959)|(09\s)|(959)|(09\.))/`)
	if possibleRe.MatchString(sanitizedNumber) {
		normalized = possibleRe.ReplaceAllString(sanitizedNumber, "09")
	}

	myanmarNumPattern := regexp.MustCompile(`[၀-၉]`)
	if myanmarNumPattern.MatchString(normalized) {
		var mappedChars []string
		for _, char := range sanitizedNumber {
			val, exists := myanmarNumbers[char]
			if exists {
				mappedChars = append(mappedChars, fmt.Sprint(val))
			} else {
				mappedChars = append(mappedChars, string(char))
			}
		}
		mappedResult := strings.Join(mappedChars, "")
		normalized = possibleRe.ReplaceAllString(mappedResult, "09")
		return normalized, nil
	}
	return normalized, nil
}

func IsValidMMPhoneNumber(phone string) bool {
	normalized, _ := NormalizeInput(phone)
	myanmarPhoneRe := regexp.MustCompile(`^(09|\+?950?9|\+?95950?9)\d{7,9}$`)
	return myanmarPhoneRe.MatchString(normalized)
}

func getTelecomName() {

}

func getPhoneNetworkType() {

}
