package proto

// ZeroBot系统配置
type System struct {
	Environment string `yaml:"env"`
	Port        int    `yaml:"port"`
	Version     string `yaml:"version"`
}

// ZeroBot框架配置
type Config struct {
	System System `yaml:"sys"`
}
