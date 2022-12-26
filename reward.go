package main
import (
"encoding/json"
"fmt"
"time"
"io/ioutil"
"net/http"
"strconv"
"net/smtp"
)

type Earnings struct {
FilAmount float64
Timestamp string
}

type Data struct {
Earnings []Earnings
}

const (
    // 按照下面时间间隔发送邮件
    SendInterval = 60
)



func main() {
    sendEmailFunc := sendEmail

    // Initialize the remainingTime variable with a zero value.
    remainingTime := time.Time{}

    // Print the remaining time until the next email is sent.
    fmt.Printf("%s remaining until the next email is sent.\n", time.Until(remainingTime).Round(time.Minute).String())
    // Create a ticker that will count down the time until the next email is sent.
    ticker := time.Tick(time.Duration(SendInterval) * time.Minute)
for {
        // Wait for the next interval.
        remainingTime = <-ticker

        // Print the remaining time until the next email is sent.
    fmt.Printf("%s remaining until the next email is sent.\n", time.Until(remainingTime).Round(time.Minute).String())
        // Send the email.
        err := sendEmailFunc()
        if err != nil {
            fmt.Println(err)
            continue
        }

        fmt.Println("Email sent successfully!")
    }
}



func sendEmail() error{
currentTime := time.Now()
currentTimestamp := currentTime.Unix()
currentTimestampInt := int(currentTimestamp)
currentTimestampStr := strconv.Itoa(currentTimestampInt)
// 打印一些信息,判断程序异常
fmt.Println("The current timestamp is:", currentTimestamp)
fmt.Println("Change ago The current timestamp is:", currentTimestampStr)
fmt.Println("https://uc2x7t32m6qmbscsljxoauwoae0yeipw.lambda-url.us-west-2.on.aws/?filAddress=f1knmlcd4ha24te7lxfn2jynniaf...............&startDate=1670945820000&endDate=" + currentTimestampStr + "000" + "&step=hour")


//使用get请求,从地址网页源码中获取收益信息
resp, err := http.Get("https://uc2x7t32m6qmbscsljxoauwoae0yeipw.lambda-url.us-west-2.on.aws/?filAddress=f1knmlcd4ha24te7lxfn2jynnia.........&startDate=1670945820000&endDate=" + currentTimestampStr + "000" + "&step=hour")

if err != nil {
    fmt.Println(err)
    return nil
}

defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
    fmt.Println(err)
    return nil
}
str := string(body)
var data Data
json.Unmarshal([]byte(str), &data)

//-=============================================输出output与output1=======================
var output string
var output1 string
var totalFilAmount float64

//遍历data.Earnings中的每一个元素
for _, v := range data.Earnings {
    output += fmt.Sprintf("filAmount: %f, timestamp: %s\n", v.FilAmount, v.Timestamp)
    totalFilAmount += v.FilAmount
}

//计算平均值
var avgFilAmount float64
avgFilAmount = totalFilAmount / float64(len(data.Earnings))

//将平均值和总和添加到output1中
output1 = fmt.Sprintf("avgFilAmount: %f, totalFilAmount: %f", avgFilAmount, totalFilAmount)

//打印output和output1
fmt.Println(output)
fmt.Println(output1)

//====================================================发送邮件======================
// 使用 smtp.PlainAuth 函数创建一个身份验证对象
auth := smtp.PlainAuth("", "perfectycc@126.com", "这里是此邮箱的IMAP_KEY", "smtp.126.com")
to := []string{"perfectyc@126.com"}
// 构造邮件正文
msg := []byte("To: perfectyc@126.com\r\n" + "Subject: Saturn_CDN收益报告\r\n" + "\r\n" + output + "\r\n" + output1 + "\r\n")
// 使用 smtp.SendMail 函数发送邮件
err = smtp.SendMail("smtp.126.com:25", auth, "perfectycc@126.com", to, msg)
if err != nil {
fmt.Printf("send mail error: %v", err) // 如果有错误，就会打印错误信息
} else {
fmt.Println("send mail success!") // 如果没有错误，就说明邮件发送成功
}
return nil
}
