package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/Models"
)

func GetProducts(c *gin.Context) {
	var product []Models.Product
	err := Models.GetAllProducts(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

//UpdateUser ... Update the user information
