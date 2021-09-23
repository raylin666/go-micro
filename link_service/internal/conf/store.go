package conf

var (
	conf *store
)

type store struct {
	*Bootstrap
}

// NewStore 初始化配置保存 /**
func NewStore(c *Bootstrap) *store {
	conf = &store{c}
	return conf
}

// GetStore 获取配置保存信息 /**
func GetStore() *store {
	return conf
}


