package selectel

import (
	"encoding/json"
	"github.com/bojanz/currency"
	"github.com/px2x/infrastructure-tg-notifier/config"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
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

func CheckBillingMessage(service *config.Services, isSchedulerCheck bool) (message string, doSendReport bool) {
	firstString := "Проверили биллинг!\n\n"
	resultString := ""
	doSendReport = false
	result := true

	if isSchedulerCheck == false || service.LastCheckBilling.Add(service.CheckIntervalSelectelBilling).Before(time.Now()) {
		service.LastCheckBilling = time.Now()
		primary, vpc, storage := checkBilling(service.SelectelAPIKey)
		indicatorVpc, statusVpc := setIndicator(vpc, 2000)
		indicatorStorage, statusStorage := setIndicator(storage, 500)
		indicatorPrimary, statusPrimary := setIndicator(primary, 1000)
		resultString = "" +
			indicatorVpc + " Облачная платформа: " + converCurrency(vpc) + "\n" +
			indicatorStorage + " Хранилище: " + converCurrency(storage) + "\n" +
			indicatorPrimary + " Основной баланс: " + converCurrency(primary)
		result = statusVpc && statusStorage && statusPrimary
	}

	if result == false && isSchedulerCheck == true {
		service.LastCheckBilling = time.Now().Add(service.MuteTimeAfterError)
		firstString = "Ни Хао!\n\nПредложение финансового характера - пополнить казну!\n\n"
		doSendReport = true
	}

	return firstString + resultString, doSendReport
}

func converCurrency(value int) string {
	locale := currency.NewLocale("ru")
	formatter := currency.NewFormatter(locale)
	formatter.MaxDigits = 2
	amount, _ := currency.NewAmount(strconv.Itoa(value/100)+"."+strconv.Itoa(value%100), "RUB")
	return formatter.Format(amount)
}

func setIndicator(value int, limit int) (string, bool) {
	if value/100 < 0 {
		return "🔴", false
	}
	if value/100 < limit {
		return "\U0001F7E0", false
	} else {
		return "\U0001F7E2", true
	}
}
