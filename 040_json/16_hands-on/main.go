package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type code struct {
	Code int
	Desc string
}

func main() {
	var data []code

	rcvd := `[{"Code":200,"Desc":"StatusOK"},{"Code":301,"Desc":"StatusMovedPermanently"},{"Code":302,"Desc":"StatusFound"},{"Code":303,"Desc":"StatusSeeOther"},{"Code":307,"Desc":"StatusTemporaryRedirect"},{"Code":400,"Desc":"StatusBadRequest"},{"Code":401,"Desc":"StatusUnauthorized"},{"Code":402,"Desc":"StatusPaymentRequired"},{"Code":403,"Desc":"StatusForbidden"},{"Code":404,"Desc":"StatusNotFound"},{"Code":405,"Desc":"StatusMethodNotAllowed"},{"Code":418,"Desc":"StatusTeapot"},{"Code":500,"Desc":"StatusInternalServerError"}]`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range data {
		fmt.Println(v.Code, "-", v.Desc)
	}
}
