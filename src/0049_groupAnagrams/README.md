# [字母异位词分组](https://leetcode-cn.com/problems/group-anagrams/)

### 题目

给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

```
输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
```

说明：  
* 所有输入均为小写字母。
* 不考虑答案输出的顺序。

### 解法

排序哈希存储或者唯一标示哈希存储，时空都不理想。

```
func groupAnagramsBySort(strs []string) [][]string {
	m := make(map[string][]string,len(strs))
	for _,str := range strs {
		sortStr := []byte(str)
		sort.Slice(sortStr, func(i, j int) bool {
			return sortStr[i] < sortStr[j]
		})
		if _, ok := m[string(sortStr)]; ok {
			m[string(sortStr)] = append(m[string(sortStr)], str)
		} else {
			m[string(sortStr)] = []string{str}
		}
	}
	res := make([][]string,0,len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		sign := make([]int, 26)
		for _, char := range str {
			sign[char-'a']++
		}
		signStr := make([]byte,0,52)
		for i := 0; i < 26; i++ {
			signStr = append(signStr, '#')
			signStr = append(signStr, byte(sign[i] + '0' - 0))
		}
		fmt.Println(string(signStr))
		if _, ok := m[string(signStr)]; ok {
			m[string(signStr)] = append(m[string(signStr)], str)
		} else {
			m[string(signStr)] = []string{str}
		}
	}
	fmt.Println(m)
	res := make([][]string,0,len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
```
