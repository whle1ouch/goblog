package flags

import (
	sys_flag "flag"
	"goblog/global"
)

type Option struct {
	Version bool
	DB      bool
}

func Parse() Option {
	version := sys_flag.Bool("v", false, "查看版本")
	db := sys_flag.Bool("db", false, "初始化数据库")
	sys_flag.Parse()
	return Option{
		Version: *version,
		DB:      *db,
	}
}

func (o Option) IsServerStop() bool {
	return o.DB || o.Version
}

func (o Option) Migration() {
	if o.DB {
		MakeMigrations()
	}
}

func (o Option) ShowVersion() {
	if o.Version {
		println("Current GoBlog Version: ", global.Config.System.Version)
	}
}
