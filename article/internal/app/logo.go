package app

import (
	"fmt"
	"github.com/fatih/color"
	"mt/config"
	"runtime"
)

// NewLogo 打印项目服务启动信息
func NewLogo(config *config.Bootstrap) {
	// see https://patorjk.com/software/taag/#p=testall&f=Graffiti&t=mt
	var logo = `
	  
		███╗   ███╗	     ████████╗
		████╗ ████║	     ╚══██╔══╝
		██╔████╔██║   		██║   
		██║╚██╔╝██║   		██║   
		██║ ╚═╝ ██║   		██║   
		╚═╝     ╚═╝   		╚═╝
`
	color.HiYellow(logo)

	contents := fmt.Sprintf(`
GO 版本及路径: %s (%s)
系统类型及架构: %s (%s) - %d 核 CPU
服务名称: %s (%s)
服务版本: %s
服务环境: %s
	`,
		runtime.Version(),
		runtime.GOROOT(),
		runtime.GOOS,
		runtime.GOARCH,
		runtime.GOMAXPROCS(0),

		config.App.Name,
		config.App.Id,
		config.App.Version,
		config.Environment,
	)

	color.HiGreen(contents)
}
