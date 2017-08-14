package main

import (
	".."
)

func main() {
	// req1 := ReqPayPageInfo{110, 44444, "This is qua number xxxx", "boss channel ", "user type", []Item{{"lewiskong", 21}, {"fuck", 22}}}
	// req2 := ReqCheckUserPayValid{5555, "user type", "userid", "1.1.1.1", "appid", "open id", "access token", "vuserid", "vusession"}

	// req, ok := types.Extract(reflect.TypeOf(ReqGetValidPay{}), req1, req2).(ReqGetValidPay)
	// if !ok {
	// 	fmt.Println("err")
	// }

	// types.Println(req)
	// fmt.Println(req)
	types.PrintlnInOneLine(req)
}
