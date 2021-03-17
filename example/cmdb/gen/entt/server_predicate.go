package entt

import (
	"cmdb/ent/predicate"
	"cmdb/ent/server"
	"github.com/gin-gonic/gin"

	"time"
)

func ServerPredicatesExec(fs ...func() (predicate.Server, error)) ([]predicate.Server, error) {
	ps := make([]predicate.Server, 0, len(fs))
	for _, f := range fs {
		p, err := f()
		if err != nil {
			return ps, err
		}
		if p != nil {
			ps = append(ps, p)
		}
	}
	return ps, nil
}

type ServerPaging struct {
	Limit int `form:"limit" json:"limit"`

	Page int `form:"page" json:"page"`
}

func (m *ServerPaging) BindPagingServer(queryer *ent.ServerQuery) error {
	if m.Page == 0 {
		return nil
	}
	queryer.Limit(m.Limit).Offset((m.Page - 1) * m.Limit)
	return nil
}

type ServerCreateTimeEQ struct {
	CreateTimeEQ *time.Time `json:"eq_create_time" form:"eq_create_time"`
}

func (m *ServerCreateTimeEQ) BindServerCreateTimeEQ() (predicate.Server, error) {
	if m.CreateTimeEQ == nil {
		return nil, nil
	}
	return server.CreateTimeEQ(*m.CreateTimeEQ), nil
}

type ServerCreateTimeOr struct {
	CreateTimeOr []time.Time `json:"or_create_time" form:"or_create_time"`
}

func (m *ServerCreateTimeOr) BindServerCreateTimeOr() (predicate.Server, error) {
	if len(m.CreateTimeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Server, 0, len(m.CreateTimeOr))
	for i, _ := range m.CreateTimeOr {
		predicate = append(predicate, server.CreateTimeEQ(m.CreateTimeOr[i]))
	}
	return server.Or(predicate...), nil
}

type ServerCreateTimeNEQ struct {
	CreateTimeNEQ *time.Time `json:"neq_create_time" form:"neq_create_time"`
}

func (m *ServerCreateTimeNEQ) BindServerCreateTimeNEQ() (predicate.Server, error) {
	if m.CreateTimeNEQ == nil {
		return nil, nil
	}
	return server.CreateTimeNEQ(*m.CreateTimeNEQ), nil
}

type ServerCreateTimeIn struct {
	CreateTimeIn []time.Time `json:"in_create_time" form:"in_create_time"`
}

func (m *ServerCreateTimeIn) BindServerCreateTimeIn() (predicate.Server, error) {
	if len(m.CreateTimeIn) == 0 {
		return nil, nil
	}
	return server.CreateTimeIn(m.CreateTimeIn...), nil
}

type ServerCreateTimeNotIn struct {
	CreateTimeNotIn []time.Time `json:"not_in_create_time" form:"not_in_create_time"`
}

func (m *ServerCreateTimeNotIn) BindServerCreateTimeNotIn() (predicate.Server, error) {
	if len(m.CreateTimeNotIn) == 0 {
		return nil, nil
	}
	return server.CreateTimeNotIn(m.CreateTimeNotIn...), nil
}

type ServerCreateTimeGT struct {
	CreateTimeGT *time.Time `json:"gt_create_time" form:"gt_create_time"`
}

func (m *ServerCreateTimeGT) BindServerCreateTimeGT() (predicate.Server, error) {
	if m.CreateTimeGT == nil {
		return nil, nil
	}
	return server.CreateTimeGT(*m.CreateTimeGT), nil
}

type ServerCreateTimeGTE struct {
	CreateTimeGTE *time.Time `json:"gte_create_time" form:"gte_create_time"`
}

func (m *ServerCreateTimeGTE) BindServerCreateTimeGTE() (predicate.Server, error) {
	if m.CreateTimeGTE == nil {
		return nil, nil
	}
	return server.CreateTimeGTE(*m.CreateTimeGTE), nil
}

