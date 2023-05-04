package main

import "strings"

func main() {
	s := "A man, a plan, a canal: Panama"
	s1 := "race a car"

	println(isPalindrome(s))
	println(isPalindrome(s1))
}

// 125. 验证回文串
// 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。
// 则可以认为该短语是一个 回文串 。
// 字母和数字都属于字母数字字符。
// 给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
// 输入: s = "A man, a plan, a canal: Panama"
// 输出：true
// 解释："amanaplanacanalpanama" 是回文串。
// 输入：s = "race a car"
// 输出：false
// 解释："raceacar" 不是回文串。

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	left, right := 0, len(s) - 1
    for left < right {
        for left < right && !isalnum(s[left]) {
            left++
        }
        for left < right && !isalnum(s[right]) {
            right--
        }
        if left < right {
            if s[left] != s[right] {
                return false
            }
            left++
            right--
        }
    }
    return true
}
//首先先将字符串中的全部内容转换为小写形式
//随后使用正则表达式来判断是否为字母或者是数字
//将过滤后的结果拿出来
//最后使用两个指针，一个在前一个在后来滑动判断是否为回文串

func isalnum(ch byte) bool {
    return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
