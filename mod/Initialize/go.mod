module Initialize

go 1.17

require (
	deployer.badassops.com/Help v0.0.0
	deployer.badassops.com/Utils v0.0.0
	deployer.badassops.com/Variables v0.0.0
	github.com/akamensky/argparse v1.3.1
)

replace deployer.badassops.com/Variables => ../Variables

replace deployer.badassops.com/Help => ../Help

replace deployer.badassops.com/Utils => ../Utils
