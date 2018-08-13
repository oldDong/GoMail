golang 实现的简单邮件发送服务，使用原生smtp包实现，也包含用gomail.v2实现的版本。
原生smtp实现的时候有以下几个注意点
1、如果发送者（From）开通了TLS(SSL)，端口时465时，用原生的smtp.SendMail，在内部调用smtp.newClient（）会卡住，最终返回EOF,通过tlsmail.go替换
原生的方法，可以实现邮件的发送。

2、在组织邮件的内容时，分隔符boundary=\"XXX\"，附件名filename=\"XXX\",双引号不能少

3、邮件的发送者From、Subject、filename部分如果包含中文，需要指定字符集和编码方式。在实际测试中发现，如果不指定，网易邮箱大师可以正常使用，
但在Windows系统中，foxmail 7.2版本邮件会乱码或没有附件（Mac上没有问题。。。。）
