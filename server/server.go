package server

import (
	"io"
	"log"
	"net/http"
	"os"
)

func loadIndex(index string) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		f, err := os.Open(index)
		if err != nil {
			log.Println(err)
			writer.Write([]byte(err.Error()))
			return
		}

		io.Copy(writer, f)
	}
}

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/nfa", loadIndex("./server/nfa.html"))
	mux.HandleFunc("/lexical", loadIndex("./server/lexical.html"))
	mux.HandleFunc("/grammar", loadIndex("./server/grammar.html"))
	mux.HandleFunc("/compiler", loadIndex("./server/compiler.html"))
	mux.HandleFunc("/api/grammar", submitGrammar)
	mux.HandleFunc("/api/lexical/nfa", submitNFA)
	mux.HandleFunc("/api/lexical/explain", submitProgram)
	mux.HandleFunc("/api/compiler", compilerSubmitProgram)
	mux.HandleFunc("/", loadIndex("./server/index.html"))

	port := "8102"

	log.Println("Server Listen at port " + port + " ... ")
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		panic(err)
	}
}
