package assets

import (
	"concurrency-experimentals/models"
	"math/rand"
	"time"
)

func GenerateOrder(product *models.Product, user *models.User, random *rand.Rand) *models.Order {
	order := new(models.Order)
	order.GenerateID()
	order.ProductID = product.ID
	order.UserID = user.ID
	diff := time.Now().Sub(user.CreatedAt)
	createdAt := user.CreatedAt.Add(time.Duration(random.Intn(int(diff))))
	order.CreatedAt = &createdAt
	order.UnitPrice = product.Price
	order.Quantity = random.Intn(1000)
	order.Amount = order.UnitPrice * float64(order.Quantity)
	order.GenerateUpdatedAt(random)
	order.GenerateDescription()
	return order
}
