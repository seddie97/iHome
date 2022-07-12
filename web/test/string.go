package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Mediaiplocation struct {
	Subdivisions []string `json:"subdivisions"`
	Country      string   `json:"country"`
}

func main() {
	iplocation := &Mediaiplocation{
		Subdivisions: []string{},
		Country:      "中国",
	}
	res, err := json.Marshal(iplocation)
	if err != nil {
		return
	}
	resStr := string(res)
	fmt.Println(resStr)
	resStr = strings.ReplaceAll(resStr, "Subdivisions", "subdivisions")
	resStr = strings.ReplaceAll(resStr, "Country", "country")
	fmt.Println(resStr)
}
