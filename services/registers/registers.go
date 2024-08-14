package registers

import (
	"net/http"
	"strconv"

	"github.com/Sahil0722/event-booking-app/models"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = event.Register(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Register For Event Successfully!"})
}

func CancelRegistrationEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var event models.Event
	event.ID = eventID

	err = event.CancelRegistration(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Un-Registered for Event Successfully"})
}

func GetRegistrations(context *gin.Context) {
	registers, err := models.GetAllRegistrations()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, registers)
}
