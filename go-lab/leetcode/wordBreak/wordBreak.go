// 140. 单词拆分 II
// https://leetcode-cn.com/problems/word-break-ii/

package wordbreak

import (
	"strings"
)

// 回溯 (超时)
// func wordBreak(s string, wordDict []string) []string {
// 	res := make([]string, 0)
// 	memo := make([]string, 0)
// 	dict := make(map[string]int)
// 	for k, v := range wordDict {
// 		dict[v] = k
// 	}
// 	backtrack(&res, &memo, &dict, s)
// 	return res
// }

// func backtrack(res, memo *[]string, dict *map[string]int, s string) {
// 	if s == "" {
// 		*res = append(*res, strings.Join(*memo, " "))
// 		return
// 	}

// 	for k := range s {
// 		if _, ok := (*dict)[s[:k+1]]; ok {
// 			*memo = append(*memo, s[:k+1])
// 			backtrack(res, memo, dict, s[k+1:])
// 			*memo = (*memo)[:len(*memo)-1]
// 		}
// 	}

// }

func wordBreak(s string, wordDict []string) (sentences []string) {
	wordSet := make(map[string]int)
	for k, w := range wordDict {
		wordSet[w] = k
	}

	n := len(s)
	dp := make([][][]string, n)
	var backtrack func(index int) [][]string
	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}
		wordsList := [][]string{}
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, has := wordSet[word]; has {
				for _, nextWords := range backtrack(i) {
					wordsList = append(wordsList, append([]string{word}, nextWords...))
				}
			}
		}
		word := s[index:]
		if _, has := wordSet[word]; has {
			wordsList = append(wordsList, []string{word})
		}
		dp[index] = wordsList
		return wordsList
	}
	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}
	return
}
