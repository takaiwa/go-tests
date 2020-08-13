package time

import (
	"time"
)

const limit = -time.Minute * 30

var now = time.Now

func CanNotify(deliveryTime time.Time) bool {

	deadline := deliveryTime.Add(limit)
	//fmt.Println("Delivery time:", deliveryTime)
	//fmt.Println("Deadline:", deadline)
	//fmt.Println("Now:", now())

	if now().Equal(deliveryTime) || now().After(deliveryTime) {
		return false
	}
	return now().Equal(deadline) || now().After(deadline)
}
