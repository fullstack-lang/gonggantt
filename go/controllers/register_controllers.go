// generated code - do not edit
package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// genQuery return the name of the column
func genQuery(columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Field   string `json:"field"`
	} `json:"body"`
}

// registerControllers register controllers
func registerControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gonggantt/go")
	{ // insertion point for registrations
		v1.GET("/v1/arrows", GetController().GetArrows)
		v1.GET("/v1/arrows/:id", GetController().GetArrow)
		v1.POST("/v1/arrows", GetController().PostArrow)
		v1.PATCH("/v1/arrows/:id", GetController().UpdateArrow)
		v1.PUT("/v1/arrows/:id", GetController().UpdateArrow)
		v1.DELETE("/v1/arrows/:id", GetController().DeleteArrow)

		v1.GET("/v1/bars", GetController().GetBars)
		v1.GET("/v1/bars/:id", GetController().GetBar)
		v1.POST("/v1/bars", GetController().PostBar)
		v1.PATCH("/v1/bars/:id", GetController().UpdateBar)
		v1.PUT("/v1/bars/:id", GetController().UpdateBar)
		v1.DELETE("/v1/bars/:id", GetController().DeleteBar)

		v1.GET("/v1/gantts", GetController().GetGantts)
		v1.GET("/v1/gantts/:id", GetController().GetGantt)
		v1.POST("/v1/gantts", GetController().PostGantt)
		v1.PATCH("/v1/gantts/:id", GetController().UpdateGantt)
		v1.PUT("/v1/gantts/:id", GetController().UpdateGantt)
		v1.DELETE("/v1/gantts/:id", GetController().DeleteGantt)

		v1.GET("/v1/groups", GetController().GetGroups)
		v1.GET("/v1/groups/:id", GetController().GetGroup)
		v1.POST("/v1/groups", GetController().PostGroup)
		v1.PATCH("/v1/groups/:id", GetController().UpdateGroup)
		v1.PUT("/v1/groups/:id", GetController().UpdateGroup)
		v1.DELETE("/v1/groups/:id", GetController().DeleteGroup)

		v1.GET("/v1/lanes", GetController().GetLanes)
		v1.GET("/v1/lanes/:id", GetController().GetLane)
		v1.POST("/v1/lanes", GetController().PostLane)
		v1.PATCH("/v1/lanes/:id", GetController().UpdateLane)
		v1.PUT("/v1/lanes/:id", GetController().UpdateLane)
		v1.DELETE("/v1/lanes/:id", GetController().DeleteLane)

		v1.GET("/v1/laneuses", GetController().GetLaneUses)
		v1.GET("/v1/laneuses/:id", GetController().GetLaneUse)
		v1.POST("/v1/laneuses", GetController().PostLaneUse)
		v1.PATCH("/v1/laneuses/:id", GetController().UpdateLaneUse)
		v1.PUT("/v1/laneuses/:id", GetController().UpdateLaneUse)
		v1.DELETE("/v1/laneuses/:id", GetController().DeleteLaneUse)

		v1.GET("/v1/milestones", GetController().GetMilestones)
		v1.GET("/v1/milestones/:id", GetController().GetMilestone)
		v1.POST("/v1/milestones", GetController().PostMilestone)
		v1.PATCH("/v1/milestones/:id", GetController().UpdateMilestone)
		v1.PUT("/v1/milestones/:id", GetController().UpdateMilestone)
		v1.DELETE("/v1/milestones/:id", GetController().DeleteMilestone)

		v1.GET("/v1/commitfrombacknb", GetController().GetLastCommitFromBackNb)
		v1.GET("/v1/pushfromfrontnb", GetController().GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func (controller *Controller) GetLastCommitFromBackNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastCommitFromBackNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gonggantt/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func (controller *Controller) GetLastPushFromFrontNb(c *gin.Context) {
	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetLastPushFromFrontNb", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gonggantt/go/models, Unkown stack", stackPath)
	}
	res := backRepo.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
