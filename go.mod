module vrlbench

go 1.18

replace github.com/mingrammer/flog => ./flog

replace github.com/gh123man/go-vrl/v10 => ../go-vrl/v10

replace github.com/gh123man/go-vrl/v5 => ../go-vrl/v5

require (
	github.com/dustin/go-humanize v1.0.0
	github.com/gh123man/go-vrl/v5 v5.0.0
	github.com/mingrammer/flog v0.0.0
	go.uber.org/atomic v1.10.0
)

require (
	github.com/brianvoe/gofakeit v3.18.0+incompatible // indirect
	github.com/fatih/color v1.6.0 // indirect
	github.com/mattn/go-colorable v0.0.9 // indirect
	github.com/mattn/go-isatty v0.0.3 // indirect
	github.com/mingrammer/cfmt v1.0.0 // indirect
	github.com/spf13/pflag v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20180416112224-2f57af4873d0 // indirect
)
