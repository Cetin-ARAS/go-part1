package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortByA(words []string) []string {
	// Kelimeleri sıralamadan önce bir özel tür tanımlayarak sıralama kriterlerini belirleyebiliriz.
	type wordWithCriteria struct {
		word        string
		numberOfAs int
		length      int
	}

	var wordsWithCriteria []wordWithCriteria

	// Her kelimenin içindeki 'a' harfi sayısını ve kelimenin uzunluğunu hesapla
	for _, word := range words {
		numberOfAs := strings.Count(word, "a")
		length := len(word)
		wordsWithCriteria = append(wordsWithCriteria, wordWithCriteria{word, numberOfAs, length})
	}

	// Sıralama kriterlerine göre sırala
	sort.Slice(wordsWithCriteria, func(i, j int) bool {
		// 'a' harfi sayısına göre azalan sıralama
		if wordsWithCriteria[i].numberOfAs > wordsWithCriteria[j].numberOfAs {
			return true
		} else if wordsWithCriteria[i].numberOfAs < wordsWithCriteria[j].numberOfAs {
			return false
		}

		// Eğer 'a' harfi sayıları eşitse uzunluğa göre artan sıralama
		return wordsWithCriteria[i].length < wordsWithCriteria[j].length
	})

	// Sıralanmış kelimeleri bir slice'e yerleştirerek döndür
	var result []string
	for _, w := range wordsWithCriteria {
		result = append(result, w.word)
	}

	return result
}

func main() {
	input := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	output := sortByA(input)
	fmt.Println(output)
}