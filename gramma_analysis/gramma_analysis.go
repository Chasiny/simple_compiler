package gramma_analysis

type LL1Analysis struct {
	start string `json:"start"`

	Grammar       map[string]map[string]struct{} `json:"grammar"`
	AnalysisTable map[string]map[string]string   `json:"analysis_table"`

	FirstSet  map[string]map[string]struct{} `json:"first_set"`
	FollowSet map[string]map[string]struct{} `json:"follow_set"`

	terminalSymbol   map[string]struct{}
	noTerminalSymbol map[string]struct{}
}
