package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/enchant97/frameworks_and_languages_module/server/core"
	cors "github.com/enchant97/go-gin-cors"
	"github.com/gin-gonic/gin"
)

// Setup & register the Routes
func InitRoutes(engine *gin.Engine) {
	engine.Use(cors.Default())
	rootRoutes := engine.Group("/")
	{
		rootRoutes.GET("/", GetIndex)
	}
	itemRoutes := engine.Group("/item")
	{
		itemRoutes.POST("/", PostNewItem)
		itemRoutes.GET("/:itemID/", GetItemByID)
		itemRoutes.DELETE("/:itemID/", DeleteItemByID)
	}
	engine.GET("/items/", GetItems)
}

// Route to get the human readable index page,
// saying that the server is operational
func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// Route for adding a new item
func PostNewItem(c *gin.Context) {
	var newItem core.ItemCreate
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.AbortWithStatus(http.StatusMethodNotAllowed)
		return
	}
	createdItem := core.CreateNewItem(newItem)
	c.JSONP(http.StatusCreated, createdItem)
}

// Route for getting an existing item by it's id
func GetItemByID(c *gin.Context) {
	rawItemID := c.Param("itemID")
	itemID, err := strconv.ParseInt(rawItemID, 10, 64)
	if err == nil {
		if item := core.GetItemByID(itemID); item != nil {
			c.JSONP(http.StatusOK, item)
			return
		}
	}
	c.AbortWithStatus(http.StatusNotFound)
}

// Route to delete an existing item by it's id
func DeleteItemByID(c *gin.Context) {
	rawItemID := c.Param("itemID")
	itemID, err := strconv.ParseInt(rawItemID, 10, 64)
	if err == nil {
		if exists := core.DeleteItemByID(itemID); exists {
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.AbortWithStatus(http.StatusNotFound)
}

// Route to get all items (with query filters)
func GetItems(c *gin.Context) {
	var filters core.ItemsFilterFromRequest
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.AbortWithError(http.StatusMethodNotAllowed, err)
		return
	}
	var dateFrom *core.PythonISOTime
	var dateTo *core.PythonISOTime

	if filters.DateFrom != nil {
		date, err := time.Parse("2006-01-02T15:04:05", *filters.DateFrom)
		date2 := core.PythonISOTime(date)
		if err != nil {
			c.AbortWithError(http.StatusMethodNotAllowed, err)
			return
		}
		dateFrom = &date2
	}
	if filters.DateTo != nil {
		date, err := time.Parse("2006-01-02T15:04:05", *filters.DateTo)
		date2 := core.PythonISOTime(date)
		if err != nil {
			c.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}
		dateTo = &date2
	}

	items := core.GetItemsFiltered(core.ItemsFilter{
		UserID:      filters.UserID,
		CSVKeywords: filters.CSVKeywords,
		Lat:         filters.Lat,
		Lon:         filters.Lon,
		Radius:      filters.Radius,
		DateFrom:    dateFrom,
		DateTo:      dateTo,
	})
	c.JSONP(http.StatusOK, items)
}
