package availability

import (
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
