package Models

import (
	"errors"
	"fmt"
	"shop/Config"
	"shop/Redis"
	"strconv"
	"time"
)

func coolDownPeriod(lastTime string) bool {
	lastOrder, err := strconv.ParseInt(lastTime, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	presentTime := strconv.FormatInt(time.Now().UTC().Unix(), 10)
	presentOrder, err := strconv.ParseInt(presentTime, 10, 64)

	if err != nil {
		fmt.Println(err)
	}
	if presentOrder-lastOrder < 300 {
		return true
	}
	return false
}

func PlaceOrder(order *Order, id uint) (err error) {
	var customer Customer
	var product Product

	customerId := order.CustomerId
	productId := order.ProductId
	quantity := order.Quantity
	LockPeriod := 30000
	value := "product_id"
	var key = strconv.FormatUint(uint64(productId), 10)
	if isLocked, _ := Redis.Lock(key, value, LockPeriod); !isLocked {
		err := errors.New("ErrorMessage : Lock is Already Acquired")
		return err
	}
	time.Sleep(10 * time.Second)
	defer Redis.Unlock(key, value)

	if err = Config.DB.Where("customer_id = ?", customerId).First(&customer).Error; err != nil {
		err := errors.New("ErrorMessage : Please First Register")
		return err
	}

	if err = Config.DB.Where("product_id = ?", productId).First(&product).Error; err != nil {
		err := errors.New("ErrorMessage : Product Id Not Found")
		return err
	}

	if product.Quantity < quantity {
		err := errors.New("ErrorMessage: Not Enough Quantity")
		return err
	}
	lastTime := customer.LastOrder

	if coolDownPeriod(lastTime) {
		err := errors.New("ErrorMessage: CoolDown Period Going On")
		return err
	}
	order.Status = "Order Placed"
	presentTime := strconv.FormatInt(time.Now().UTC().Unix(), 10)

	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}

	customer.LastOrder = presentTime
	Config.DB.Save(customer)
	product.Quantity -= quantity
	Config.DB.Save(product)

	return nil
}

func GetAllCustomers(customer *[]Customer) (err error) {
	if err = Config.DB.Find(customer).Error; err != nil {
		return err
	}
	return nil
}

func RegisterCustomer(customer *Customer) (err error) {
	customer.LastOrder = "0"
	if err = Config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func GetCustomerOrders(order *[]Order, id string) (err error) {
	if err = Config.DB.Where("customer_id = ?", id).Find(order).Error; err != nil {
		return err
	}
	return nil
}
