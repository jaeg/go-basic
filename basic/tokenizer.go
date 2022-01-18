package basic

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type TokenType int

const (
	TokenTypeVariable TokenType = iota
	TokenTypeString
	TokenTypeNumber
	TokenTypeFunction
	TokenTypeControl
	TokenTypeOp
	TokenTypeEquals
	TokenTypeParens
	TokenTypeWord
	TokenTypeLabel
	TokenTypeInvalid
	TokenTypeNewLine
)

type Token struct {
	Value string
	Type  TokenType
}

func Tokenize(input []string) ([]Token, error) {
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
		//inputChunked[lineI] = append(inputChunked[lineI], "\n")
	}

	//Get rid of empty lines
	for lineI := 0; lineI < len(inputChunked); lineI++ {
		if len(inputChunked[lineI]) == 0 {
			inputChunked = append(inputChunked[:lineI], inputChunked[lineI+1:]...)
			lineI--
		}
	}

	state := "search" //  string, search, number
	invalid := false
	specialCharacters := "=+-*/<>()!"
	builtInfunctions := []string{"write"}
	controlFunctions := []string{"if", "endif", "goto", "else", "print"}

	tokenizedProgram := []Token{}
	tempToken := ""
	for lineI := 0; lineI < len(inputChunked); lineI++ {
		for chunkI := 0; chunkI < len(inputChunked[lineI]); chunkI++ {
			chunk := inputChunked[lineI][chunkI]
			switch state {
			case "search":
				tempToken = ""

				if strings.Contains(specialCharacters, chunk) {
					t := TokenTypeOp
					if chunk == "=" {
						t = TokenTypeEquals
					}
					if chunk == "(" || chunk == ")" {
						t = TokenTypeParens
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
				} else if _, err := strconv.Atoi(chunk); err == nil {
					state = "number"
					tempToken = chunk
					break
				} else if chunk == "\"" {
					state = "string"
					break
				} else if inArray(builtInfunctions, chunk) {
					tokenizedProgram = append(tokenizedProgram, Token{Value: chunk, Type: TokenTypeFunction})
					break
				} else if inArray(controlFunctions, chunk) {
					tokenizedProgram = append(tokenizedProgram, Token{Value: chunk, Type: TokenTypeControl})
					break
				} else if chunk == "\n" {
					tokenizedProgram = append(tokenizedProgram, Token{Value: chunk, Type: TokenTypeNewLine})
					break
				} else {
					state = "word"
					tempToken = chunk

					if chunkI == len(inputChunked[lineI])-1 && lineI == len(inputChunked)-1 {
						invalid = true
						tokenizedProgram = append(tokenizedProgram, Token{Value: chunk, Type: TokenTypeInvalid})
					}
					break
				}
			case "number":
				//Handling decimal numbers
				if chunk == "." {
					tempToken += "."
				} else if tempToken[len(tempToken)-1] == '.' {
					//Handle the numbers that come after the decimal while checking to make sure they are actual numbers.
					if _, err := strconv.Atoi(chunk); err != nil {
						tokenizedProgram = append(tokenizedProgram, Token{Value: tempToken, Type: TokenTypeInvalid})
						state = "search" // Return to searching for the next token type.
						invalid = true
					} else {
						tempToken += chunk
					}
				} else {
					//The number is complete add to the program
					tokenizedProgram = append(tokenizedProgram, Token{Value: tempToken, Type: TokenTypeNumber})
					state = "search"

					// Because this tokenization is happening to tokens in passing if we have a
					// completed number at chunk 0 it it means we had a new line.
					if chunkI == 0 {
						tokenizedProgram = append(tokenizedProgram, Token{Value: "", Type: TokenTypeNewLine})
					}

					chunkI-- //We take a step back so we make sure the scan search state can check it out
				}

				//Handle being the last chunk..
				if chunkI == len(inputChunked[lineI])-1 && lineI == len(inputChunked)-1 {
					tokenizedProgram = append(tokenizedProgram, Token{Value: tempToken, Type: TokenTypeNumber})
				}

			case "string":
				if chunk == "\"" {
					tokenizedProgram = append(tokenizedProgram, Token{Value: tempToken, Type: TokenTypeString})
					state = "search"
				} else {
					if tempToken != "" {
						tempToken += " "
					}
					tempToken += chunk
				}

			case "word":
				if chunk == "_" {
					tempToken += chunk
				} else if chunk == ":" { //Colons indicate labels.
					tokenizedProgram = append(tokenizedProgram, Token{Value: tempToken, Type: TokenTypeLabel})
					state = "search"
				} else {
					tokenizedProgram = append(tokenizedProgram, Token{Value: tempToken, Type: TokenTypeWord})
					state = "search"
					chunkI--
				}
			}
		}

		if state == "search" {
			tokenizedProgram = append(tokenizedProgram, Token{Value: "", Type: TokenTypeNewLine})
		}

	}

	//Do a quick check for invalid t
	var err error
	if invalid {
		err = errors.New("invalid syntax while tokenizing")
	}

	return tokenizedProgram, err
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
