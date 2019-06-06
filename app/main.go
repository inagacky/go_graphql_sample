package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/inagacky/go_graphql_sample/app/application/db"
	"github.com/inagacky/go_graphql_sample/app/application/graphql_util"
	"io/ioutil"
	"net/http"
)

func main() {

	// DB接続
	db.Init()

	// graphqlの受付
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// bodyの読み取り
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		// query 実行
		result := executeQuery(fmt.Sprintf("%s", body), graphql_util.Schema)
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)
}


func executeQuery(query string, schema graphql.Schema) *graphql.Result {

	// schemaとqueryで処理実行
	result := graphql.Do(graphql.Params {
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
