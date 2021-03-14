package entt

import "github.com/gin-gonic/gin"

type AlertID struct {
	ID int `json:"id,omitempty"`
}

type AlertNode struct {
	AlertID
	AlertNodeNotID
}

type AlertEdges struct {
}

type AlertNodeNotID struct {
	Name  string     `json:"name,omitempty"   `
	Edges AlertEdges `json:"edges"`
}

type AlertQuery struct {
}

// @Summary create one alert
// @Accept  json
// @Produce  json
// @Tags Alert
// @Param body body AlertNodeNotID true " "
// @Success 200 {object} RestReturn{data=AlertNode}
// @Router /alert [post]
func AlertCreateOne(c *gin.Context) {
}

// @Summary create list alert
// @Accept  json
// @Produce  json
// @Tags Alert
// @Param body body []AlertNode true " "
// @Success 200 {object} RestReturn{data=[]AlertNodeNotID}
// @Router /alerts [post]
func AlertCreateList(c *gin.Context) {
}

// @Summary delete one alert
// @Accept  json
// @Produce  json
// @Tags Alert
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=AlertNode}
// @Router /alert/{id} [delete]
func AlertDeleteOne(c *gin.Context) {
}

// @Summary delete list alert
// @Accept  json
// @Produce  json
// @Tags Alert
// @Param ids query IdsQuery true " "
// @Success 200 {object} RestReturn{data=AlertNode}
// @Router /alerts [delete]
func AlertDeleteList(c *gin.Context) {
}

// @Summary get one alert
// @Accept  json
// @Produce  json
// @Tags Alert
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=AlertNode}
// @Router /alert/{id} [get]
func AlertGetOne(c *gin.Context) {
}

// @Summary get list alert
// @Accept  json
// @Produce  json
// @Tags Alert
// @Param data query AlertQuery true " "
// @Header 200 {string} Count "The total amount"
// @Success 200 {object} RestReturn{data=GetAlertListData}
// @Router /alerts [get]
func AlertGetList(c *gin.Context) {
}

// @Summary update one alert
// @Accept  json
// @Produce  json
// @Tags Alert
// @Param id path int true " "
// @Param body body AlertNodeNotID true " "
// @Success 200 {object} RestReturn{data=AlertNode}
// @Router /alert/{id} [put]
func AlertUpdateOne(c *gin.Context) {
}

// @Summary update list alert
// @Accept  json
// @Produce  json
// @Tags Alert
// @Param body body []AlertNode true " "
// @Success 200 {object} RestReturn{data=AlertNode}
// @Router /alerts [put]
func AlertUpdateList(c *gin.Context) {
}
