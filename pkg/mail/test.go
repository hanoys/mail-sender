package mail

import (
	"fmt"
	"net/smtp"
	"os"
)

const (
	from     = "hanoys@yandex.ru"
	password = "hjvmfstkbbvtrgwh"
	host     = "smtp.yandex.com"
	port     = "587"
	msg      = "Test Message"
)

func test() {
	auth := smtp.PlainAuth("", from, password, host)
	body := []byte(msg)
	receivers := []string{"thehanoy@gmail.com"}
	err := smtp.SendMail(host+":"+port, auth, from, receivers, body)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Println("Message has sent")
}
