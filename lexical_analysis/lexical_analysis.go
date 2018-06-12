package lexical_analysis

import (
	"strconv"
	"strings"
)

func NewLexicalAnalysisDemo() *LexicalAnalysis {
	lexical := &LexicalAnalysis{}

	lowChar := "a|b|c|d|e|f|g|h|i|j|k|l|m|n|o|p|q|r|s|t|u|v|w|x|y|z"
	upChar := "A|B|C|D|E|F|G|H|I|J|K|L|M|N|O|P|Q|R|S|T|U|V|W|X|Y|Z"
	digitalChar := "0|1|2|3|4|5|6|7|8|9"
	charSet := lowChar + "|" + upChar

	variableReg := "(_|" + charSet + ")(_|" + charSet + "|" + digitalChar + ")*"
	intReg := "(" + digitalChar + ")*"
	floatReg := "(" + digitalChar + ")*.(" + digitalChar + ")*"

	lexical.variableDFA = NewDFA(NewNFA(variableReg))
	lexical.intDFA = NewDFA(NewNFA(intReg))
	lexical.floatDFA = NewDFA(NewNFA(floatReg))

	lexical.KeywordMap = make(map[string]int)
	lexical.wordMap = make(map[string]*Word)

	lexical.KeywordMap["="] = 101
	lexical.KeywordMap["+"] = 102
	lexical.KeywordMap["-"] = 103
	lexical.KeywordMap["*"] = 104
	lexical.KeywordMap["/"] = 105
	lexical.KeywordMap["<"] = 106
	lexical.KeywordMap[">"] = 107
	lexical.KeywordMap["<="] = 108
	lexical.KeywordMap[">="] = 109
	lexical.KeywordMap["=="] = 110
	lexical.KeywordMap["!="] = 111
	lexical.KeywordMap[";"] = 112
	lexical.KeywordMap[":"] = 113
	lexical.KeywordMap[","] = 114
	lexical.KeywordMap["{"] = 115
	lexical.KeywordMap["}"] = 116
	lexical.KeywordMap["["] = 117
	lexical.KeywordMap["]"] = 118
	lexical.KeywordMap["("] = 119
	lexical.KeywordMap[")"] = 120

	lexical.KeywordMap["main"] = 201
	lexical.KeywordMap["if"] = 202
	lexical.KeywordMap["else"] = 203
	lexical.KeywordMap["return"] = 204
	lexical.KeywordMap["void"] = 205
	lexical.KeywordMap["while"] = 206
	lexical.KeywordMap["int"] = 207
	lexical.KeywordMap["float"] = 208

	lexical.VariableType = 1
	lexical.IntType = 2
	lexical.FloatType = 3

	return lexical
}

type LexicalAnalysis struct {
	variableDFA *DFA
	intDFA      *DFA
	floatDFA    *DFA

	KeywordMap map[string]int

	VariableType int
	IntType      int
	FloatType    int

	variableID int

	wordMap map[string]*Word
}

func (l *LexicalAnalysis) Explain(input string) (res []Word) {

	input = strings.Replace(input, "\r\n", "\n", -1)

	preWord := Word{
		WordType: -1,
	}
	prePos := 0
	line := 0
	index := 0
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '\n':
			line++
			index = 0
		case ' ':
			if preWord.WordType != -1 {
				res = append(res, preWord)
				preWord.WordType = -1
			}
			prePos = i + 1
			index++
		default:
			newWord := l.findWord(input[prePos : i+1])
			if newWord.WordType != -1 {
				preWord = *newWord
				preWord.Line = line
				preWord.Index = index
			} else {
				prePos = i
				if preWord.WordType != -1 {
					res = append(res, preWord)
					preWord.WordType = -1
					i--
					index--
				}
			}
			index++
		}
	}

	if preWord.WordType != -1 {
		res = append(res, preWord)
	}

	return
}

func (l *LexicalAnalysis) findWord(s string) (res *Word) {
	if s == "" {
		return &Word{
			Word:     s,
			WordType: -1,
			Value:    s,
		}
	}

	if _, ok := l.wordMap[s]; ok {
		return l.wordMap[s]

	} else if _, ok := l.KeywordMap[s]; ok {
		res = &Word{
			Word:     s,
			WordType: l.KeywordMap[s],
			Value:    "",
		}

	} else if _, ok := l.KeywordMap[s]; ok {
		res = &Word{
			Word:     s,
			WordType: l.KeywordMap[s],
			Value:    "",
		}

	} else if l.variableDFA.Accept(s) {
		l.variableID++
		res = &Word{
			Word:     s,
			WordType: l.VariableType,
			Value:    strconv.Itoa(l.variableID),
		}

	} else if l.intDFA.Accept(s) {
		res = &Word{
			Word:     s,
			WordType: l.IntType,
			Value:    s,
		}

	} else if l.floatDFA.Accept(s) {
		res = &Word{
			Word:     s,
			WordType: l.FloatType,
			Value:    s,
		}

	} else {
		res = &Word{
			Word:     s,
			WordType: -1,
			Value:    s,
		}
	}

	l.wordMap[s] = res

	return
}

func (l *LexicalAnalysis) WordsToString(words []Word) (res [][]string) {

	res = append(res, []string{"单词", "类型", "值", "行号", "列号"})

	for i := range words {
		res = append(res, []string{words[i].Word, strconv.Itoa(words[i].WordType),
			words[i].Value, strconv.Itoa(words[i].Line), strconv.Itoa(words[i].Index)})
	}
	return
}
