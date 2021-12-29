package wordbreak

import (
	"log"
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		res      []string
	}{
		{
			"catsanddog",
			[]string{"cat", "cats", "and", "sand", "dog"},
			[]string{
				"cats and dog",
				"cat sand dog",
			},
		},

		{
			"pineapplepenapple",
			[]string{"apple", "pen", "applepen", "pine", "pineapple"},
			[]string{
				"pine apple pen apple",
				"pineapple pen apple",
				"pine applepen apple",
			},
		},

		{
			"catsandog",
			[]string{"cats", "dog", "sand", "and", "cat"},
			[]string{},
		},
		// {
		// 	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		// 	[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
		// 	[]string{},
		// },
	}

	for _, tt := range tests {
		actual := wordBreak(tt.s, tt.wordDict)
		log.Printf("wordBreak(%v %v) = %v, expcted: %v", tt.s, tt.wordDict, actual, tt.res)
	}
}
