package api_models

import (
	"ent_samp/genent"
)

type (
	Query struct {
		genent.UserNameEQ
	}

	Uri struct {
		genent.IdUri
	}

	Body struct {
		genent.UserNode
	}

	Header struct {
	}

	In struct {
		Query Query
	}
)
