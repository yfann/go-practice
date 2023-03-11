package leetcode

func LengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	lastOccurred := make(map[byte]int)
	start, maxlength := 0, 0
	for i := 0; i < n; i++ {
		if lastOccurred[s[i]] >= start {
			start = lastOccurred[s[i]] + 1
		}
		lastOccurred[s[i]] = i
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
	}
	return maxlength
}
