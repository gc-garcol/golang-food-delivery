package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Restaurant struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func main() {
	// CONNECT
	dsn := os.Getenv("MYSQL_CONNECTION")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(db, err)

	r := gin.Default()

	// GET /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	// POST restaurants
	v1.POST("/restaurants", func(c *gin.Context) {
		var data Restaurant
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		db.Create(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	v1.PATCH("/restaurants/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		db.Where("id = ?", id).Updates(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	v1.GET("/restaurants/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data Restaurant
		db.Where("id = ?", id).First(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	v1.GET("/restaurants", func(c *gin.Context) {
		type Paging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}
		var pagingData Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if pagingData.Page <= 0 {
			pagingData.Page = 1
		}

		if pagingData.Page <= 0 {
			pagingData.Page = 5
		}

		var data []Restaurant
		db.Offset((pagingData.Page - 1) * pagingData.Limit).
			Order("id desc").
			Limit(pagingData.Limit).
			Find(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	v1.DELETE("/restaurants/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil)
		c.JSON(http.StatusOK, gin.H{
			"data": 1,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") + Block here

	// CREATE
	//newRestaurant := Restaurant{Name: "mei yuan neef", Address: "chung nha voi Thai"}
	//if err := db.Create(&newRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//log.Println("new ID: ", newRestaurant.Id)

	// READ
	var readRestaurant Restaurant
	if err := db.Where("id = ?", 3).First(&readRestaurant).Error; err != nil {
		log.Println(err)
	}
	log.Println(readRestaurant)

	// UPDATE
	//nilName := ""
	var nilName string
	updateRestaurant := RestaurantUpdate{Name: &nilName}

	if err := db.Where("id = ?", 3).Updates(&updateRestaurant).Error; err != nil {
		log.Println(err)
	}

	// DELETE
	if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 1).Delete(nil).Error; err != nil {
		log.Println(err)
	}
}
