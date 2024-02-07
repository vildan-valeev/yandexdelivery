# Library fo API [Yandex Delivery](https://yandex.ru/dev/logistics/api/about/intro.html)

---
If `debugmode` is false, request and response json data wiil be printed to console.

---
### Example
```go
package main

import (
	"fmt"
	"os"
	
	"github.com/vildan-valeev/yandexdelivery"
	yaModels "github.com/vildan-valeev/yandexdelivery/models"

)

func main() {
	cl := yandexdelivery.NewYandexClient(os.Getenv("YandexURL"), os.Getenv("YandexToken"), true)
	methods, err := cl.DeliveryMethods(
		yaModels.DeliveryMethodsRequest{
			FullName: "улица Льва Толстого, 16, Москва",
			StartPoint: []float64{
				37.588074, 55.733924,
			},
		})
	if err != nil {
		return
	}

	fmt.Println(methods)
}


```
