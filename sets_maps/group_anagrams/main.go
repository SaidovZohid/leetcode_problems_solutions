package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)
	ns := [][]string{}
	for _, v := range strs {
		temp := []byte(v)
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
		if _, ok := mp[string(temp)]; ok {
			mp[string(temp)] = append(mp[string(temp)], v)
			continue
		}
		mp[string(temp)] = []string{v}
	}

	for _, v := range mp {
		ns = append(ns, v)
	}

	return ns
}

// * Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// * An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	strs = []string{""}
	fmt.Println(groupAnagrams(strs))
	strs = []string{"a"}
	fmt.Println(groupAnagrams(strs))
}
