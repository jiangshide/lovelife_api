package main

// import (
// 	"sanskrit_api/tencentyun"
// 	"fmt"
// )

// const (
// 	sdkappid = 1400287594
// 	key = "72113ec138c0a5b05ab5865fbe278b23627dafa8dfbdb6764327f639bc00e709"
// )

// func main()  {
// 	sig, err := tencentyun.GenSig(sdkappid, key, "administrator", 86400*180)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println(sig)
// 	}
// 	sig, err = tencentyun.GenSigWithUserBuf(sdkappid, key, "administrator", 86400*180, []byte("abc"))
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println(sig)
// 	}
// }
