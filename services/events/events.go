package events

import (
	"net/http"
	"strconv"

	"github.com/Sahil0722/event-booking-app/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)
}

func GetEvent(context *gin.Context) {
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

	context.JSON(http.StatusOK, event)
}

func CreateEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	event.UserID = context.GetInt64("userID")

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func UpdateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID := context.GetInt64("userID")
	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to Update this event!"})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updatedEvent.ID = eventID

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Updated Event Successfully!"})
}

func DeleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userID := context.GetInt64("userID")
	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to Delete this event!"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted Event Successfully!"})
}
