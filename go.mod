module deployer

go 1.17

require (
	api v0.0.0
	config v0.0.0
	initialize v0.0.0
	logs v0.0.0
	utils v0.0.0
	vars v0.0.0
)

require (
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/akamensky/argparse v1.3.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	help v0.0.0 // indirect
	msg v0.0.0 // indirect
)

replace vars => ./mod/vars

replace help => ./mod/help

replace initialize => ./mod/initialize

replace utils => ./mod/utils

replace config => ./mod/config

replace msg => ./mod/msg

replace logs => ./mod/logs

replace api => ./mod/api
