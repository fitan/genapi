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

func (g *getConfKey) GetTs(key string) *Ts {
	for _, v:= range g.conf.Gen.Ts {
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

func (g *getConfKey) GetCallBack(Name string) *CallBackConf {
	for _, v := range g.conf.Plugin.CallBack {
		if v.TagName == Name {
			return &CallBackConf{CallBack: v}
		}
	}
	return nil
}

type CallBackConf struct {
	CallBack CallBack
}

func (g *getConfKey) GetPoint(Name string) *PointConf {
	for _,v := range g.conf.Plugin.Point {
		if v.TagName == Name {
			return &PointConf{Point: v}
		}
	}
	return nil
}

type PointConf struct {
	Point Point
}

