package generator

import (
	"crypto/rand"
	"fmt"
	"log"
)

func UUIDV4() (token string) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Println("Error generate uuid: ", err)
		return
	}

	token = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return
}
