package types

type AuthConfig struct {
	Token   string `yaml:"token,omitempty"`
	TokenAt string `yaml:"tokenAt,omitempty"`
	OrgId   int    `yaml:"orgId"`
}

type Config struct {
	Endpoint        string     `yaml:"endpoint"`
	Username        string     `yaml:"username"`
	Password        string     `yaml:"password"`
	ApplicationType int        `yaml:"applicationType"`
	Auth            AuthConfig `yaml:"auth,omitempty"`
}
