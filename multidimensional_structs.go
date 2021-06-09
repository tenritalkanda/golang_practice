package main

import (
	"encoding/json"
	"fmt"
)

type Resp struct {
	Spf Param `json:"spf"`
}

type Param struct {
	Is_verified    int            `json:"is_verified"`
	Spf_matched    int            `json:"spf_matched"`
	Required_value string         `json:"required_value"`
	Current_value  map[int]string `json:"current_value"`
}

func main() {

	str := make(map[int]string)

	str[0] = "v=spf1 a -all,"
	resp := Resp{Spf: Param{Is_verified: 0, Spf_matched: 0, Required_value: "v=spf1 a  include:32782.pppp.com ~all", Current_value: str}}
	js, _ := json.Marshal(resp)
	fmt.Printf("%s", js)
}
