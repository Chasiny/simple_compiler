package server

import (
	"Compiler/gramma_analysis"
	"Compiler/utils/data_structure/queue"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func submitGrammar(writer http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		log.Println(err)
		writer.Write([]byte(err.Error()))
		return
	}

	grammar := request.FormValue("grammar")
	inputdata := request.FormValue("inputdata")

	ll1, err := gramma_analysis.NewLL1Analysis(grammar)
	if err != nil {
		log.Println(err)
		writer.Write([]byte(err.Error()))
		return
	}

	table, tree, ok := ll1.Analysis(inputdata)

	treeString := []string{}
	if ok {
		q := queue.NewQueue()
		q.PushBack(tree)
		for q.Len() > 0 {
			cur := q.Pop().(*gramma_analysis.AnalysisTreeNode)
			for i := range cur.Next {
				treeString = append(treeString, fmt.Sprintf("%s(%d) -> %s(%d)", cur.Value, cur.ID, cur.Next[i].Value, cur.Next[i].ID))
				q.PushBack(cur.Next[i])
			}
		}
	}

	b, err := json.Marshal(map[string]interface{}{
		"first_set":      ll1.GetFirstSet(),
		"follow_set":     ll1.GetFollowSet(),
		"analysis_table": ll1.GetAnalysisTable(),
		"tree":           treeString,
		"analysis":       table,
		"analysis_ok":    ok,
	})
	if err != nil {
		log.Println(err)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Write(b)

}
