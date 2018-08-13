package email

import (
	"net/smtp"
	"bytes"
	"strings"
	"fmt"
	"encoding/base64"
)

const (
	tencentFromUser         = "**********@qq.com"
	tencentFromUserPassward = "**********"
	tencentHost             = "smtp.qq.com"
	tencentPort             = ":25"
)

/**
	根据邮件服务器选择不同的发送方式
	如果服务器开通了TLS(SSL),使用SendMailUsingTLS,否则使用原生方法
	如果tencentPort为25，选择smtp.SendMail;如果端口为465，选择SendMailUsingTLS
 */
func send(to []string, b []byte) error {
	auth := smtp.PlainAuth("", tencentFromUser, tencentFromUserPassward, tencentHost)
	//return SendMailUsingTLS(tencentHost+tencentPort, auth, tencentFromUser, to, b)
	return smtp.SendMail(tencentHost+tencentPort, auth, tencentFromUser, to, b)
}

func SendMailWithATTs(fromNick string, toMails []string, title string, content string, bytesMap map[string][]byte) (int, string) {
	mime := bytes.NewBuffer(nil)
	to := strings.Join(toMails, ",")
	boundary := "=_NextPart_="

	//首部
	mime.WriteString(fmt.Sprintf("From: =?UTF-8?B?%s?=<%s>\r\n", base64.StdEncoding.EncodeToString([]byte(fromNick)), tencentFromUser))
	mime.WriteString(fmt.Sprintf("To: %s\r\n", to))
	mime.WriteString(fmt.Sprintf("Subject: =?UTF-8?B?%s?=\r\n", base64.StdEncoding.EncodeToString([]byte(title))))
	mime.WriteString(fmt.Sprintf("Content-Type: multipart/mixed;charset=UTF-8; boundary=\"%s\"\r\n", boundary))
	mime.WriteString("MIME-Version: 1.0\r\n")

	//正文
	mime.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	mime.WriteString("Content-Type: text/html;\r\n")
	mime.WriteString(content)
	mime.WriteString("\r\n\r\n")

	//附件
	for filename := range bytesMap {
		mime.WriteString(fmt.Sprintf("--%s\r\n", boundary))
		mime.WriteString("Content-Type: application/octet-stream;\r\n")
		mime.WriteString("Content-Transfer-Encoding: base64\r\n")
		mime.WriteString("Content-Disposition: attachment; filename=\"=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(filename)) + "?=\"" + "\r\n\r\n")

		b := make([]byte, base64.StdEncoding.EncodedLen(len(bytesMap[filename])))
		base64.StdEncoding.Encode(b, bytesMap[filename])
		mime.Write(b)
		mime.WriteString("\r\n\r\n")
	}
	mime.WriteString("--" + boundary + "--")

	fmt.Println("SendMailWithATTs邮件报文：")
	fmt.Println(mime.String())

	error := send(toMails, mime.Bytes())
	if error != nil {
		return 1, error.Error()
	}
	return 0, ""
}
