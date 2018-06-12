package server

import (
	"Compiler/lexical_analysis"
	"encoding/json"
	"log"
	"net/http"
)

func submitNFA(writer http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		log.Println(err)
		writer.Write([]byte(err.Error()))
		return
	}

	regexp := request.FormValue("regexp")

	n := lexical_analysis.NewNFA(regexp)
	d := lexical_analysis.NewDFA(n)

	b, err := json.Marshal(map[string]interface{}{
		"nfa":       n.String(),
		"dfa_table": d.String(),
	})
	if err != nil {
		log.Println(err)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Write(b)

}
