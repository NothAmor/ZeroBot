package proto

// ZeroBot系统配置
type (
	System struct {
		Environment string `yaml:"env"`     // 环境, dev/prod
		Port        int    `yaml:"port"`    // 端口
		Version     string `yaml:"version"` // 版本
	}
)

// ZeroBot插件配置
type (
	Plugins struct {
		Command   string `yaml:"command"`   // 命令前缀
		Separator string `yaml:"separator"` // 命令分隔符
	}
)

// 机器人配置
type (
	Bot struct {
		QQ     int64  `yaml:"qq"`     // QQ号
		Name   string `yaml:"name"`   // 名字
		Age    int    `yaml:"age"`    // 年龄
		Gender string `yaml:"gender"` // 性别
	}
)

// ZeroBot通知配置
type (
	ServerChan struct {
		Key string `yaml:"key"` // ServerChan key
	}
	Email struct {
		Host      string   `yaml:"host"`      // SMTP服务器地址
		Account   string   `yaml:"account"`   // SMTP账号
		Password  string   `yaml:"password"`  // SMTP密码
		Port      int      `yaml:"port"`      // SMTP端口
		Receivers []string `yaml:"receivers"` // 接收者
	}
	Notify struct {
		Enable     bool       `yaml:"enable"`      // 是否启用推送通知
		Use        []string   `yaml:"use"`         // 使用的推送通知方式
		ServerChan ServerChan `yaml:"server-chan"` // ServerChan配置
		Email      Email      `yaml:"email"`       // Email配置
	}
)

// ZeroBot框架配置
type Config struct {
	System  System  `yaml:"sys"`     // 系统配置
	Bot     Bot     `yaml:"bot"`     // 机器人配置
	Plugins Plugins `yaml:"plugins"` // 插件配置
	Notify  Notify  `yaml:"notify"`  // 通知配置
}
