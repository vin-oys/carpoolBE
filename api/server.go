package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vin-oys/api-carpool/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.Use(CORS)

	userRoutes := router.Group("/user")

	userRoutes.POST("/create", server.createUser)
	userRoutes.GET("/get", server.getUser)
	userRoutes.PUT("/update", server.updateUser)
	userRoutes.DELETE("/delete", server.deleteUser)

	carRoutes := router.Group("/car")

	carRoutes.POST("/create", server.createCar)
	carRoutes.GET("/get", server.getCar)
	carRoutes.PUT("/update", server.updateCar)
	carRoutes.DELETE("/delete", server.deleteCar)

	passengerRoute := router.Group("/passenger")

	passengerRoute.POST("/", server.createSchedulePassenger)
	passengerRoute.GET("/", server.getScheduledPassenger)
	passengerRoute.GET("/list", server.listScheduledPassengers)
	passengerRoute.PUT("/schedule", server.updatePassengerSchedule)
	passengerRoute.PUT("/seat", server.updatePassengerSeat)
	passengerRoute.DELETE("/", server.deleteSchedulePassenger)

	scheduleRoutes := router.Group("/schedule")
	scheduleRoutes.POST("/create", server.createSchedule)
	scheduleRoutes.GET("/get", server.getSchedule)
	scheduleRoutes.GET("/list", server.listSchedule)
	scheduleRoutes.PUT("/update/departureDate", server.updateScheduleDepartureDate)
	scheduleRoutes.PUT("/update/departureTime", server.updateScheduleDepartureTime)
	scheduleRoutes.PUT("/update/driverId", server.updateScheduleDriverId)
	scheduleRoutes.PUT("/update/dropOff", server.updateScheduleDropOff)
	scheduleRoutes.PUT("/update/pickup", server.updateSchedulePickup)
	scheduleRoutes.PUT("/update/plateId", server.updateSchedulePlateId)
	scheduleRoutes.DELETE("/delete", server.deleteSchedule)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	//handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
