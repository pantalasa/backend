module earthly/example-go
go 1.22
toolchain go1.22.10

require (
	github.com/aws/aws-sdk-go v1.44.298
	github.com/aws/aws-sdk-go-v2 v1.24.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.47.5
	github.com/sirupsen/logrus v1.9.0
	golang.org/x/exp v0.0.0-20230206171751-46f607a40771
)

require golang.org/x/sys v0.1.0 // indirect
