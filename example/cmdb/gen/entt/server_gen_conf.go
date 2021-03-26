package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
	"cmdb/ent/server"
)

type ServerIncludes struct {
	Includes []string `form:"includes" json:"includes" binding:"dive,oneof=service.role_binding.user.alert service service.project service.project.role_binding service.project.role_binding.user service.project.role_binding.user.alert service.role_binding.user service.role_binding"`
}

func ServerSelete(queryer *ent.ServerQuery) {
	queryer.Select(

		server.FieldCreateTime,

		server.FieldUpdateTime,

		server.FieldIP,

		server.FieldMachineType,

		server.FieldPlatformType,

		server.FieldSystemType,
	)
}

func ServerCreateMutation(m *ent.ServerMutation, v *ent.Server) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetIP(v.IP)

	m.SetMachineType(v.MachineType)

	m.SetPlatformType(v.PlatformType)

	m.SetSystemType(v.SystemType)

	m.AddServiceIDs(ServiceGetIDs(v.Edges.Services)...)

}

func ServerUpdateMutation(m *ent.ServerMutation, v *ent.Server) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetIP(v.IP)

	m.SetMachineType(v.MachineType)

	m.SetPlatformType(v.PlatformType)

	m.SetSystemType(v.SystemType)

	m.AddServiceIDs(ServiceGetIDs(v.Edges.Services)...)

}

func ServerGetIDs(servers ent.Servers) []int {
	IDs := make([]int, 0, len(servers))
	for _, server := range servers {
		IDs = append(IDs, server.ID)
	}
	return IDs
}

type ServerDefaultQuery struct {
	ServerIncludes
}

func (s *ServerDefaultQuery) PredicatesExec() ([]predicate.Server, error) {
	return ServerPredicatesExec()
}

func (s *ServerDefaultQuery) Exec(queryer *ent.ServerQuery) error {
	ps, err := s.PredicatesExec()
	if err != nil {
		return err
	}
	QueryerIncludes(queryer, s.Includes)

	queryer.Where(server.And(ps...))

	return nil
}
