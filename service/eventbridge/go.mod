module github.com/aws/aws-sdk-go-v2/service/eventbridge

go 1.15

require (
<<<<<<< HEAD
<<<<<<< HEAD
	github.com/aws/aws-sdk-go-v2 v1.19.0
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.35
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.29
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.27
	github.com/aws/smithy-go v1.13.5
=======
	github.com/aws/aws-sdk-go-v2 v1.4.0
	github.com/aws/aws-sdk-go-v2/internal/configsources v0.0.0-00010101000000-000000000000
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.0.0-00010101000000-000000000000
	github.com/aws/aws-sdk-go-v2/internal/v4a v0.0.0-00010101000000-000000000000
	github.com/aws/smithy-go v1.4.0
	github.com/google/go-cmp v0.5.4
>>>>>>> 1b06413dd8 (Add endpoint-based auth scheme resolution (#2158))
=======
	github.com/aws/aws-sdk-go-v2 v1.18.1
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.34
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.28
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.26
	github.com/aws/smithy-go v1.13.5
	github.com/google/go-cmp v0.5.8
>>>>>>> 8bd2a48e49 (fix borked go mod files)
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/internal/v4a => ../../internal/v4a/
