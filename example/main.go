package main

import (
	"fmt"
	"github.com/vildan-valeev/yandexdelivery"
	yaModels "github.com/vildan-valeev/yandexdelivery/models"
	"os"
)

func main() {
	deliveryMethods()
}

func deliveryMethods() {
	cl := yandexdelivery.NewYandexClient("https://b2b.taxi.tst.yandex.net", true)
	methods, err := cl.DeliveryMethods(os.Getenv("YandexToken"),
		yaModels.DeliveryMethodsRequest{
			FullName: "улица Льва Толстого, 16, Москва",
			StartPoint: []float64{
				37.588074, 55.733924,
			},
		})
	if err != nil {
		return
	}

	fmt.Println(methods.ExpressDelivery)
}
