package availability

import (
	"log"
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
		log.Fatalln(err)
	}

	if response.StatusCode == 200 {
		return true
	}
	return false

}
