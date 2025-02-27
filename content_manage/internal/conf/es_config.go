package conf

type Elasticsearch struct {
	Addresses     []string `json:"addresses" yaml:"addresses"`
	Username      string   `json:"username" yaml:"username"`
	Password      string   `json:"password" yaml:"password"`
	EnableLogging bool     `json:"enable_logging" yaml:"enable_logging"`
}
