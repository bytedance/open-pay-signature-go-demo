// Package utils sign_test.go 签名算法的单测
//
// Package utils sign_test.go Unit testing of signature algorithms

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_RequestSign 测试担保支付请求签名算法
//
// Test_RequestSign Testing the Guaranteed Payment Request Signature Algorithm
func Test_RequestSign(t *testing.T) {
	// 以预下单接口为例
	//
	// Take the create_order interface as an example
	paramsMap := map[string]interface{}{
		"app_id":        "ttcfdbb9xxxxxxxxxxx",
		"thirdparty_id": "tta4bad200000xxxxxx",
		"out_order_no":  "test-02167569xxxxxx",
		"total_amount":  2376,
		"subject":       "test-payment_subject-test-paym...",
		"body":          "强烈推荐！经典腊汁肉夹馍团购价仅需7.92元！",
		"valid_time":    172800,
		"notify_url":    "https://www.xxx.com",
		"disable_msg":   0,
		"msg_page":      "pages/user/orderDetail/orderDetail?id=997979879879879879",
		"sign":          "edc608b160a1be3de0xxxxxx",
	}

	assert.Equal(t, "54f102e7115f8a6a3e6af4633dc33959", RequestSign(paramsMap))
}

// Test_RequestSign 测试担保支付回调签名算法
//
// Test_RequestSign Test Guaranteed Payment Callback Signature Algorithm
func Test_CallbackSign(t *testing.T) {
	// 以支付回调为例
	//
	// Take the payment callback as an example
	callbackToken := "fdsifakhflasjfxxxxxxxxx" // callbackToken 是平台上配置的token (callbackToken is the token configured on the platform)
	timestamp := "1652675265"
	nonce := "9999"
	msg := "80850852"
	sortedString := []string{callbackToken, timestamp, nonce, msg}
	assert.Equal(t, "c9df04a40645c4ec15c13bc542cea589eac57e64", CallbackSign(sortedString))
}
