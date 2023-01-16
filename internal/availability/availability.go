package availability

import (
	"github.com/px2x/infrastructure-tg-notifier/config"
	"net/http"
)

func CheckAvailability(url string) bool {
	client := &http.Client{}
	request, _ := http.NewRequest(
		"GET",
		url,
		nil)

	response, err := client.Do(request)
	if err != nil {
		return false
	}

	if response.StatusCode == 200 {
		return true
	}
	return false

}

func CheckAvailabilityEnv(service *config.Services) string {
	resultString := ""
	for _, env := range service.Env {

		for _, link := range env.Link {
			if CheckAvailability(link.Url) {

			} else {

			}

		}
	}
	return resultString
}
