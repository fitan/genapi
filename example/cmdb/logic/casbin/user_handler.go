package casbin

import (
	"cmdb/ent"
	"cmdb/ent/user"
	"cmdb/gen/entt"
	"context"
	"github.com/gin-gonic/gin"
)
var userCURD *UserCURDV2

//type UserCURDV2 struct {
//	Db *ent.Client
//
//	RoleBindingObj *RoleBindingCURD
//
//	AlertObj *AlertCURD
//}

func NewUserCURDV2(db *ent.Client) {
	userCURD = &UserCURDV2{
		Db: db,

		RoleBindingObj: &entt.RoleBindingCURD{
			Db: db,
		},

		AlertObj: &entt.AlertCURD{
			Db: db,
		},
	}
}

func GetUserCURDV2() *UserCURDV2 {
	return userCURD
}

type GetUserIn struct {
	Uri entt.IdUri
}

func GetUser(c *gin.Context, in *GetUserIn) (*ent.User, error) {
	curd := GetUserCURDV2()
	queryer := curd.Db.User.Query().Where(user.IDEQ(in.Uri.ID))
	entt.UserSelete(queryer)
	return queryer.Only(context.Background())
}

type GetUsersIn struct {
	Query entt.UserDefaultQuery
}

func GetUsers(c *gin.Context, in *GetUsersIn) (*entt.GetUserListData, error) {
	userCURD := GetUserCURDV2()
	countQueryer := userCURD.Db.User.Query()
	listQueryer := userCURD.Db.User.Query()
	bg := context.Background()
	err := userCURD.DefaultGetListCount(countQueryer, &in.Query)
	if err != nil {
		return nil, err
	}
	count, err := countQueryer.Count(bg)
	if err != nil {
		return nil, err
	}
	err = userCURD.DefaultGetListQueryer(listQueryer, &in.Query)
	if err != nil {
		return nil,err
	}

	list, err := listQueryer.All(bg)
	if err != nil {
		return nil, err
	}

	return &entt.GetUserListData{count,list}, nil
}

type CreateUserIn struct {
	Body ent.User
}

func CreateUser(c *gin.Context, in *CreateUserIn) (*ent.User, error) {
	creater := GetUserCURDV2().Db.User.Create()
	entt.UserCreateMutation(creater.Mutation(), &in.Body)
	return creater.Save(context.Background())
}

type CreateUsersIn struct {
	Body []ent.User
}

func CreateUsers(c *gin.Context, in *CreateUsersIn)  {
	bulk := make([]*ent.UserCreate, 0, len(in.Body))
	for _, v := range body {
		creater :=
	}
}
