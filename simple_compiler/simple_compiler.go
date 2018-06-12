package simple_compiler

import (
	"Compiler/gramma_analysis"
	"Compiler/lexical_analysis"
	"fmt"
	"strconv"
)

type SimpleCompiler struct {
	lexicalAnalysisor   *lexical_analysis.LexicalAnalysis
	operationAnalysisor *gramma_analysis.LL1Analysis
	statementAnalysisor *gramma_analysis.LL1Analysis
	spitSymbol          string
	variableMap         map[int]int

	tagMap map[string]struct{}
	keyMap map[int]string
}

func NewSimpleCompiler() *SimpleCompiler {
	operationGrammar := `E->v=A
A->TE'
E'->+TE'|@
T->FT'
T'->*FT'|@
F->(E)|v|i`
	assignGrammar := `E->TF
T->k
F->v`

	operationll1, err := gramma_analysis.NewLL1Analysis(operationGrammar)
	if err != nil {
		panic(err)
		return nil
	}
	assignll1, err := gramma_analysis.NewLL1Analysis(assignGrammar)
	if err != nil {
		panic(err)
		return nil
	}

	res := &SimpleCompiler{
		spitSymbol:          ";",
		operationAnalysisor: operationll1,
		statementAnalysisor: assignll1,
		lexicalAnalysisor:   lexical_analysis.NewLexicalAnalysisDemo(),
		variableMap:         make(map[int]int),
		tagMap:              make(map[string]struct{}),
		keyMap:              make(map[int]string),
	}
	res.tagMap["+"] = struct{}{}
	res.tagMap["-"] = struct{}{}
	res.tagMap["*"] = struct{}{}
	res.tagMap["/"] = struct{}{}
	res.tagMap["i"] = struct{}{}
	res.tagMap["v"] = struct{}{}
	res.tagMap["k"] = struct{}{}
	res.tagMap["="] = struct{}{}

	return res
}

func (sc *SimpleCompiler) Run(p string) (res [][]string, err error) {
	words := sc.lexicalAnalysisor.Explain(p)

	curPos := 0
	for i := 0; i < len(words); i++ {
		if words[i].Word == sc.spitSymbol {
			err := sc.Exec(words[curPos:i])
			if err != nil {
				return res, err
			}
			curPos = i + 1
		}
	}

	for k, v := range sc.variableMap {
		res = append(res, []string{sc.keyMap[k], strconv.Itoa(v)})
	}

	return res, nil
}

func (sc *SimpleCompiler) Exec(words []lexical_analysis.Word) error {
	fmt.Println(sc.wordsToString(words))
	if _, _, ok := sc.statementAnalysisor.Analysis(sc.wordsToString(words)); ok {
		if err := sc.statementFunc(words); err != nil {
			return err
		}
	} else if _, tree, ok := sc.operationAnalysisor.Analysis(sc.wordsToString(words)); ok {
		if err := sc.operationFunc(words, tree); err != nil {
			return err
		}
	}

	return nil
}

func (sc *SimpleCompiler) wordsToString(words []lexical_analysis.Word) (res string) {
	for i := range words {
		if words[i].WordType == sc.lexicalAnalysisor.KeywordMap["int"] {
			res += "k"
		} else if words[i].WordType == sc.lexicalAnalysisor.KeywordMap["="] ||
			words[i].WordType == sc.lexicalAnalysisor.KeywordMap["+"] ||
			words[i].WordType == sc.lexicalAnalysisor.KeywordMap["-"] ||
			words[i].WordType == sc.lexicalAnalysisor.KeywordMap["*"] ||
			words[i].WordType == sc.lexicalAnalysisor.KeywordMap["/"] {
			res += words[i].Word
		} else if words[i].WordType == sc.lexicalAnalysisor.IntType {
			res += "i"
		} else if words[i].WordType == sc.lexicalAnalysisor.VariableType {
			res += "v"
		}
	}

	return res + "#"
}
