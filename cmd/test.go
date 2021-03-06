package main

import (
	"encoding/json"
	"fmt"
	"time"
	Server "web-token-server/internal/server"
	"web-token-server/pkg/generator"
	"web-token-server/testing"
)

func main() {
	for j := 0; j < 50000000000; j++ {
		for i := 0; i < 100; i++ {
			go test(j * 90)
			//fmt.Println(i)
		}
		fmt.Println(j * 100)
		time.Sleep(1 * time.Second)
	}

}

func test(count int) {
	name := ""
	name = generator.RandStringRunes(5)
	//fmt.Println("New name Generate", name)
	newToken := testing.TestCreateNewToken("http://127.0.0.1:8080/new", name, "POST")
	token := Server.TokenStruct{}
	_ = json.Unmarshal(newToken, &token)
	//fmt.Println("Accept", token.Token)
	checkToken := Server.ValidateStruct{}
	valid := testing.TestTokenCheckInfo("http://127.0.0.1:8080/check", token.Token, "POST")
	_ = json.Unmarshal(valid, &checkToken)
	if checkToken.Valid {
		info := Server.TokenInfoStruct{}
		tokenInfo := testing.TestTokenCheckInfo("http://127.0.0.1:8080/info", token.Token, "POST")
		_ = json.Unmarshal(tokenInfo, &info)
		if name == info.Name {
			fmt.Println("Test", count, "OK")
		}
	}

}
