package repository

import (
	"time"

	"ordersystem/model"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	// Init the drinks slice with some test data
	drinks := []model.Drink{
		{ID: 1, Name: "Beer", Price: 2.5, Description: "A refreshing beer"},
		{ID: 2, Name: "Spritzer", Price: 1.4, Description: "Wine mixed with soda water"},
		{ID: 3, Name: "Coffe", Price: 0, Description: "Classic black coffee"},
	}
	// Init orders slice with some test data
	orders := []model.Order{} // start empty

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
// calculate total orders
// key = DrinkID, value = Amount of orders
// totalledOrders map[uint64]uint64
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	totalledOrders := make(map[uint64]uint64)
	for _, o := range db.orders {
		totalledOrders[o.DrinkID] += uint64(o.Amount)
	}
	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	// todo
	// add order to db.orders slice
	if order.CreatedAt.IsZero() {
		order.CreatedAt = time.Now()
	}
	db.orders = append(db.orders, *order)
}