type ServerCreateTimeLT struct {
	CreateTimeLT *time.Time `json:"lt_create_time" form:"lt_create_time"`
}

func (m *ServerCreateTimeLT) BindServerCreateTimeLT() (predicate.Server, error) {
	if m.CreateTimeLT == nil {
		return nil, nil
	}
	return server.CreateTimeLT(*m.CreateTimeLT), nil
}

type ServerCreateTimeLTE struct {
	CreateTimeLTE *time.Time `json:"lte_create_time" form:"lte_create_time"`
}

func (m *ServerCreateTimeLTE) BindServerCreateTimeLTE() (predicate.Server, error) {
	if m.CreateTimeLTE == nil {
		return nil, nil
	}
	return server.CreateTimeLTE(*m.CreateTimeLTE), nil
}

type ServerUpdateTimeEQ struct {
	UpdateTimeEQ *time.Time `json:"eq_update_time" form:"eq_update_time"`
}

func (m *ServerUpdateTimeEQ) BindServerUpdateTimeEQ() (predicate.Server, error) {
	if m.UpdateTimeEQ == nil {
		return nil, nil
	}
	return server.UpdateTimeEQ(*m.UpdateTimeEQ), nil
}

type ServerUpdateTimeOr struct {
	UpdateTimeOr []time.Time `json:"or_update_time" form:"or_update_time"`
}

func (m *ServerUpdateTimeOr) BindServerUpdateTimeOr() (predicate.Server, error) {
	if len(m.UpdateTimeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Server, 0, len(m.UpdateTimeOr))
	for i, _ := range m.UpdateTimeOr {
		predicate = append(predicate, server.UpdateTimeEQ(m.UpdateTimeOr[i]))
	}
	return server.Or(predicate...), nil
}

type ServerUpdateTimeNEQ struct {
	UpdateTimeNEQ *time.Time `json:"neq_update_time" form:"neq_update_time"`
}

func (m *ServerUpdateTimeNEQ) BindServerUpdateTimeNEQ() (predicate.Server, error) {
	if m.UpdateTimeNEQ == nil {
		return nil, nil
	}
	return server.UpdateTimeNEQ(*m.UpdateTimeNEQ), nil
}

type ServerUpdateTimeIn struct {
	UpdateTimeIn []time.Time `json:"in_update_time" form:"in_update_time"`
}

func (m *ServerUpdateTimeIn) BindServerUpdateTimeIn() (predicate.Server, error) {
	if len(m.UpdateTimeIn) == 0 {
		return nil, nil
	}
	return server.UpdateTimeIn(m.UpdateTimeIn...), nil
}

type ServerUpdateTimeNotIn struct {
	UpdateTimeNotIn []time.Time `json:"not_in_update_time" form:"not_in_update_time"`
}

func (m *ServerUpdateTimeNotIn) BindServerUpdateTimeNotIn() (predicate.Server, error) {
	if len(m.UpdateTimeNotIn) == 0 {
		return nil, nil
	}
	return server.UpdateTimeNotIn(m.UpdateTimeNotIn...), nil
}

type ServerUpdateTimeGT struct {
	UpdateTimeGT *time.Time `json:"gt_update_time" form:"gt_update_time"`
}

func (m *ServerUpdateTimeGT) BindServerUpdateTimeGT() (predicate.Server, error) {
	if m.UpdateTimeGT == nil {
		return nil, nil
	}
	return server.UpdateTimeGT(*m.UpdateTimeGT), nil
}

type ServerUpdateTimeGTE struct {
	UpdateTimeGTE *time.Time `json:"gte_update_time" form:"gte_update_time"`
}

func (m *ServerUpdateTimeGTE) BindServerUpdateTimeGTE() (predicate.Server, error) {
	if m.UpdateTimeGTE == nil {
		return nil, nil
	}
	return server.UpdateTimeGTE(*m.UpdateTimeGTE), nil
}

