package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Maged-Zaki/gin-rest-api/models"
	"github.com/Maged-Zaki/gin-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func GetAllEvents(c *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": events,
	})
}
func GetEvent(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(400, utils.FormatResponse(idStr+" is not a valid id", nil))
		return
	}

	event, err := models.GetEvent(id)

	if err != nil {
		c.JSON(404, utils.FormatResponse(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": event,
	})
}

func CreateEvent(c *gin.Context) {
	// parsedToken := c.MustGet("parsedToken")
	userId := c.GetInt64("userId")

	var event models.Event
	// Bind the JSON data to the event struct
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, utils.FormatResponse(err.Error(), nil))
		return
	}

	event.UserID = userId

	if err := event.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error saving: " + err.Error()})
		return
	}

	response := utils.FormatResponse("Created Successfully", event)
	c.JSON(http.StatusCreated, response)
}

func DeleteEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventIdStr := c.Param("id")

	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("%s is not a valid event id", eventIdStr),
		})
		return
	}

	// Validate the event belongs to the user
	event, err := models.GetEvent(eventId)
	if err != nil {
		c.JSON(404, utils.FormatResponse("Event not found", nil))
		return
	}
	if event.UserID != userId {
		c.JSON(403, utils.FormatResponse("You are not authorized to delete this event", nil))
		return
	}

	// Delete the event
	err = event.Delete()
	if err != nil {
		c.JSON(400, utils.FormatResponse("Error Deleting: "+err.Error(), nil))
		return
	}

	response := utils.FormatResponse("Deleted Successfully", nil)

	c.JSON(200, response)
}
func UpdateEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventIdStr := c.Param("id")

	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("%s is not a valid id", eventIdStr),
		})
		return
	}

	event, err := models.GetEvent(eventId)
	if err != nil {
		c.JSON(404, utils.FormatResponse("Event not found", nil))
		return
	}
	// Validate the event belongs to the user
	if event.UserID != userId {
		c.JSON(403, utils.FormatResponse("You are not authorized to update this event", nil))
		return
	}

	var updatedEvent models.Event

	// Bind the JSON data to the event struct
	err = c.ShouldBindJSON(&updatedEvent) // Careful with binding as the json input might include someone else's event id and then it's binded and updated
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error Updating: " + err.Error(),
		})
		return
	}

	response := utils.FormatResponse("Updated Successfully", updatedEvent)

	c.JSON(http.StatusCreated, response)
}
