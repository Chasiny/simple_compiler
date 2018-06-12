package server

import (
	"Compiler/lexical_analysis"
	"encoding/json"
	"log"
	"net/http"
)

func submitProgram(writer http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		log.Println(err)
		writer.Write([]byte(err.Error()))
		return
	}

	program := request.FormValue("program")

	l:=lexical_analysis.NewLexicalAnalysisDemo()
	words := l.Explain(program)

	b, err := json.Marshal(map[string]interface{}{
		"words": l.WordsToString(words),
	})
	if err != nil {
		log.Println(err)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Write(b)

}
