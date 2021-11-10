module fliecloud

go 1.16

replace (
	github.com/yddeng/dnet => ../dnet
	github.com/yddeng/dutil => ../utils
	github.com/yddeng/filecloud => ./
)

require (
	github.com/BurntSushi/toml v0.4.1
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/shirou/gopsutil v3.21.7+incompatible
	github.com/yddeng/dnet v0.0.0-00010101000000-000000000000
	github.com/yddeng/dutil v0.0.0-00010101000000-000000000000
	github.com/yddeng/filecloud v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
)
