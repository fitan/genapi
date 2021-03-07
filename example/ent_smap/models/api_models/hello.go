package api_models

import genrest "ent_samp/service"

type (
Query struct {
	genrest.UserNameEQ
}

Uri struct {
	genrest.IdUri
}

Body struct {
	genrest.UserNode
}

Header struct {

}

In struct {
	Query Query
	Uri Uri
	Body Body
}
)

