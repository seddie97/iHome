package main

import "fmt"

type ClueBizTag struct {
	Value *TagValue `thrift:"value,1,required" json:"value"`
}

type TagValue struct {
	SingleTagValue *string  `thrift:"single_tag_value,1" json:"single_tag_value,omitempty"`
	MultiTagValue  []string `thrift:"multi_tag_value,2" json:"multi_tag_value,omitempty"`
}

func main() {
	bizTags := make(map[string]*ClueBizTag, 10)
	test := "123123"
	bizTags["123123"] = &ClueBizTag{Value: &TagValue{SingleTagValue: &test}}
	//bizTags["123123"].Value.SingleTagValue = &test
	fmt.Println(bizTags)
	fmt.Println(*bizTags["123123"])
	fmt.Println(*bizTags["123123"].Value)
	fmt.Println(*bizTags["123123"].Value.SingleTagValue)
}
