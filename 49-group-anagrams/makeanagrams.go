/*
49. Group Anagrams

Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]


Constraints:

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.
*/

package groupanagrams

func groupAnagrams(strs []string) [][]string {
	result := make(map[[26]int][]string)
	for _, v := range strs {
		var count [26]int
		for _, n := range v {
			count[n-'a']++
		}
		result[count] = append(result[count], v)

	}
	resultMap := make([][]string, len(result))

	idx := 0
	for _, x := range result {
		resultMap[idx] = x
		idx++
	}
	return resultMap

}
