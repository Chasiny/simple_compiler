package gramma_analysis

func (ll1 *LL1Analysis) createAnalysisTable() {
	ll1.AnalysisTable = make(map[string]map[string]string)

	for gk := range ll1.noTerminalSymbol {
		_, ok := ll1.AnalysisTable[gk]
		if !ok {
			ll1.AnalysisTable[gk] = make(map[string]string)
		}

		for nk := range ll1.terminalSymbol {
			if nk == "@" {
				continue
			}

			if _, ok := ll1.FirstSet[nk]; ok {
				ll1.addTerminalAnalysisTable(gk, nk)
			}
		}
	}

	for gk := range ll1.noTerminalSymbol {
		if _, ok := ll1.FirstSet[gk]["@"]; ok {
			for f := range ll1.FollowSet[gk] {
				if _, ok := ll1.terminalSymbol[f]; ok {
					ll1.AnalysisTable[gk][f] = "@"
				}
			}
		}
	}

}

func (ll1 *LL1Analysis) addTerminalAnalysisTable(noTerminalSymbol, terminalSymbol string) {

	for k := range ll1.Grammar[noTerminalSymbol] {
		l, r := 0, 1
		if len(k) > 1 && k[1:2] == `'` {
			r++
		}

		if _, ok := ll1.FirstSet[k[l:r]][terminalSymbol]; ok {
			ll1.AnalysisTable[noTerminalSymbol][terminalSymbol] = k
		}
	}
}

func (ll1 *LL1Analysis) GetAnalysisTable() (res [][]string) {
	var head []string
	head = append(head, "")
	for k := range ll1.terminalSymbol {
		if k == "@" {
			continue
		}
		head = append(head, k)
	}

	res = append(res, head)
	for gk, _ := range ll1.noTerminalSymbol {
		var addRow []string
		addRow = append(addRow, gk)
		for i := range head {
			if head[i] == "" {
				continue
			}

			if _, ok := ll1.AnalysisTable[gk][head[i]]; !ok {
				addRow = append(addRow, "")
			} else {
				addRow = append(addRow, ll1.AnalysisTable[gk][head[i]])
			}
		}

		res = append(res, addRow)
	}

	return
}
