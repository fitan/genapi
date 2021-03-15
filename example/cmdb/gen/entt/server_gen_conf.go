package entt

import (
	"cmdb/ent"
	"cmdb/ent/server"
)

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

	m.AddServiceIDs(curd.ServiceObj.GetIDs(v.Edges.Services)...)

}

func ServerUpdateMutation(m *ent.ServerMutation, v *ent.Server) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetIP(v.IP)

	m.SetMachineType(v.MachineType)

	m.SetPlatformType(v.PlatformType)

	m.SetSystemType(v.SystemType)

	m.AddServiceIDs(curd.ServiceObj.GetIDs(v.Edges.Services)...)

}
