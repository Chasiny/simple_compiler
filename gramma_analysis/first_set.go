package gramma_analysis

import "fmt"

func (ll1 *LL1Analysis) creteFirstSet() error {
	if ll1.FirstSet == nil {
		ll1.FirstSet = make(map[string]map[string]struct{})
	}

	for k, _ := range ll1.terminalSymbol {
		m, ok := ll1.FirstSet[k]
		if !ok {
			ll1.FirstSet[k] = make(map[string]struct{})
			m = ll1.FirstSet[k]
		}
		m[k] = struct{}{}
	}

	l := 0
	r := 0

	isNext := true
	for {
		isNext = false
		for gk, gv := range ll1.Grammar {
			if ll1.FirstSet[gk] == nil {
				ll1.FirstSet[gk] = make(map[string]struct{})
			}

			for nk, _ := range gv {
				l = 0
				r = 1

				if len(nk) != 1 && nk[1:2] == `'` {
					r++
				}

				_, ok := ll1.noTerminalSymbol[nk[l:r]]
				if ok {
					count, err := ll1.mergeFirstSet(nk[l:r], gk)
					if err != nil {
						continue
					}
					if count != 0 {
						isNext = true
					}
				} else if _, exit := ll1.FirstSet[gk][nk[l:r]]; !exit {
					ll1.FirstSet[gk][nk[l:r]] = struct{}{}
					isNext = true
				}

				if _, ok := ll1.FirstSet[gk]["@"]; ll1.isAllEmpty(nk) && !ok {
					ll1.FirstSet[gk]["@"] = struct{}{}
					isNext = true
				}
			}
		}

		if !isNext {
			break
		}
	}

	return nil
}

func (ll1 *LL1Analysis) mergeFirstSet(src, dest string) (int, error) {
	srcSet, ok := ll1.FirstSet[src]
	if !ok {
		return 0, fmt.Errorf("%s first set not exist", src)
	}
	destSet, ok := ll1.FirstSet[dest]
	if !ok {
		return 0, fmt.Errorf("%s first set not exist", dest)
	}

	mergeCount := 0
	for k := range srcSet {
		if k == "@" {
			continue
		}

		if _, ok := destSet[k]; !ok {
			destSet[k] = struct{}{}
			mergeCount++
		}
	}

	return mergeCount, nil
}

func (ll1 *LL1Analysis) isContainEmpty(src string) bool {
	if src == "@" {
		return true
	}

	_, ok := ll1.FirstSet[src]["@"]
	return ok
}

func (ll1 *LL1Analysis) isAllEmpty(src string) bool {
	l, r := 0, 1
	for l < len(src) {
		r = l + 1
		if l != len(src)-1 && src[l+1:l+2] == "'" {
			r++
		}

		if !ll1.isContainEmpty(src[l:r]) {
			return false
		}
		l = r
	}

	return true
}

func (ll1 *LL1Analysis) GetFirstSet() (res []string) {
	for k, v := range ll1.FirstSet {
		t := "FIRST(" + k + ")" + "={"
		for s, _ := range v {
			t += s + " "
		}
		t += "}"

		res = append(res, t)
	}

	return
}
