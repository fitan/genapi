package nest

type Nest struct {
	School  string
	Parents string `json:"parents"`
	Depth DepthNest `json:"depth"`
}

type DepthNest struct {
	Hello string `json:"hello"`
}

type Fater string