type ServerUpdateTimeLT struct {
	UpdateTimeLT *time.Time `json:"lt_update_time" form:"lt_update_time"`
}

func (m *ServerUpdateTimeLT) BindServerUpdateTimeLT() (predicate.Server, error) {
	if m.UpdateTimeLT == nil {
		return nil, nil
	}
	return server.UpdateTimeLT(*m.UpdateTimeLT), nil
}

type ServerUpdateTimeLTE struct {
	UpdateTimeLTE *time.Time `json:"lte_update_time" form:"lte_update_time"`
}

func (m *ServerUpdateTimeLTE) BindServerUpdateTimeLTE() (predicate.Server, error) {
	if m.UpdateTimeLTE == nil {
		return nil, nil
	}
	return server.UpdateTimeLTE(*m.UpdateTimeLTE), nil
}

type ServerIPEQ struct {
	IPEQ *string `json:"eq_ip" form:"eq_ip"`
}

func (m *ServerIPEQ) BindServerIPEQ() (predicate.Server, error) {
	if m.IPEQ == nil {
		return nil, nil
	}
	return server.IPEQ(*m.IPEQ), nil
}

type ServerIPOr struct {
	IPOr []string `json:"or_ip" form:"or_ip"`
}

func (m *ServerIPOr) BindServerIPOr() (predicate.Server, error) {
	if len(m.IPOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Server, 0, len(m.IPOr))
	for i, _ := range m.IPOr {
		predicate = append(predicate, server.IPEQ(m.IPOr[i]))
	}
	return server.Or(predicate...), nil
}

type ServerIPNEQ struct {
	IPNEQ *string `json:"neq_ip" form:"neq_ip"`
}

func (m *ServerIPNEQ) BindServerIPNEQ() (predicate.Server, error) {
	if m.IPNEQ == nil {
		return nil, nil
	}
	return server.IPNEQ(*m.IPNEQ), nil
}

type ServerIPIn struct {
	IPIn []string `json:"in_ip" form:"in_ip"`
}

func (m *ServerIPIn) BindServerIPIn() (predicate.Server, error) {
	if len(m.IPIn) == 0 {
		return nil, nil
	}
	return server.IPIn(m.IPIn...), nil
}

type ServerIPNotIn struct {
	IPNotIn []string `json:"not_in_ip" form:"not_in_ip"`
}

func (m *ServerIPNotIn) BindServerIPNotIn() (predicate.Server, error) {
	if len(m.IPNotIn) == 0 {
		return nil, nil
	}
	return server.IPNotIn(m.IPNotIn...), nil
}

type ServerIPGT struct {
	IPGT *string `json:"gt_ip" form:"gt_ip"`
}

func (m *ServerIPGT) BindServerIPGT() (predicate.Server, error) {
	if m.IPGT == nil {
		return nil, nil
	}
	return server.IPGT(*m.IPGT), nil
}

type ServerIPGTE struct {
	IPGTE *string `json:"gte_ip" form:"gte_ip"`
}

func (m *ServerIPGTE) BindServerIPGTE() (predicate.Server, error) {
	if m.IPGTE == nil {
		return nil, nil
	}
	return server.IPGTE(*m.IPGTE), nil
}

type ServerIPLT struct {
	IPLT *string `json:"lt_ip" form:"lt_ip"`
}

func (m *ServerIPLT) BindServerIPLT() (predicate.Server, error) {
	if m.IPLT == nil {
		return nil, nil
	}
	return server.IPLT(*m.IPLT), nil
}

type ServerIPLTE struct {
	IPLTE *string `json:"lte_ip" form:"lte_ip"`
}

func (m *ServerIPLTE) BindServerIPLTE() (predicate.Server, error) {
	if m.IPLTE == nil {
		return nil, nil
	}
	return server.IPLTE(*m.IPLTE), nil
}

type ServerIPContains struct {
	IPContains *string `json:"contains_ip" form:"contains_ip"`
}

