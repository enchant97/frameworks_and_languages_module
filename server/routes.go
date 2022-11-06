package main

/*
This contains the setup code for initialising
the routes and the actual route methods, which are called by gin
*/

import (
	"net/http"
	"strconv"

	"github.com/enchant97/frameworks_and_languages_module/server/core"
	cors "github.com/enchant97/go-gin-cors"
	"github.com/gin-gonic/gin"
)

// Setup & register the Routes
func InitRoutes(engine *gin.Engine) {
	// apply to cors to all routes, using
	// default configs which allows all origins by default
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
	// SOURCE: https://gin-gonic.com/docs/examples/html-rendering/
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// Route for adding a new item
func PostNewItem(c *gin.Context) {
	// SOURCE: https://gin-gonic.com/docs/examples/bind-query-or-post/
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
	// SOURCE: https://pkg.go.dev/strconv#ParseInt
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
	// SOURCE: https://gin-gonic.com/docs/examples/only-bind-query-string/
	var filters core.ItemsFilterFromRequest
	if err := c.ShouldBindQuery(&filters); err != nil {
		c.AbortWithError(http.StatusMethodNotAllowed, err)
		return
	}
	var dateFrom *core.PythonISOTime
	var dateTo *core.PythonISOTime

	// if a date-from filter has been set,
	// convert into valid date object
	if filters.DateFrom != nil {
		date, err := core.ParsePythonISOTime(*filters.DateFrom)
		if err != nil {
			c.AbortWithError(http.StatusMethodNotAllowed, err)
			return
		}
		dateFrom = &date
	}
	// if a date-to filter has been set,
	// convert into valid date object
	if filters.DateTo != nil {
		date, err := core.ParsePythonISOTime(*filters.DateTo)
		if err != nil {
			c.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}
		dateTo = &date
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
