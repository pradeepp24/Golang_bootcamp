package Models

import "shop/Config"

func GetAllProducts(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *Product, id string) (err error) {

	if err = Config.DB.Where("product_id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}
