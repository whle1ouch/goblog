package config

import "reflect"

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	JWT      JWT      `yaml:"jwt"`
	QQInfo   QQInfo   `yaml:"qq"`
}

func Update[T SiteInfo | JWT | QQInfo](oldS *T, newS T) {
	valNew := reflect.ValueOf(newS)
	valS := reflect.ValueOf(oldS).Elem()
	for i := 0; i < valNew.NumField(); i++ {
		newValue := valNew.Field(i)
		if newValue.Interface() != reflect.Zero(newValue.Type()).Interface() {
			valS.Field(i).Set(newValue)
		}
	}
}
