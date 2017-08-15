package main

import (
	".."
	"fmt"
	// "reflect"
)

func main() {
	req1 := ReqPayPageInfo{110, 44444, "This is qua number xxxx", "boss channel ", "user type"}
	req2 := ReqCheckUserPayValid{5555, "user type", "userid", "1.1.1.1", "appid", "open id", "access token", "vuserid", "vusession"}

	// req, ok := types.Extract(reflect.TypeOf(ReqGetValidPay{}), req1, req2).(ReqGetValidPay)
	ret := ReqGetValidPay{}
	req, ok := types.Extract(ret, req1, req2).(ReqGetValidPay)
	if !ok {
		fmt.Println("err")
	}

	// types.Println(req)

	types.PrintlnInOneLine(req)
	fmt.Println("————————————————————————————————————————————————")
	types.PrintlnInOneLine(ret)

}
