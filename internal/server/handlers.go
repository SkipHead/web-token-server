package Server

import (
	"encoding/json"
	"log"
	"net/http"
)

//NewToken Handler function accept post request for generate new token from json body
func NewToken(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := TokenName{}
		decoder := json.NewDecoder(r.Body)
		errDecoder := decoder.Decode(&name)
		if errDecoder != nil {
			log.Println("New token request", errDecoder)
		}
		sessionToken := TokenStruct{
			Token: newToken(name.Name, r.Host),
		}
		w.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(sessionToken)
		if err != nil {
			log.Println("Handle function newToken token encoder:", err)
		}
	} else if r.Method == "GET" {
		w.WriteHeader(400)
	}
}

//ChekToken Handler for check on valid token
func ChekToken(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		token := TokenStruct{}
		decoder := json.NewDecoder(r.Body)
		errDecoder := decoder.Decode(&token)
		if errDecoder != nil {
			log.Println("New token request", errDecoder)
		}
		sessionToken := ValidateStruct{Valid: isExpired(token.Token)}
		w.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(sessionToken)
		if err != nil {
			log.Println("Handle function ChekToken token encoder:", err)
		}
	} else if r.Method == "GET" {
		w.WriteHeader(400)
	}
}

//TokenInfo Handler return about info for token string and name
func TokenInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		token := TokenStruct{}
		decoder := json.NewDecoder(r.Body)
		errDecoder := decoder.Decode(&token)
		if errDecoder != nil {
			log.Println("New token request", errDecoder)
		}

		tokenInfo := TokenInfoStruct{}
		for _, t := range storeTokens {
			if t.Token == token.Token {
				tokenInfo.Token = token.Token
				tokenInfo.Name = t.Name
			}
		}
		tokenInfo.Token = token.Token
		if tokenInfo.Name == "" {
			tokenInfo.Name = "Not found"
		}

		w.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(tokenInfo)
		if err != nil {
			log.Println("Handle function tokenInfo encoder:", err)
		}
	} else if r.Method == "GET" {
		w.WriteHeader(400)
	}
}

//Version return about version for server
func Version(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(version())
	if err != nil {
		log.Println("Handle function version encoder:", err)
	}

}
