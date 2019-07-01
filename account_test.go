package goanda_test

import (
	"fmt"
	"os"

	"github.com/stewi1014/goanda"

	"gopkg.in/yaml.v2"
)

func getAccountAndToken() (string, string) {
	f, _ := os.Open("token.yaml")
	d := yaml.NewDecoder(f)
	a := struct {
		Token     string
		AccountID string
	}{}
	d.Decode(&a)
	return a.AccountID, a.Token
}

func ExampleConnection_GetAccount() {
	id, token := getAccountAndToken()

	conn, err := goanda.NewConnection(id, token, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := conn.Request("/accounts/" + id + "/summary")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
	//Output:
}
