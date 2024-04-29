package main

import (
	"fmt"
	"sort"
)

type Frequency struct {
	Char  rune
	Count int
}

type CharFrequency struct {
	frequencies map[rune]int
}

func NewCharFrequency(str string) CharFrequency {
	return CharFrequency{frequencies: countCharacters(str)}
}

func (freq CharFrequency) Values() []Frequency {
	var values []Frequency
	for k, v := range freq.frequencies {
		values = append(values, Frequency{k, v})
	}
	return values
}

func (freq CharFrequency) ValuesDESC() []Frequency {
	values := freq.Values()
	sort.Slice(values, func(i, j int) bool {
		return values[i].Count > values[j].Count
	})

	return values
}

func (freq CharFrequency) ToString() string {
	str := ""
	for _, value := range freq.ValuesDESC() {
		str += fmt.Sprintf("'%c' => %d\n", value.Char, value.Count)
	}

	return str
}

func countCharacters(input string) map[rune]int {
	charFrequency := make(map[rune]int)
	for _, char := range input {
		charFrequency[char]++
	}
	return charFrequency
}
