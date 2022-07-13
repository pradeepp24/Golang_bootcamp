package Models

import (
	"fmt"
	"retail_shop/Config"
)

func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Where("Product_id = ?", id).Delete(product)
	return nil
}
