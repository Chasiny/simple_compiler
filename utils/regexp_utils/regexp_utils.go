package regexp_utils

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func EraseBrackets(s string) (res string) {
	left := 0

	i := 0
	for ; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			break
		}
	}

	min := left
	right := 0
	for ; i < len(s); i++ {
		if s[i] == '(' {
			left -= right
			right = 0
			min = Min(min, left)
			left++
		} else if s[i] == ')' {
			right++
		} else {
			left -= right
			right = 0
			min = Min(min, left)
		}
	}

	return s[min : len(s)-min]
}

func SpitAnd(s string) (res []string) {
	left := 0
	bracket := 0

	for i := range s {
		switch s[i] {

		case '(':
			if left != i && bracket == 0 {
				res = append(res, s[left:i])
				left = i
			}
			bracket++

		case ')':
			bracket--
			if left != i && bracket == 0 && !(i != len(s)-1 && s[i+1] == '*') {
				res = append(res, s[left:i+1])
				left = i + 1
			}

		case '*':
			if bracket == 0 {
				res = append(res, s[left:i+1])
				left = i + 1
			}

		default:
			if left != i && bracket == 0 {
				res = append(res, s[left:i])
				left = i
			}
		}

	}

	if left != len(s) {
		res = append(res, s[left:])
	}

	for i := range res {
		res[i] = EraseBrackets(res[i])
	}

	return
}

func SpitOr(s string) (res []string) {
	left := 0
	bracket := 0

	for i := range s {
		switch s[i] {

		case '(':
			bracket++

		case ')':
			bracket--

		case '|':
			if bracket == 0 {
				res = append(res, s[left:i])
				left = i + 1
			}

		}

	}

	if left != len(s) {
		res = append(res, s[left:])
	}

	for i := range res {
		res[i] = EraseBrackets(res[i])
	}

	return
}

func SpitClosure(s string) (res string, closure bool) {
	s = EraseBrackets(s)

	if s[len(s)-1] != '*' || len(s) < 2 {
		return s, false
	}

	if len(s) == 2 {
		return s[0:1], true
	}

	bracket := 0
	for i := 0; i < len(s)-2; i++ {
		switch s[i] {

		case '(':
			bracket++

		case ')':
			bracket--
			if bracket == 0 {
				return s, false
			}

		}
	}

	if bracket != 1 {
		return s, false
	}

	return EraseBrackets(s[1 : len(s)-2]), true
}
