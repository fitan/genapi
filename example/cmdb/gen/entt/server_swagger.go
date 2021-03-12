package entt

import (
	"cmdb/ent/server"
	"time"

	"github.com/gin-gonic/gin"
)

type ServerNode struct {
	ID int `json:"id,omitempty"`

	CreateTime   time.Time           `json:"create_time,omitempty"  format:"date-time" `
	UpdateTime   time.Time           `json:"update_time,omitempty"  format:"date-time" `
	IP           string              `json:"ip,omitempty"   `
	MachineType  server.MachineType  `json:"machine_type,omitempty"   enums:"physical,virtual" binding:"oneof=physical,virtual"`
	PlatformType server.PlatformType `json:"platform_type,omitempty"   enums:"zstack,k8s,openstack" binding:"oneof=zstack,k8s,openstack"`
	SystemType   server.SystemType   `json:"system_type,omitempty"   enums:"linux,windows" binding:"oneof=linux,windows"`
}

type ServerNodeNotID struct {
	CreateTime   time.Time           `json:"create_time,omitempty"  format:"date-time" `
	UpdateTime   time.Time           `json:"update_time,omitempty"  format:"date-time" `
	IP           string              `json:"ip,omitempty"   `
	MachineType  server.MachineType  `json:"machine_type,omitempty"   enums:"physical,virtual" binding:"oneof=physical,virtual"`
	PlatformType server.PlatformType `json:"platform_type,omitempty"   enums:"zstack,k8s,openstack" binding:"oneof=zstack,k8s,openstack"`
	SystemType   server.SystemType   `json:"system_type,omitempty"   enums:"linux,windows" binding:"oneof=linux,windows"`
}

type ServerQuery struct {
}

// @Summary create one server
// @Accept  json
// @Produce  json
// @Tags Server
// @Param body body ServerNodeNotID true " "
// @Success 200 {object} RestReturn{data=ServerNode}
// @Router /server [post]
func ServerCreateOne(c *gin.Context) {
}

// @Summary create list server
// @Accept  json
// @Produce  json
// @Tags Server
// @Param body body []ServerNode true " "
// @Success 200 {object} RestReturn{data=[]ServerNodeNotID}
// @Router /servers [post]
func ServerCreateList(c *gin.Context) {
}

// @Summary delete one server
// @Accept  json
// @Produce  json
// @Tags Server
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=ServerNode}
// @Router /server/{id} [delete]
func ServerDeleteOne(c *gin.Context) {
}

// @Summary delete list server
// @Accept  json
// @Produce  json
// @Tags Server
// @Param ids query IdsQuery true " "
// @Success 200 {object} RestReturn{data=ServerNode}
// @Router /servers [delete]
func ServerDeleteList(c *gin.Context) {
}

// @Summary get one server
// @Accept  json
// @Produce  json
// @Tags Server
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=ServerNode}
// @Router /server/{id} [get]
func ServerGetOne(c *gin.Context) {
}

// @Summary get list server
// @Accept  json
// @Produce  json
// @Tags Server
// @Param data query ServerQuery true " "
// @Header 200 {string} Count "The total amount"
// @Success 200 {object} RestReturn{data=GetServerListData}
// @Router /servers [get]
func ServerGetList(c *gin.Context) {
}

// @Summary update one server
// @Accept  json
// @Produce  json
// @Tags Server
// @Param id path int true " "
// @Param body body ServerNodeNotID true " "
// @Success 200 {object} RestReturn{data=ServerNode}
// @Router /server/{id} [put]
func ServerUpdateOne(c *gin.Context) {
}

// @Summary update list server
// @Accept  json
// @Produce  json
// @Tags Server
// @Param body body []ServerNode true " "
// @Success 200 {object} RestReturn{data=ServerNode}
// @Router /servers [put]
func ServerUpdateList(c *gin.Context) {
}

// @Summary create list service
// @Accept  json
// @Produce  json
// @Tags Service by Server
// @Param id path int true " "
// @Param body body  []ServiceNodeNotID true " "
// @Success 200 {object} RestReturn{data=[]ServiceNodeNotID}
// @Router /server/{id}/services [post]
func CreateListServicesByServer(c *gin.Context) {
}

// @Summary delete one service
// @Accept  json
// @Produce  json
// @Tags Service by Server
// @Param id path int true " "
// @Param query query  ServiceQuery false " "
// @Success 200 {object} RestReturn{data=ServiceNode}
// @Router /server/{id}/services [delete]
func DeleteListServicesByServer(c *gin.Context) {
}

// @Summary get list service
// @Accept  json
// @Produce  json
// @Tags Service by Server
// @Param id path int true " "
// @Param query query  ServiceQuery false " "
// @Success 200 {object} RestReturn{data=[]ServiceNode}
// @Router /server/{id}/services [get]
func GetListServicesByServer(c *gin.Context) {
}
