package availability

import (
	"github.com/px2x/infrastructure-tg-notifier/config"
	"net/http"
	"strconv"
	"time"
)

func CheckAvailability(url string) (bool, int) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	request, _ := http.NewRequest(
		"GET",
		url,
		nil)

	response, err := client.Do(request)
	if err != nil {
		return false, 0
	}

	if response.StatusCode == 200 {
		return true, response.StatusCode
	}
	return false, response.StatusCode

}

func CheckAvailabilityEnv(service *config.Services, isSchedulerCheck bool) (message string, doSendReport bool) {
	firstString := "Проверили доступность сервисов!\n\n"
	resultString := ""
	doSendReport = false
	result := true
	if isSchedulerCheck == false || service.LastCheckSite.Add(service.CheckIntervalSite).Before(time.Now()) {
		service.LastCheckSite = time.Now()
		for _, env := range service.Env {
			resultString += "Окружение: <strong>" + env.Name + "</strong>\n"
			for _, link := range env.Link {
				status, code := CheckAvailability(link.Url)
				if status {
					resultString += "✅"
				} else {
					resultString += "💩"
					result = false
				}
				resultString += " - " + link.Url + " (" + strconv.Itoa(code) + ")"
				resultString += "\n"

			}
			resultString += "\n"
		}

	}

	if result == false && isSchedulerCheck == true {
		service.LastCheckSite = time.Now().Add(time.Duration(60) * time.Second)
		firstString = "Буэнос диас нахуй!\n\nЕсть проблема с доступностью хостов!\n\n"
		doSendReport = true
	}
	return firstString + resultString, doSendReport
}
