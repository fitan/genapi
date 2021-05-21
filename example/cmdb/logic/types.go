package logic


type Policy struct {
	User string `json:"user"`
	Path string `json:"path"`
	Method string `json:"method"`
}

type Policies []Policy

func (p *Policies) Serialize() [][]string {
	ps := make([][]string, 0, len(*p))
	for _, v := range *p {
		ps = append(ps, append([]string{}, v.User,v.Path,v.Method))
	}
	return ps
}


type IdPolicy struct {
	Id int `json:"id"`
	Policy
}

type Uri struct {
	Id int `uri:"id"`
}

type Query struct {
	User string `form:"user"`
}


type GetListIn struct {
	Query Query
}

type GetOneIn struct {
	Uri Uri
}

type AddListIn struct {
	Body []Policy
}

type UpdateBody struct {
	Old Policies `json:"old"`
	New Policies `json:"new"`
}

type UpdateIn struct {
	Body UpdateBody
}

type DeleteIn struct {
	Query Query
}
