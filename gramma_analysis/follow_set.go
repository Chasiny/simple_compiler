package gramma_analysis

import "fmt"

func (ll1 *LL1Analysis) creteFollowSet() {
	if ll1.FollowSet == nil {
		ll1.FollowSet = make(map[string]map[string]struct{})
	}

	if _, ok := ll1.FollowSet[ll1.start]; !ok {
		ll1.FollowSet[ll1.start] = make(map[string]struct{})
	}
	ll1.FollowSet[ll1.start]["#"] = struct{}{}

	isNext := true
	for {
		isNext = false

		l1, r1 := 0, 0
		l2, r2 := 0, 0

		for gk, gv := range ll1.Grammar {
			for nk, _ := range gv {
				l1 = 0
				r1 = 1

				if l1 != len(nk)-1 && nk[l1+1:l1+2] == "'" {
					r1++
				}
				if r1 >= len(nk) {
					continue
				}

				if n, _ := ll1.addFollowByParent(gk, nk, l1, r1); n > 0 {
					isNext = true
				}

				for {
					l2 = r1
					r2 = l2 + 1
					if l2 < len(nk)-1 && nk[l2+1:l2+2] == "'" {
						r2++
					}

					if r2 > len(nk) {
						break
					}

					if n, _ := ll1.addFollowByParent(gk, nk, l2, r2); n > 0 {
						isNext = true
					}

					if n, _ := ll1.mergeFirstToFollow(nk[l2:r2], nk[l1:r1]); n > 0 {
						isNext = true
					}

					l1 = l2
					r1 = r2

				}

			}
		}

		if !isNext {
			break
		}
	}
}

func (ll1 *LL1Analysis) addFollowByParent(parent, child string, l, r int) (int, error) {
	if r == len(child) || (ll1.isAllEmpty(child[r:]) && r < len(child)) {
		return ll1.mergeFollowToFollow(parent, child[l:r])
	}

	return 0, nil
}

func (ll1 *LL1Analysis) mergeFirstToFollow(src, dest string) (int, error) {
	if _, ok := ll1.noTerminalSymbol[dest]; !ok {
		return 0, fmt.Errorf("dest is not in no terminal symbol")
	}

	srcSet, ok := ll1.FirstSet[src]
	if !ok {
		return 0, fmt.Errorf("%s first set not exist", src)
	}
	destSet, ok := ll1.FollowSet[dest]
	if !ok {
		ll1.FollowSet[dest] = make(map[string]struct{})
		destSet = ll1.FollowSet[dest]
	}

	mergeCount := 0
	for k, _ := range srcSet {
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

func (ll1 *LL1Analysis) mergeFollowToFollow(src, dest string) (int, error) {
	_, oksrc := ll1.noTerminalSymbol[src]
	_, okdest := ll1.noTerminalSymbol[dest]
	if !oksrc || !okdest {
		return 0, fmt.Errorf("src or dest is not in no terminal symbol")
	}

	srcSet, ok := ll1.FollowSet[src]
	if !ok {
		ll1.FollowSet[src] = make(map[string]struct{})
		srcSet = ll1.FollowSet[src]
	}
	destSet, ok := ll1.FollowSet[dest]
	if !ok {
		ll1.FollowSet[dest] = make(map[string]struct{})
		destSet = ll1.FollowSet[dest]
	}

	mergeCount := 0
	for k, _ := range srcSet {
		if _, ok := destSet[k]; !ok {
			destSet[k] = struct{}{}
			mergeCount++
		}
	}

	return mergeCount, nil
}

func (ll1 *LL1Analysis) GetFollowSet() (res []string) {
	for k, v := range ll1.FollowSet {
		t := "FOLLOW(" + k + ")" + "={"
		for s, _ := range v {
			t += s + " "
		}
		t += "}"

		res = append(res, t)
	}

	return
}
