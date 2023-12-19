package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetOrderNumber(t *testing.T) {

	sss()

}

func sss() {
	str := `{\"code\":0,\"msg\":\"success\",\"data\":{\"payQr\":\"https://open.huilianpay.com/output/openId?redirectUrl=https://s.hxt5.cn/toyPre.html?attach=ZZ9KDT845Z,38,585945537024963,1989,1226862,1369596012470\",\"agencyNo\":\"1226862\",\"money\":1989,\"hlMerchantId\":\"1369596012470\",\"orderId\":585945537024963},\"ServerTime\":\"2023-12-20 14:18:29\"}`

	// 替换所有的转义符
	str = strings.ReplaceAll(str, "\\", "")

	//var jsonObj map[string]interface{}
	//err := json.Unmarshal([]byte(str), &jsonObj)
	//if err != nil {
	//	fmt.Println("JSON unmarshal error:", err)
	//	return
	//}

	fmt.Println(str)

}
