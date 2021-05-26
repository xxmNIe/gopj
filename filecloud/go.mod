module fliecloud

go 1.16

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/StackExchange/wmi v0.0.0-20210224194228-fe8f1750fd46 // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/shirou/gopsutil v3.21.4+incompatible
	github.com/yddeng/dnet v0.0.0-20210517045540-29976bdf5083
	github.com/yddeng/dutil v0.0.0-20210518033953-18807197300c
	github.com/yddeng/filecloud v0.0.0-20210414054422-4e5902e3a463

)

replace (
	github.com/yddeng/filecloud  => ./
	github.com/yddeng/dnet => ../dnet
	github.com/yddeng/dutil => ../utils
)
