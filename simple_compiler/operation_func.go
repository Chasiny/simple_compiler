package simple_compiler

import (
	"Compiler/gramma_analysis"
	"Compiler/lexical_analysis"
	"fmt"
	"strconv"
)

func (sc *SimpleCompiler) statementFunc(words []lexical_analysis.Word) error {
	if len(words) < 2 {
		return fmt.Errorf("statement func fail: words len <2")
	}

	if words[0].WordType != sc.lexicalAnalysisor.KeywordMap["int"] {
		return fmt.Errorf("statement func fail: statement type isn't int")
	}

	for i := 1; i < len(words); i++ {
		vInt, err := strconv.Atoi(words[i].Value)
		if err != nil {
			return err
		}
		if _, ok := sc.variableMap[vInt]; ok {
			return fmt.Errorf("statement func fail: %s exist", words[i].Word)
		}

		sc.variableMap[vInt] = 0
		sc.keyMap[vInt] = words[i].Word
	}

	return nil
}

func (sc *SimpleCompiler) operationFunc(words []lexical_analysis.Word, tree *gramma_analysis.AnalysisTreeNode) error {
	n, err := sc.setTreeValue(words, tree, 0)
	if err != nil {
		return err
	}

	sc.countTree(tree)

	fmt.Println(n)

	return nil
}

func (sc *SimpleCompiler) setTreeValue(words []lexical_analysis.Word, tree *gramma_analysis.AnalysisTreeNode, l int) (int, error) {

	if tree == nil {
		return l, nil
	}

	fmt.Println(tree.Value)

	if _, ok := sc.tagMap[tree.Value]; ok {

		fmt.Println("set", tree.Value, "to", words[l].Value, words[l].Word)
		tree.RealValue = words[l].Value
		l = l + 1
	}
	for i := range tree.Next {
		n, err := sc.setTreeValue(words, tree.Next[i], l)
		if err != nil {
			return l, err
		}
		l = n
	}

	return l, nil
}

func (sc *SimpleCompiler) countTree(tree *gramma_analysis.AnalysisTreeNode) (res []*gramma_analysis.AnalysisTreeNode, err error) {
	if tree == nil {
		return
	}
	if len(tree.Next) < 1 {
		if _, ok := sc.tagMap[tree.Value]; ok {
			return []*gramma_analysis.AnalysisTreeNode{tree}, nil
		}
		return
	}

	child := []*gramma_analysis.AnalysisTreeNode{}
	for i := range tree.Next {
		tres, err := sc.countTree(tree.Next[i])
		if err != nil {
			return nil, err
		}

		for t := range tres {
			if _, ok := sc.tagMap[tres[t].Value]; ok {
				child = append(child, tres[t])
			}
		}
	}

	if len(child) < 3 {
		return child, nil
	}

	fmt.Println(child)

	return sc.execOperation(child)
}

func (sc *SimpleCompiler) execOperation(in []*gramma_analysis.AnalysisTreeNode) (res []*gramma_analysis.AnalysisTreeNode, err error) {
	if len(in) < 3 {
		return in, nil
	}

	fmt.Println(len(in))
	for i := range in {
		fmt.Println(fmt.Sprintf("%+v", in[i]))
	}

	switch in[1].Value {
	case "=":
		value, err := sc.getValue(in[2])
		if err != nil {
			return nil, err
		}

		err = sc.setValue(in[0], value)
		if err != nil {
			return nil, err
		}

		fmt.Println("assign", in[0].ID, "to", value)
		res = append(res, in[2])

	case "-":
		value1, err := sc.getValue(in[2])
		if err != nil {
			return nil, err
		}
		value2, err := sc.getValue(in[0])
		if err != nil {
			return nil, err
		}
		res = append(res, &gramma_analysis.AnalysisTreeNode{Value: "i", RealValue: strconv.Itoa(value2 - value1)})

	case "+":
		value1, err := sc.getValue(in[2])
		if err != nil {
			return nil, err
		}
		value2, err := sc.getValue(in[0])
		if err != nil {
			return nil, err
		}
		res = append(res, &gramma_analysis.AnalysisTreeNode{Value: "i", RealValue: strconv.Itoa(value2 + value1)})

	}

	fmt.Println("return")
	for i := range res {
		fmt.Println(fmt.Sprintf("%+v", res[i]))
	}

	return res, nil
}

func (sc *SimpleCompiler) getValue(in *gramma_analysis.AnalysisTreeNode) (int, error) {
	if in.Value == "v" {
		id, err := strconv.Atoi(in.RealValue.(string))
		if err != nil {
			return 0, err
		}
		if v, ok := sc.variableMap[id]; ok {
			return v, nil
		}
		return 0, fmt.Errorf("can't find %s", in.Value)
	} else if in.Value == "i" {
		return strconv.Atoi(in.RealValue.(string))
	}

	return 0, fmt.Errorf("can't find %s", in.Value)
}

func (sc *SimpleCompiler) setValue(in *gramma_analysis.AnalysisTreeNode, v int) error {
	id, err := strconv.Atoi(in.RealValue.(string))
	if err != nil {
		return err
	}
	sc.variableMap[id] = v

	return nil
}
