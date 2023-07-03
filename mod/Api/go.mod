module api

go 1.17

require (
	deployer.badassops.com/Config v0.0.0
	deployer.badassops.com/Logs v0.0.0
	deployer.badassops.com/Msg v0.0.0
)

require (
	deployer.badassops.com/Variables v0.0.0 // indirect
	github.com/BurntSushi/toml v0.4.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace deployer.badassops.com/Variables => ../Variables

replace deployer.badassops.com/Config => ../Config

replace deployer.badassops.com/Msg => ../Msg

replace deployer.badassops.com/Logs => ../Logs
