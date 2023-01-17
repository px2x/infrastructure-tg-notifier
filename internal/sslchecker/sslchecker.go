package sslchecker

import (
	"crypto/tls"
	"fmt"
	"github.com/px2x/infrastructure-tg-notifier/config"
	"net/url"
	"time"
)

func CheckSSL(uri string) (bool, string) {

	u, err := url.Parse(uri)
	conn, err := tls.Dial("tcp", u.Hostname()+":443", nil)
	if err != nil {
		return false, "no such host"
	}

	err = conn.VerifyHostname(u.Hostname())
	if err != nil {
		return false, "no SSL"
	}
	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	fmt.Printf("Issuer: %s\nExpiry: %v\n", conn.ConnectionState().PeerCertificates[0].Issuer, expiry.Format(time.RFC850))
	if time.Now().Before(expiry) {
		return true, expiry.Format("2.1.2006")
	} else {
		return false, expiry.Format("2.1.2006")
	}

}

func CheckSSLEnv(service *config.Services) string {
	resultString := "Проверили SSL сервисов!\n\n"
	for _, env := range service.Env {
		resultString += "Окружение: <strong>" + env.Name + "</strong>\n"
		for _, link := range env.Link {
			status, expiry := CheckSSL(link.Url)
			if status {
				resultString += "✅"
			} else {
				resultString += "💩"
			}
			resultString += " - " + link.Url + " (" + expiry + ")"
			resultString += "\n"
		}
		resultString += "\n"
	}
	return resultString
}
