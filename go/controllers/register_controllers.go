package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gonggantt/go/orm"
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

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	v1 := r.Group("/api/github.com/fullstack-lang/gonggantt/go")
	{ // insertion point for registrations
		v1.GET("/v1/arrows", GetArrows)
		v1.GET("/v1/arrows/:id", GetArrow)
		v1.POST("/v1/arrows", PostArrow)
		v1.PATCH("/v1/arrows/:id", UpdateArrow)
		v1.PUT("/v1/arrows/:id", UpdateArrow)
		v1.DELETE("/v1/arrows/:id", DeleteArrow)

		v1.GET("/v1/bars", GetBars)
		v1.GET("/v1/bars/:id", GetBar)
		v1.POST("/v1/bars", PostBar)
		v1.PATCH("/v1/bars/:id", UpdateBar)
		v1.PUT("/v1/bars/:id", UpdateBar)
		v1.DELETE("/v1/bars/:id", DeleteBar)

		v1.GET("/v1/gantts", GetGantts)
		v1.GET("/v1/gantts/:id", GetGantt)
		v1.POST("/v1/gantts", PostGantt)
		v1.PATCH("/v1/gantts/:id", UpdateGantt)
		v1.PUT("/v1/gantts/:id", UpdateGantt)
		v1.DELETE("/v1/gantts/:id", DeleteGantt)

		v1.GET("/v1/groups", GetGroups)
		v1.GET("/v1/groups/:id", GetGroup)
		v1.POST("/v1/groups", PostGroup)
		v1.PATCH("/v1/groups/:id", UpdateGroup)
		v1.PUT("/v1/groups/:id", UpdateGroup)
		v1.DELETE("/v1/groups/:id", DeleteGroup)

		v1.GET("/v1/lanes", GetLanes)
		v1.GET("/v1/lanes/:id", GetLane)
		v1.POST("/v1/lanes", PostLane)
		v1.PATCH("/v1/lanes/:id", UpdateLane)
		v1.PUT("/v1/lanes/:id", UpdateLane)
		v1.DELETE("/v1/lanes/:id", DeleteLane)

		v1.GET("/v1/laneuses", GetLaneUses)
		v1.GET("/v1/laneuses/:id", GetLaneUse)
		v1.POST("/v1/laneuses", PostLaneUse)
		v1.PATCH("/v1/laneuses/:id", UpdateLaneUse)
		v1.PUT("/v1/laneuses/:id", UpdateLaneUse)
		v1.DELETE("/v1/laneuses/:id", DeleteLaneUse)

		v1.GET("/v1/milestones", GetMilestones)
		v1.GET("/v1/milestones/:id", GetMilestone)
		v1.POST("/v1/milestones", PostMilestone)
		v1.PATCH("/v1/milestones/:id", UpdateMilestone)
		v1.PUT("/v1/milestones/:id", UpdateMilestone)
		v1.DELETE("/v1/milestones/:id", DeleteMilestone)

		v1.GET("/commitfrombacknb", GetLastCommitFromBackNb)
		v1.GET("/pushfromfrontnb", GetLastPushFromFrontNb)
	}
}

// swagger:route GET /commitfrombacknb backrepo GetLastCommitFromBackNb
func GetLastCommitFromBackNb(c *gin.Context) {
	res := orm.GetLastCommitFromBackNb()

	c.JSON(http.StatusOK, res)
}

// swagger:route GET /pushfromfrontnb backrepo GetLastPushFromFrontNb
func GetLastPushFromFrontNb(c *gin.Context) {
	res := orm.GetLastPushFromFrontNb()

	c.JSON(http.StatusOK, res)
}
