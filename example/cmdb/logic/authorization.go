package logic

import (
	"cmdb/models"
	"github.com/gin-gonic/gin"
)

type AuthorizationBody struct {
	Role string `json:"role"`
	Path string `json:"path"`
	Method string `json:"method"`
}

type AuthorizationIn struct {
	Body AuthorizationBody
}

type AuthorizationOut struct {
	Success bool `json:"msg"`
}

type GetPolicyIn struct {
	Uri GetPolicyQuery
}

type GetPolicyQuery struct {
	Id int `uri:"id"`
}

type UpdatePolicyIn struct {
	Body UpdatePolicyBody
}

type UpdatePolicyBody struct {
	Old AuthorizationBody
	New AuthorizationBody
}

// @GenApi /api/add_policy [post]
func AddPolicy(c *gin.Context, in *AuthorizationIn) (*AuthorizationOut, error) {
	has, err := models.GetCasbin().AddPolicy(in.Body.Role, in.Body.Path, in.Body.Method)
	out := AuthorizationOut{has}
	return &out, err
}

// @GenApi /api/get_policy [get]
func GetPolicyList(c *gin.Context, in *AuthorizationIn) ([][]string,error) {
	return models.GetCasbin().GetPolicy(), nil
}

// @GenApi /api/get_policy/{id} [get]
func GetPolicyOne(c *gin.Context, in *GetPolicyIn) ([][]string, error) {
	return models.GetCasbin().GetFilteredPolicy(in.Uri.Id), nil
}

// @GenApi /api/delete_policy/:id [delete]
func DeletePolicy(c *gin.Context, in *AuthorizationIn) (bool, error) {
	return models.GetCasbin().RemovePolicy(in.Body.Role, in.Body.Path, in.Body.Method)
}

// @GenApi /api/update_policy [put]
func UpdatePolicy(c *gin.Context, in *UpdatePolicyIn) (bool, error)  {
	oldL := []string{in.Body.Old.Role, in.Body.Old.Path, in.Body.Old.Method}
	newL := []string{in.Body.New.Role, in.Body.New.Path, in.Body.New.Method}
	return models.GetCasbin().UpdatePolicy(oldL, newL)
}

type SayHelloIn struct {
}



// @GenApi /api/say_hello [get]
func SayHello(c *gin.Context, in *SayHelloIn) (string, error) {
	return "hello", nil
}


