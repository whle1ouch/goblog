package config

type SiteInfo struct {
	Name       string `yaml:"name" json:"name"`
	CreateAt   string `yaml:"create_at" json:"create_at"`
	Title      string `yaml:"title" json:"title"`
	WebVersion string `yaml:"web_version" json:"web_version"`
	Addr       string `yaml:"addr" json:"addr"`
	Email      string `yaml:"email" json:"email"`
}
