package basic

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Token struct {
	Value string
	Type  string
}

func Tokenize(input []string) []Token {
	//Clean up input and chunk it apart.
	inputChunked := make([][]string, 1)
	for lineI := 0; lineI < len(input); lineI++ {
		tokenizedLine := Split(regexp.MustCompile("([^a-zA-Z0-9])"), input[lineI], -1)
		inputChunked = append(inputChunked, []string{})
		for i := 0; i < len(tokenizedLine); i++ {
			if strings.TrimSpace(tokenizedLine[i]) != "" {
				inputChunked[lineI] = append(inputChunked[lineI], tokenizedLine[i])
			}
		}
	}

	//Get rid of empty lines
	for lineI := 0; lineI < len(inputChunked); lineI++ {
		if len(inputChunked[lineI]) == 0 {
			inputChunked = append(inputChunked[:lineI], inputChunked[lineI+1:]...)
			lineI--
		}
	}

	state := "search" //  string, search, number
	specialCharacters := "=+-*/<>()!"
	builtInfunctions := []string{"write"}
	controlFunctions := []string{"if", "endif", "goto", "print", "move", "else"}

	tokenizedProgram := []Token{}
	tempToken := ""
	for lineI := 0; lineI < len(inputChunked); lineI++ {
		for chunkI := 0; chunkI < len(inputChunked[lineI]); chunkI++ {
			chunk := inputChunked[lineI][chunkI]
			switch state {
			case "search":
				tempToken = ""
				if strings.Contains(specialCharacters, chunk) {
					t := "op"
					if chunk == "=" {
						t = "equals"
					}
					if chunk == "(" || chunk == ")" {
						t = "paren"
					}
					if chunk == "!" {
						if chunkI < len(inputChunked[lineI])-1 {
							if strings.Contains(specialCharacters, inputChunked[lineI][chunkI+1]) {
								chunk += inputChunked[lineI][chunkI+1]
								chunkI++
							}
						}
					}
					tokenizedProgram = append(tokenizedProgram, Token{Value: chunk, Type: t})
					break
				} else if _, err := strconv.Atoi(chunk); err == nil {
					state = "number"
					tempToken = chunk
				} else if chunk == "\"" {
					state = "string"
				} else if inArray(builtInfunctions, chunk) {
					tokenizedProgram = append(tokenizedProgram, Token{Value: chunk, Type: "function"})
				} else if inArray(controlFunctions, chunk) {
					tokenizedProgram = append(tokenizedProgram, Token{Value: chunk, Type: "control"})
				} else {
					state = "word"
					tempToken = chunk
				}

				if chunkI == len(inputChunked[lineI])-1 && lineI == len(inputChunked)-1 {
					tokenizedProgram = append(tokenizedProgram, Token{Value: chunk, Type: state})
				}
				break
			case "number":
				if chunk == "." {
					tempToken += "."
				} else if len(tempToken) > 1 {
					if tempToken[len(tempToken)-1] == '.' {
						if _, err := strconv.Atoi(chunk); err == nil {
							fmt.Println("Failed to tokenize this line: ", lineI)
							return nil
						} else {
							tempToken += chunk
						}
					}
				} else {
					tokenizedProgram = append(tokenizedProgram, Token{Value: tempToken, Type: "number"})
					state = "search"
					chunkI-- //We take a step back so we make sure the scan search state can check it out
				}
				break

			case "string":
				if chunk == "\"" {
					tokenizedProgram = append(tokenizedProgram, Token{Value: tempToken, Type: "string"})
					state = "search"
				} else {
					if tempToken != "" {
						tempToken += " "
					}
					tempToken += chunk
				}
				break
			case "word":
				if chunk == "_" {
					tempToken += chunk
				} else if chunk == ":" { //Colons indicate labels.
					tokenizedProgram = append(tokenizedProgram, Token{Value: tempToken, Type: "label"})
					state = "search"
				} else {
					tokenizedProgram = append(tokenizedProgram, Token{Value: tempToken, Type: "word"})
					state = "search"
					chunkI--
				}
			}
		}
	}
	return tokenizedProgram
}

func Split(re *regexp.Regexp, s string, n int) []string {
	if n == 0 {
		return nil
	}

	matches := re.FindAllStringIndex(s, n)
	strings := make([]string, 0, len(matches))

	beg := 0
	end := 0
	for _, match := range matches {
		if n > 0 && len(strings) >= n-1 {
			break
		}

		end = match[0]
		if match[1] != 0 {
			strings = append(strings, s[beg:end])
		}
		beg = match[1]
		// This also appends the current match
		strings = append(strings, s[match[0]:match[1]])
	}

	if end != len(s) {
		strings = append(strings, s[beg:])
	}

	return strings
}

func inArray(a []string, s string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == s {
			return true
		}
	}
	return false
}