func (m *ServerIPContains) BindServerIPContains() (predicate.Server, error) {
	if m.IPContains == nil {
		return nil, nil
	}
	return server.IPContains(*m.IPContains), nil
}

type ServerIPHasPrefix struct {
	IPHasPrefix *string `json:"has_prefix_ip" form:"has_prefix_ip"`
}

func (m *ServerIPHasPrefix) BindServerIPHasPrefix() (predicate.Server, error) {
	if m.IPHasPrefix == nil {
		return nil, nil
	}
	return server.IPHasPrefix(*m.IPHasPrefix), nil

}

type ServerIPHasSuffix struct {
	IPHasSuffix *string `json:"has_suffix_ip" form:"has_suffix_ip"`
}

func (m *ServerIPHasSuffix) BindServerIPHasSuffix() (predicate.Server, error) {
	if m.IPHasSuffix == nil {
		return nil, nil
	}
	return server.IPHasSuffix(*m.IPHasSuffix), nil
}

type ServerMachineTypeEQ struct {
	MachineTypeEQ *server.MachineType `json:"eq_machine_type" form:"eq_machine_type"`
}

func (m *ServerMachineTypeEQ) BindServerMachineTypeEQ() (predicate.Server, error) {
	if m.MachineTypeEQ == nil {
		return nil, nil
	}
	return server.MachineTypeEQ(*m.MachineTypeEQ), nil
}

type ServerMachineTypeOr struct {
	MachineTypeOr []server.MachineType `json:"or_machine_type" form:"or_machine_type"`
}

func (m *ServerMachineTypeOr) BindServerMachineTypeOr() (predicate.Server, error) {
	if len(m.MachineTypeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Server, 0, len(m.MachineTypeOr))
	for i, _ := range m.MachineTypeOr {
		predicate = append(predicate, server.MachineTypeEQ(m.MachineTypeOr[i]))
	}
	return server.Or(predicate...), nil
}

type ServerMachineTypeNEQ struct {
	MachineTypeNEQ *server.MachineType `json:"neq_machine_type" form:"neq_machine_type"`
}

func (m *ServerMachineTypeNEQ) BindServerMachineTypeNEQ() (predicate.Server, error) {
	if m.MachineTypeNEQ == nil {
		return nil, nil
	}
	return server.MachineTypeNEQ(*m.MachineTypeNEQ), nil
}

type ServerMachineTypeIn struct {
	MachineTypeIn []server.MachineType `json:"in_machine_type" form:"in_machine_type"`
}

func (m *ServerMachineTypeIn) BindServerMachineTypeIn() (predicate.Server, error) {
	if len(m.MachineTypeIn) == 0 {
		return nil, nil
	}
	return server.MachineTypeIn(m.MachineTypeIn...), nil
}

type ServerMachineTypeNotIn struct {
	MachineTypeNotIn []server.MachineType `json:"not_in_machine_type" form:"not_in_machine_type"`
}

func (m *ServerMachineTypeNotIn) BindServerMachineTypeNotIn() (predicate.Server, error) {
	if len(m.MachineTypeNotIn) == 0 {
		return nil, nil
	}
	return server.MachineTypeNotIn(m.MachineTypeNotIn...), nil
}

type ServerPlatformTypeEQ struct {
	PlatformTypeEQ *server.PlatformType `json:"eq_platform_type" form:"eq_platform_type"`
}

func (m *ServerPlatformTypeEQ) BindServerPlatformTypeEQ() (predicate.Server, error) {
	if m.PlatformTypeEQ == nil {
		return nil, nil
	}
	return server.PlatformTypeEQ(*m.PlatformTypeEQ), nil
}

type ServerPlatformTypeOr struct {
	PlatformTypeOr []server.PlatformType `json:"or_platform_type" form:"or_platform_type"`
}

