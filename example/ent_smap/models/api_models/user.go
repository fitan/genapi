package api_models

import (
	"ent_samp/ent"
	"ent_samp/gen/entt"
)

type (
	Query struct {
		entt.UserNameEQ
	}

	Uri struct {
		entt.IdUri
	}

	Body struct {
		entt.UserNode
	}

	Header struct {
	}

	In struct {
		Query Query
	}
)
type UserOut struct {
	User ent.Users
	Len  int
}
