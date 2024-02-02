package main

import (
	"fmt"
	"github.com/vildan-valeev/yandexdelivery"
	yaModels "github.com/vildan-valeev/yandexdelivery/models"
	"os"
)

func main() {
	//deliveryMethods()
	//checkPrice()
	tariffs()
	offersCalc()
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

func checkPrice() {
	cl := yandexdelivery.NewYandexClient("https://b2b.taxi.tst.yandex.net", true)
	price, err := cl.CheckPrice(os.Getenv("YandexToken"),
		yaModels.CheckPriceRequest{
			Items: []yaModels.DeliveryItem{
				{
					Quantity: 1,
					Size: &yaModels.Size{
						Height: 0.1,
						Length: 0.1,
						Width:  0.1,
					},
					Weight: 2,
				},
			},
			Requirements: yaModels.CheckPriceRequirement{
				CargoLoaders: 0,
				CargoOptions: []string{"thermobag", "auto_courier"},
				CargoType:    "lcv_m",
				ProCourier:   false,
				TaxiClass:    "courier",
				SameDayData: yaModels.SameDayData{
					DeliveryInterval: yaModels.DeliveryInterval{
						From: "2024-01-26T00:00:00+00:00",
						To:   "2024-01-27T00:00:00+00:00",
					},
				},
			},
			RoutePoints: []yaModels.RoutePoint{
				{Coordinates: []float64{37.588074, 55.733924}},
				{Coordinates: []float64{37.588074, 55.733924}},
			},
			SkipDoorToDoor: false,
		})
	if err != nil {
		return
	}

	fmt.Println(price.Price)
}

func tariffs() {
	cl := yandexdelivery.NewYandexClient("https://b2b.taxi.tst.yandex.net", true)
	trfs, err := cl.Tariffs(os.Getenv("YandexToken"),
		yaModels.TariffsRequest{
			FullName: "улица Льва Толстого, 16, Москва",
			StartPoint: []float64{
				37.588074, 55.733924,
			},
		})
	if err != nil {
		return
	}

	fmt.Println(trfs.AvailableTariffs)
}

func offersCalc() {
	cl := yandexdelivery.NewYandexClient("https://b2b.taxi.tst.yandex.net", true)
	of, err := cl.OffersCalculate(os.Getenv("YandexToken"),
		yaModels.OffersCalculateRequest{
			Items: []yaModels.DeliveryItem{
				{
					Quantity: 1,
					Size: &yaModels.Size{
						Height: 0.1,
						Length: 0.1,
						Width:  0.1,
					},
					Weight: 2,
				},
			},
		})
	if err != nil {
		return
	}

	fmt.Println(of)
}