func (m *ServerPlatformTypeOr) BindServerPlatformTypeOr() (predicate.Server, error) {
	if len(m.PlatformTypeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Server, 0, len(m.PlatformTypeOr))
	for i, _ := range m.PlatformTypeOr {
		predicate = append(predicate, server.PlatformTypeEQ(m.PlatformTypeOr[i]))
	}
	return server.Or(predicate...), nil
}

type ServerPlatformTypeNEQ struct {
	PlatformTypeNEQ *server.PlatformType `json:"neq_platform_type" form:"neq_platform_type"`
}

func (m *ServerPlatformTypeNEQ) BindServerPlatformTypeNEQ() (predicate.Server, error) {
	if m.PlatformTypeNEQ == nil {
		return nil, nil
	}
	return server.PlatformTypeNEQ(*m.PlatformTypeNEQ), nil
}

type ServerPlatformTypeIn struct {
	PlatformTypeIn []server.PlatformType `json:"in_platform_type" form:"in_platform_type"`
}

func (m *ServerPlatformTypeIn) BindServerPlatformTypeIn() (predicate.Server, error) {
	if len(m.PlatformTypeIn) == 0 {
		return nil, nil
	}
	return server.PlatformTypeIn(m.PlatformTypeIn...), nil
}

type ServerPlatformTypeNotIn struct {
	PlatformTypeNotIn []server.PlatformType `json:"not_in_platform_type" form:"not_in_platform_type"`
}

func (m *ServerPlatformTypeNotIn) BindServerPlatformTypeNotIn() (predicate.Server, error) {
	if len(m.PlatformTypeNotIn) == 0 {
		return nil, nil
	}
	return server.PlatformTypeNotIn(m.PlatformTypeNotIn...), nil
}

type ServerSystemTypeEQ struct {
	SystemTypeEQ *server.SystemType `json:"eq_system_type" form:"eq_system_type"`
}

func (m *ServerSystemTypeEQ) BindServerSystemTypeEQ() (predicate.Server, error) {
	if m.SystemTypeEQ == nil {
		return nil, nil
	}
	return server.SystemTypeEQ(*m.SystemTypeEQ), nil
}

type ServerSystemTypeOr struct {
	SystemTypeOr []server.SystemType `json:"or_system_type" form:"or_system_type"`
}

func (m *ServerSystemTypeOr) BindServerSystemTypeOr() (predicate.Server, error) {
	if len(m.SystemTypeOr) == 0 {
		return nil, nil
	}
	predicate := make([]predicate.Server, 0, len(m.SystemTypeOr))
	for i, _ := range m.SystemTypeOr {
		predicate = append(predicate, server.SystemTypeEQ(m.SystemTypeOr[i]))
	}
	return server.Or(predicate...), nil
}

type ServerSystemTypeNEQ struct {
	SystemTypeNEQ *server.SystemType `json:"neq_system_type" form:"neq_system_type"`
}

func (m *ServerSystemTypeNEQ) BindServerSystemTypeNEQ() (predicate.Server, error) {
	if m.SystemTypeNEQ == nil {
		return nil, nil
	}
	return server.SystemTypeNEQ(*m.SystemTypeNEQ), nil
}

type ServerSystemTypeIn struct {
	SystemTypeIn []server.SystemType `json:"in_system_type" form:"in_system_type"`
}

func (m *ServerSystemTypeIn) BindServerSystemTypeIn() (predicate.Server, error) {
	if len(m.SystemTypeIn) == 0 {
		return nil, nil
	}
	return server.SystemTypeIn(m.SystemTypeIn...), nil
}

type ServerSystemTypeNotIn struct {
	SystemTypeNotIn []server.SystemType `json:"not_in_system_type" form:"not_in_system_type"`
}

func (m *ServerSystemTypeNotIn) BindServerSystemTypeNotIn() (predicate.Server, error) {
	if len(m.SystemTypeNotIn) == 0 {
		return nil, nil
	}
	return server.SystemTypeNotIn(m.SystemTypeNotIn...), nil
}
