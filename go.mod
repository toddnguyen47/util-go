module github.com/toddnguyen47/util-go

go 1.19

// Retracting versions that no longer work
retract v1.0.0

retract v1.1.0

retract v1.2.0

retract v1.3.0

retract v1.4.0

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/aws/aws-sdk-go v1.44.302
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.10.31
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression v1.4.58
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.20.1
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.18.3
	github.com/google/go-cmp v0.5.9
	github.com/onsi/ginkgo/v2 v2.5.1
	github.com/onsi/gomega v1.24.1
	github.com/rs/zerolog v1.29.1
	github.com/stretchr/testify v1.8.4
	golang.org/x/crypto v0.11.0
)

require (
	github.com/aws/aws-sdk-go-v2 v1.19.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.35 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.29 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.14.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.29 // indirect
	github.com/aws/smithy-go v1.13.5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
