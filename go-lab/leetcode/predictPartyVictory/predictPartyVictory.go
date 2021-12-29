// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/dota2-senate/solution/dota2-can-yi-yuan-by-leetcode-solution-jb7l/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// 投票是一轮一轮来的, 被ban了之后就永久不能投了

package predictpartyvictory

func predictPartyVictory(senate string) string {
	var radiant, dire []int
	for i, s := range senate {
		if s == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}
	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}
	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
