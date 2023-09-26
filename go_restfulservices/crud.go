package main

import (
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Vehicle struct {
	Id    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

var vehicles = []Vehicle{
	{Id: 1, Brand: "Toyota", Model: "Vios"},
	{Id: 2, Brand: "Toyota", Model: "Altis"},
	{Id: 3, Brand: "Honda", Model: "City"},
	{Id: 4, Brand: "Honda", Model: "Civic"},
	{Id: 5, Brand: "Nissan", Model: "Almera"},
	{Id: 6, Brand: "Nissan", Model: "Sylphy"},
	{Id: 7, Brand: "Mazda", Model: "3"},
	{Id: 8, Brand: "Mazda", Model: "6"},
	{Id: 9, Brand: "Proton", Model: "Saga"},
	{Id: 10, Brand: "Proton", Model: "Persona"},
}

func main() {
	router := gin.Default()

	//GET for all vehicles
	router.GET("/vehicles", GetAllVehicles)

	//GET for vehicle by id
	router.GET("/vehicles/:vid", GetVehicleById)

	//POST for new vehicle
	router.POST("/vehicles", AddVehicle)

	//DELETE for vehicle by id
	router.DELETE("/vehicles/:vid", DeleteVehicle)

	//PUT for vehicle by id
	router.PUT("/vehicles/:vid", ModifyVehicle)

	var port string = ":9999"
	fmt.Println("Server listening on ", port)
	router.Run(port)
}

func GetAllVehicles(c *gin.Context) {
	c.JSON(http.StatusOK, vehicles)
}
func GetVehicleById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("vid"))
	for _, v := range vehicles {
		if v.Id == id {
			c.JSON(http.StatusOK, v)
			return
		}
	}
	c.JSON(http.StatusNotFound, nil)
}
func AddVehicle(c *gin.Context) {
	//read from request body
	var v Vehicle
	err := c.BindJSON(&v)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	//add to slice
	vehicles = append(vehicles, v)
	//return status created
	c.JSON(http.StatusCreated, v)
}

func DeleteVehicle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("vid"))
	for i, v := range vehicles {
		if v.Id == id {
			vehicles = append(vehicles[:i], vehicles[i+1:]...)
			c.JSON(http.StatusOK, v)
			return
		}
	}
	c.JSON(http.StatusNotFound, nil)
}

func ModifyVehicle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("vid"))
	var vh Vehicle
	err := c.BindJSON(&vh)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	for i, v := range vehicles {
		if v.Id == id {
			vehicles[i].Brand = vh.Brand
			vehicles[i].Model = vh.Model
			c.JSON(http.StatusOK, vh)
			return
		}
	}
	c.JSON(http.StatusNotFound, nil)
}
