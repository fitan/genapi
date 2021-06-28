package public

func GetConfKey() *getConfKey {
	return &getConfKey{GetGenConf()}
}

type getConfKey struct {
	conf *GenConf
}

func (g *getConfKey) GetApi(key string) *API {
	for _, v := range g.conf.Gen.API {
		if v.Name == key {
			return &v
		}
	}
	return nil
}

func (g *getConfKey) GetEnt(key string) *Ent {
	for _, v := range g.conf.Gen.Ent {
		if v.Name == key {
			return &v
		}
	}
	return nil
}

func (g *getConfKey) GetPlugin(key string) *plugin {
	for _, v := range g.conf.Plugin {
		if v.Name == key {
			return &plugin{conf: v}
		}
	}
	return nil
}

type plugin struct {
	conf Plugin
}

func (p *plugin) GetInInterface(key string) *InBindInterfaceName {
	for _, v := range p.conf.InBindInterfaceName {
		if v.Name == key {
			return &v
		}
	}
	return nil
}

func (p *plugin) GetOutInterface(key string) *OutBindInterfaceName {
	for _, v := range p.conf.OutBindInterfaceName {
		if v.Name == key {
			return &v
		}
	}
	return nil
}
