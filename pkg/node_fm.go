package pkg

import "encoding/json"

type RestNodeOp struct {
	Paging Paging     `json:"paging"`
	Order  Order      `json:"order"`
	Method NodeMethod `json:"method"`
}

func (r RestNodeOp) Name() string {
	return RestNodeType
}


type NodeMethod struct {
	GetOne    NodeMethodOp
	GetList NodeMethodOp
	CreateOne NodeMethodOp
	CreateList NodeMethodOp
	UpdateOne NodeMethodOp
	UpdateList NodeMethodOp
	DeleteOne NodeMethodOp
	DeleteList NodeMethodOp
}

type NodeMethodOp struct {
	Has GenRestSwitch `json:"has"`
	RouterTag string `json:"router_tag"`
	Comments []string `json:"comments"`
}

type Order struct {
	DefaultAcsOrder   []string `json:"default_acs_order"`
	DefaultDescOrder  []string `json:"default_desc_order"`
	OpenOptionalOrder bool     `json:"open_optional_order"`
	OptionalOrder     []string `json:"optional_order"`
}

type Paging struct {
	Open     bool    `json:"open"`
	Must     bool    `json:"must"`
	MaxLimit float64 `json:"max_limit"`
}

func PaseRestNodePaging(m map[string]interface{}) Paging {
	if _, ok := m[RestNodeType]; ok {
		b, err := json.Marshal(m[RestNodeType])
		if err != nil {
			panic(err.Error())
		}
		op := RestNodeOp{}
		err = json.Unmarshal(b, &op)
		if err != nil {
			panic(err.Error())
		}
		return op.Paging
	}
	return Paging{}
}

func CheckMethodHasSwitch(s GenRestSwitch) bool {
	if s == GenRestDefault || s == GenRestTrue {
		return true
	}
	return false
}

func PaseRestNodeMethod(m map[string]interface{}) NodeMethod {
	op := RestNodeOp{}
	if _, ok := m[RestNodeType]; ok {
		b, err := json.Marshal(m[RestNodeType])
		if err != nil {
			panic(err.Error())
		}

		op := RestNodeOp{}

		err = json.Unmarshal(b, &op)
		if err != nil {
			panic(err.Error())
		}
		return op.Method
		//res := map[string]bool{}
		//if op.Method.GetOne == GenRestDefault || op.Method.GetOne == GenRestTrue {
		//	res[GetOneKey] = true
		//}
		//if op.Method.GetList == GenRestDefault || op.Method.GetList == GenRestTrue {
		//	res[GetListKey] = true
		//}
		//if op.Method.CreateOne == GenRestDefault || op.Method.CreateOne == GenRestTrue {
		//	res[CreateOneKey] = true
		//}
		//if op.Method.CreateList == GenRestDefault || op.Method.CreateList == GenRestTrue {
		//	res[CreateListKey] = true
		//}
		//if op.Method.UpdateOne == GenRestDefault || op.Method.UpdateOne == GenRestTrue {
		//	res[UpdateOneKey] = true
		//}
		//
		//if op.Method.UpdateList == GenRestDefault || op.Method.UpdateList == GenRestTrue {
		//	res[UpdateListKey] = true
		//}
		//
		//if op.Method.DeleteOne == GenRestDefault || op.Method.DeleteOne == GenRestTrue {
		//	res[DeleteOneKey] = true
		//}
		//
		//if op.Method.DeleteList == GenRestDefault || op.Method.DeleteList == GenRestTrue {
		//	res[DeleteListKey] = true
		//}
		//
		//return res
	}
	//return map[string]bool{
	//	GetOneKey:     true,
	//	GetListKey:    true,
	//	CreateOneKey:  true,
	//	CreateListKey: true,
	//	UpdateOneKey:  true,
	//	UpdateListKey: true,
	//	DeleteOneKey:  true,
	//	DeleteListKey: true,
	//}
	return op.Method
}
