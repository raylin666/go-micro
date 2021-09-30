package casbin

import (
	"auth_service/internal/conf"
	"fmt"
	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var (
	enforcer *casbin.Enforcer
)

type Options struct {

}

func NewEnforcer(modelType string) error {
	source := strings.Split(conf.GetStore().GetData().GetDatabase().GetSource(), "/")
	adapter, err := gormadapter.NewAdapter(
		conf.GetStore().GetData().GetDatabase().GetDriver(),
		fmt.Sprintf("%s/", source[0]),
		)
	if err != nil {
		return err
	}

	var model casbinmodel.Model
	switch modelType {
	case "rbac":
		model, err = RBACModel()
		break
	}
	if err != nil {
		return err
	}

	enforcer, err = casbin.NewEnforcer(model, adapter)
	return err
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}