package yaml

import (
	"fmt"
	"strconv"
	"strings"
)

// Token types.
const (
	TokenKey   = 1
	TokenIndex = 2
)

// Token presents the token by its type and values.
type Token struct {
	Type  int
	Key   string
	Index int
}

// split will take a dot separated line of YAML keys to be converted
// to a slice of tokens.
//
// i.e.: 'foo.bar[2].baz' translates to {'foo', 'bar', '[2]', 'baz'}
//
// TODO: Write a reliable parser that does not abuse string splits.
func split(key string) ([]Token, error) {
	tokens := make([]Token, 0)

	for _, str := range strings.Split(key, ".") {
		str = strings.TrimSpace(str)

		// Determine if array.
		if strings.Contains(str, "[") && strings.Contains(str, "]") {
			list := strings.Replace(str, "[", " ", -1)
			list = strings.Replace(list, "]", "", -1)
			array := strings.Split(list, " ")

			// Expected to have the array name and index separated as two strings.
			if len(array) != 2 {
				return tokens, fmt.Errorf("could not determine index from array '%s'", array[0])
			}

			// Test if index is an actual number
			index, err := strconv.Atoi(array[1])

			if err != nil {
				return tokens, fmt.Errorf("index '%s' is not a number", array[1])
			}

			tokenKey := Token{
				Key:  array[0],
				Type: TokenKey,
			}
			TokenIndex := Token{
				Index: index,
				Type:  TokenIndex,
			}
			tokens = append(tokens, tokenKey)
			tokens = append(tokens, TokenIndex)
			continue
		}

		token := Token{
			Key:  str,
			Type: TokenKey,
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}
