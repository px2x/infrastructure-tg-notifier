package selectel

import (
	"encoding/json"
	"github.com/bojanz/currency"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type billingResponse struct {
	Status string `json:"status"`
	Data   struct {
		Currency  string `json:"currency"`
		IsPostpay bool   `json:"is_postpay"`
		Discount  int    `json:"discount"`
		Primary   struct {
			Main  int `json:"main"`
			Bonus int `json:"bonus"`
			VkRub int `json:"vk_rub"`
			Ref   int `json:"ref"`
			Hold  struct {
				Main  int `json:"main"`
				Bonus int `json:"bonus"`
				VkRub int `json:"vk_rub"`
			} `json:"hold"`
		} `json:"primary"`
		Storage struct {
			Main       int         `json:"main"`
			Bonus      int         `json:"bonus"`
			VkRub      int         `json:"vk_rub"`
			Prediction interface{} `json:"prediction"`
			Debt       int         `json:"debt"`
			Sum        int         `json:"sum"`
		} `json:"storage"`
		Vpc struct {
			Main       int         `json:"main"`
			Bonus      int         `json:"bonus"`
			VkRub      int         `json:"vk_rub"`
			Prediction interface{} `json:"prediction"`
			Debt       int         `json:"debt"`
			Sum        int         `json:"sum"`
		} `json:"vpc"`
		Vmware struct {
			Main       int         `json:"main"`
			Bonus      int         `json:"bonus"`
			VkRub      int         `json:"vk_rub"`
			Prediction interface{} `json:"prediction"`
			Debt       int         `json:"debt"`
			Sum        int         `json:"sum"`
		} `json:"vmware"`
		WithdrawDay interface{} `json:"withdraw_day"`
	} `json:"data"`
}

func checkBilling(apiKey string) (primary int, vpc int, storage int) {

	client := &http.Client{}
	request, _ := http.NewRequest(
		"GET",
		"https://api.selectel.ru/v3/billing/balance",
		nil,
	)
	request.Header.Set("X-token", apiKey)

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	var data billingResponse
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to go struct pointer
		log.Println("Can not unmarshal JSON")
	}
	return data.Data.Primary.Main, data.Data.Vpc.Main, data.Data.Storage.Main
}

func CheckBillingMessage(apiKey string) string {
	primary, vpc, storage := checkBilling(apiKey)
	resultString := "Биллинг Selectel\n\n" +
		setIndicator(vpc, 2000) + " Облачная платформа: " + converCurrency(vpc) + "\n" +
		setIndicator(storage, 500) + " Хранилище: " + converCurrency(storage) + "\n" +
		setIndicator(primary, 1000) + " Основной баланс: " + converCurrency(primary)
	return resultString
}

func converCurrency(value int) string {
	locale := currency.NewLocale("ru")
	formatter := currency.NewFormatter(locale)
	formatter.MaxDigits = 2
	amount, _ := currency.NewAmount(strconv.Itoa(value/100)+"."+strconv.Itoa(value%100), "RUB")
	return formatter.Format(amount)
}

func setIndicator(value int, limit int) string {
	if value/100 < 0 {
		return "🔴"
	}
	if value/100 < limit {
		return "\U0001F7E0"
	} else {
		return "\U0001F7E2"
	}
}
