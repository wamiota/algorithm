package recursion2

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	var res []string
	parenthesis(n, 1, 0, "(", res)
	return res
}

func parenthesis(max int, lnum int, rnum int, cur string, list []string) []string {
	if lnum < rnum || max < lnum || max < rnum {
		return list
	} else if lnum == max && rnum == max {
		fmt.Println(cur)
		list = append(list, cur)
		return list
	}
	list = parenthesis(max, lnum+1, rnum, cur+"(", list)
	list = parenthesis(max, lnum, rnum+1, cur+")", list)

	return list
}
