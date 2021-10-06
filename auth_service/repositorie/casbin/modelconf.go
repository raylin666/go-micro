package casbin

import casbinmodel "github.com/casbin/casbin/v2/model"

func RBACModel() (casbinmodel.Model, error) {
	return casbinmodel.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
	`)
}