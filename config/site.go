package config

import "reflect"

type SiteInfo struct {
	Name       string `yaml:"name" json:"name"`
	CreateAt   string `yaml:"create_at" json:"create_at"`
	Title      string `yaml:"title" json:"title"`
	WebVersion string `yaml:"web_version" json:"web_version"`
	Addr       string `yaml:"addr" json:"addr"`
	Email      string `yaml:"email" json:"email"`
}

func (s *SiteInfo) Update(newInfo SiteInfo) {
	val := reflect.ValueOf(newInfo)
	valS := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		newValue := val.Field(i)
		if newValue.Interface() != reflect.Zero(newValue.Type()).Interface() {
			valS.Field(i).Set(newValue)
		}
	}
}
