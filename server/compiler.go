package server

import (
	"github.com/Chasiny/simple_compiler/simple_compiler"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func compilerSubmitProgram(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Println(err)
		writer.Write([]byte(err.Error()))
		return
	}

	program := request.FormValue("program")

	fmt.Println(program)

	c := simple_compiler.NewSimpleCompiler()

	res, err := c.Run(program)

	b, err := json.Marshal(map[string]interface{}{
		"words": res,
		"err":   err,
	})
	writer.Write(b)

}
