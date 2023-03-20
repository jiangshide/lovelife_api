# lovelife_api
this is api for lovelife



短信测试:
package main

import (
        "fmt"

        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
        sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

func main() {

        credential := common.NewCredential(
                "AKIDyGPMsE9apkW3RnLD4mcD1conXz9Dv3IK",
                "4sOfkYobliU1GFzSeZOwDCtqHRM1SeBc",
        )
        cpf := profile.NewClientProfile()
        cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
        client, _ := sms.NewClient(credential, "", cpf)
        
        request := sms.NewSendSmsRequest()

        params := "{\"PhoneNumberSet\":[\"+8618311271399\"],\"TemplateID\":\"718159\",\"Sign\":\"梵山科技\",\"TemplateParamSet\":[\"23423\",\"10\"],\"SmsSdkAppid\":\"1400425697\"}"
        err := request.FromJsonString(params)
        if err != nil {
                panic(err)
        }
        response, err := client.SendSms(request)
        if _, ok := err.(*errors.TencentCloudSDKError); ok {
                fmt.Printf("An API error has returned: %s", err)
                return
        }
        if err != nil {
                panic(err)
        }
        fmt.Printf("%s", response.ToJsonString())
} 



{
  "Response": {
    "SendStatusSet": [
      {
        "SerialNo": "2019:-7601187759052096569",
        "PhoneNumber": "+8618311271399",
        "Fee": 1,
        "SessionContext": "",
        "Code": "Ok",
        "Message": "send success",
        "IsoCode": "CN"
      }
    ],
    "RequestId": "036b8ef9-7083-483c-a20d-cef3c3fd15db"
  }
}




Server:nginx

Date:Sat, 12 Sep 2020 16:50:33 GMT

Content-Type:application/json

Content-Length:235

Connection:close


		
curl -X POST https://sms.tencentcloudapi.com -H "Authorization: TC3-HMAC-SHA256 Credential=AKIDyGPMsE9apkW3RnLD4mcD1conXz9Dv3IK/2020-09-12/sms/tc3_request, SignedHeaders=content-type;host, Signature=4cc4a74fe0cc8c9daaca510f5d2f6a171eaf157a2f1053e21f4706d93a5d5d81" -H "Content-Type: application/json" -H "Host: sms.tencentcloudapi.com" -H "X-TC-Action: SendSms" -H "X-TC-Timestamp: 1599929433" -H "X-TC-Version: 2019-07-11" -H "X-TC-Language: zh-CN" -d '{"PhoneNumberSet":["+8618311271399"],"TemplateID":"718159","Sign":"梵山科技","TemplateParamSet":["23423","10"],"SmsSdkAppid":"1400425697"}'
