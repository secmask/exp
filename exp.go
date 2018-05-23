package exp

import (
	"log"
	"encoding/json"
)
// CheckErr panic if error
func CheckErr(err error) {
	if err!=nil{
		log.Panic(err)
	}
}
// Dump
func Dump(o interface{}){
	d,_ := json.MarshalIndent(o,""," ")
	log.Println(string(d))
}
// MustDump
func MustDump(o interface{}, err error){
	CheckErr(err)
	Dump(o)
}