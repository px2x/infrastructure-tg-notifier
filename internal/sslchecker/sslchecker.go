package sslchecker

import (
	"crypto/tls"
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
	//fmt.Printf("Issuer: %s\nExpiry: %v\n", conn.ConnectionState().PeerCertificates[0].Issuer, expiry.Format(time.RFC850))
	if time.Now().Before(expiry.Add(time.Duration(-24) * time.Hour)) {
		return true, expiry.Format("2.1.2006")
	} else {
		return false, expiry.Format("2.1.2006")
	}

}

func CheckSSLEnv(service *config.Services, isSchedulerCheck bool) (message string, doSendReport bool) {
	firstString := "Проверили SSL сервисов!\n\n"
	resultString := ""
	doSendReport = false
	result := true
	if isSchedulerCheck == false || service.LastCheckSSL.Add(service.CheckIntervalSSL).Before(time.Now()) {
		service.LastCheckSSL = time.Now()
		for _, env := range service.Env {
			resultString += "Окружение: <strong>" + env.Name + "</strong>\n"
			for _, link := range env.Link {
				status, expiry := CheckSSL(link.Url)
				if status {
					resultString += "✅"
				} else {
					resultString += "💩"
					result = false
				}
				resultString += " - " + link.Url + "\n(истекает " + expiry + ")"
				resultString += "\n"

			}
			resultString += "\n"
		}

	}

	if result == false && isSchedulerCheck == true {
		service.LastCheckSSL = time.Now().Add(service.MuteTimeAfterError)
		//service.LastCheckSSL = time.Now().Add(time.Duration(120) * time.Second)
		firstString = "Досиделись блэть?!\n\nПора обновить почти протухшие SSL!\n" +
			"или там вообще уже нет SSL - предприми что нибудь! 🤬\n\n"
		doSendReport = true
	}
	return firstString + resultString, doSendReport
}
