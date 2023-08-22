package cfg

type configType struct {
	SNI    string `mapstructure:"sni" default:"www.cloudflare.com"`
	DbPath string `mapstructure:"dbPath" default:"scan.db"`
}

var Instance = &configType{}
