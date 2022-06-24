package utils

import (
	"encoding/json"
	"net/http"
	"regexp"
)

var (
	allCarRe  = regexp.MustCompile("[ ./-]")
	allNumRe  = regexp.MustCompile("[^0-9]")
	allSame14 = regexp.MustCompile("0{14}|1{14}|2{14}|3{14}|4{14}|5{14}|6{14}|7{14}|8{14}|9{14}")
	allSame11 = regexp.MustCompile("0{11}|1{11}|2{11}|3{11}|4{11}|5{11}|6{11}|7{11}|8{11}|9{11}")
)

func RespondWithError(w http.ResponseWriter, code, errorCode int, message string) {
	RespondWithJSON(w, code, map[string]interface{}{"error": message, "code": errorCode})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
