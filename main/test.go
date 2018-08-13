package main

import (
	"io/ioutil"
	"fmt"
	"../email"
	"bytes"
)
func getContent() string {
	var buffer bytes.Buffer
	buffer.WriteString("\n任务总数: ")
	buffer.WriteString("83")
	buffer.WriteString("<br>")
	buffer.WriteString("\n成功数: ")
	buffer.WriteString("<font color='green'>")
	buffer.WriteString("82")
	buffer.WriteString("</font><br>")
	buffer.WriteString("\n失败数: ")
	buffer.WriteString("<font color='red'>")
	buffer.WriteString("1")
	buffer.WriteString("</font><br>")
	buffer.WriteString("\n总开始时间: ")
	buffer.WriteString(" 2018-08-02 04:15:31")
	buffer.WriteString("<br>")
	buffer.WriteString("\n总结束时间: ")
	buffer.WriteString("2018-08-02 06:58:52")
	buffer.WriteString("<br>")
	buffer.WriteString("\n总用时: ")
	buffer.WriteString("2h : 43m : 20s")
	buffer.WriteString("<br>")
	buffer.WriteString("\n任务明细: ")
	buffer.WriteString("<br>")

	buffer.WriteString("\n<table bordercolor=\"#008888\" style=\"BORDER-COLLAPSE: collapse\" border=1>")

	buffer.WriteString("<tr>")
	buffer.WriteString("<td>")
	buffer.WriteString("任务名称")
	buffer.WriteString("</td>")
	buffer.WriteString("<td>")
	buffer.WriteString("状态")
	buffer.WriteString("</td>")
	buffer.WriteString("<td>")
	buffer.WriteString("开始时间")
	buffer.WriteString("</td>")
	buffer.WriteString("<td>")
	buffer.WriteString("结束时间")
	buffer.WriteString("</td>")
	buffer.WriteString("<td>")
	buffer.WriteString("用时（m:s）")
	buffer.WriteString("</td>")
	buffer.WriteString("</tr>")

	buffer.WriteString("<tr>")
	buffer.WriteString("<td>")
	buffer.WriteString("cube.job.user.a.SjUserCubeJob")
	buffer.WriteString("</td>")
	buffer.WriteString("<td>")
	buffer.WriteString("<font color='green'>成功</font>")
	buffer.WriteString("</td>")
	buffer.WriteString("<td>")
	buffer.WriteString("2018-08-02 04:15:31")
	buffer.WriteString("</td>")
	buffer.WriteString("<td>")
	buffer.WriteString("2018-08-02 04:17:11")
	buffer.WriteString("</td>")
	buffer.WriteString("<td>")
	buffer.WriteString("1m:40s")
	buffer.WriteString("</td>")
	buffer.WriteString("</tr>")

	buffer.WriteString("</table>")

	content := buffer.String()
	return content
}


func main() {
	toMails := []string{"dongzhijun0314@163.com"}
	fromNick := "离线任务通知"
	title := "邀请有礼201802"
	content := getContent()
	bytesMap := make(map[string][]byte)

	filename := "text.txt"
	attachment, err := ioutil.ReadFile("/Users/dongzj/Downloads/text.txt")
	if err != nil {
		fmt.Printf("get attachment error: %v", err)
	}
	bytesMap[filename] = attachment

	rc, msg := email.SendMailWithATTs(fromNick, toMails, title, content, bytesMap)
	fmt.Println(rc, msg)
}
