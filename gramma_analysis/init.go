package gramma_analysis

import (
	"fmt"
	"strings"
)

func (ll1 *LL1Analysis) initTerminalSymbol() error {
	if ll1.terminalSymbol == nil {
		ll1.terminalSymbol = make(map[string]struct{})
	}

	l, r := 0, 0
	for _, m := range ll1.Grammar {
		for k, _ := range m {
			for i := 0; i < len(k); i++ {
				if k[i:i+1] == `'` {
					continue
				}

				l = i
				r = i + 1
				if i != len(k)-1 && k[i+1:i+2] == `'` {
					r++
				}
				if _, ok := ll1.noTerminalSymbol[k[l:r]]; !ok {
					ll1.terminalSymbol[k[l:r]] = struct{}{}
				}
			}
		}
	}

	ll1.terminalSymbol["#"] = struct{}{}

	return nil
}

func (ll1 *LL1Analysis) initNoTerminalSymbol() error {
	if ll1.noTerminalSymbol == nil {
		ll1.noTerminalSymbol = make(map[string]struct{})
	}

	for k, _ := range ll1.Grammar {
		ll1.noTerminalSymbol[k] = struct{}{}
	}

	return nil
}

func NewLL1Analysis(grammar string) (ll1 *LL1Analysis, err error) {
	ll1 = &LL1Analysis{}
	grammar = strings.Replace(grammar, " ", "", -1)
	grammar = strings.Replace(grammar, "	", "", -1)
	grammars := strings.Split(grammar, "\n")

	if ll1.Grammar == nil {
		ll1.Grammar = make(map[string](map[string]struct{}))
	}

	ll1.start = strings.SplitN(grammars[0], "->", 2)[0]

	for i := 0; i < len(grammars); i++ {
		left := strings.SplitN(grammars[i], "->", 2)
		if len(left) < 2 {
			return nil, fmt.Errorf("in grammar fail: spit -> fail")
		}

		tGrammer := strings.Split(left[1], "|")
		if len(tGrammer) < 1 {
			return nil, fmt.Errorf("in grammar fail: spit | fail")
		}

		_, ok := ll1.Grammar[left[0]]
		if !ok {
			ll1.Grammar[left[0]] = make(map[string]struct{})
		}

		for i := 0; i < len(tGrammer); i++ {
			ll1.Grammar[left[0]][tGrammer[i]] = struct{}{}
		}

	}

	ll1.initNoTerminalSymbol()
	ll1.initTerminalSymbol()
	ll1.creteFirstSet()
	ll1.creteFollowSet()
	ll1.createAnalysisTable()

	return ll1, nil
}

func (ll1 *LL1Analysis) Show() {
	fmt.Println(fmt.Sprintf("%+v", ll1))
}
