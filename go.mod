module deployer

go 1.17

require (
	deployer.badassops.com/Initialize v0.0.0
	deployer.badassops.com/Utils v0.0.0
	deployer.badassops.com/Variables v0.0.0

	deployer.badassops.com/Config v0.0.0
	deployer.badassops.com/Msg v0.0.0
	deployer.badassops.com/Logs v0.0.0
	deployer.badassops.com/Api v0.0.0
	deployer.badassops.com/Help v0.0.0
)

require (
	github.com/BurntSushi/toml v0.4.1 // indirect
	github.com/akamensky/argparse v1.3.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace deployer.badassops.com/Variables => ./mod/Variables

replace deployer.badassops.com/Help => ./mod/Help

replace deployer.badassops.com/Initialize => ./mod/Initialize

replace deployer.badassops.com/Utils => ./mod/Utils

replace deployer.badassops.com/Config => ./mod/Config

replace deployer.badassops.com/Msg => ./mod/Msg

replace deployer.badassops.com/Logs => ./mod/Logs

replace deployer.badassops.com/Api => ./mod/Api
