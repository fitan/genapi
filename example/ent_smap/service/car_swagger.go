package genrest

import (
	"time"

	"github.com/gin-gonic/gin"
)

type CarNode struct {
	ID int `json:"id,omitempty"`

	Model        string    `json:"model,omitempty"   `
	RegisteredAt time.Time `json:"registered_at,omitempty" time_format:"2006-01-02 15:04:05" time_utc:"1"  format:"date-time" `
}

type CarNodeNotID struct {
	Model        string    `json:"model,omitempty"   `
	RegisteredAt time.Time `json:"registered_at,omitempty" time_format:"2006-01-02 15:04:05" time_utc:"1"  format:"date-time" `
}

type CarQuery struct {
}

// @Summary create one car
// @Accept  json
// @Produce  json
// @Tags Car
// @Param body body CarNodeNotID true " "
// @Success 200 {object} RestReturn{data=CarNode}
// @Router /car [post]
func CarCreateOne(c *gin.Context) {
}

// @Summary create list car
// @Accept  json
// @Produce  json
// @Tags Car
// @Param body body []CarNode true " "
// @Success 200 {object} RestReturn{data=[]CarNodeNotID}
// @Router /cars [post]
func CarCreateList(c *gin.Context) {
}

// @Summary delete one car
// @Accept  json
// @Produce  json
// @Tags Car
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=CarNode}
// @Router /car/{id} [delete]
func CarDeleteOne(c *gin.Context) {
}

// @Summary delete list car
// @Accept  json
// @Produce  json
// @Tags Car
// @Param ids query IdsQuery true " "
// @Success 200 {object} RestReturn{data=CarNode}
// @Router /cars [delete]
func CarDeleteList(c *gin.Context) {
}

// @Summary get one car
// @Accept  json
// @Produce  json
// @Tags Car
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=CarNode}
// @Router /car/{id} [get]
func CarGetOne(c *gin.Context) {
}

// @Summary get list car
// @Accept  json
// @Produce  json
// @Tags Car
// @Param data query CarQuery true " "
// @Header 200 {string} Count "The total amount"
// @Success 200 {object} RestReturn{data=GetCarListData}
// @Router /cars [get]
func CarGetList(c *gin.Context) {
}

// @Summary update one car
// @Accept  json
// @Produce  json
// @Tags Car
// @Param id path int true " "
// @Param body body CarNodeNotID true " "
// @Success 200 {object} RestReturn{data=CarNode}
// @Router /car/{id} [put]
func CarUpdateOne(c *gin.Context) {
}

// @Summary update list car
// @Accept  json
// @Produce  json
// @Tags Car
// @Param body body []CarNode true " "
// @Success 200 {object} RestReturn{data=CarNode}
// @Router /cars [put]
func CarUpdateList(c *gin.Context) {
}
