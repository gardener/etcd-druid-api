module github.com/gardener/etcd-druid-api

go 1.20

// Unfortunately go does not clearly distinguish between compile-time/test/tool dependencies.
// To demarcate this clearly contributors should clearly seggreate them into separate `require` blocks.
// This will help the maintainers clearly understand the category of dependencies for this project.
// Dependencies for an API project should be kept to the minimum.

// The following dependencies are compile-time dependencies.
require (
	k8s.io/api v0.26.2
	k8s.io/apimachinery v0.26.2
)

// The following dependencies are not compile-time dependencies.
// This is a bit unfortunate that this gets added as a dependency and that is due to the usage of
// init function for which an empty import needs to be included for each of these dependencies.
require (
	// required to generate api-docs.
	github.com/ahmetb/gen-crd-api-reference-docs v0.3.0
	// required to generate manifests, DeepCopy functions etc.
	sigs.k8s.io/controller-tools v0.11.3
)

// The following dependencies are test dependencies
require (
	github.com/onsi/gomega v1.24.2
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/gobuffalo/flect v0.3.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/tools v0.4.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apiextensions-apiserver v0.26.1 // indirect
	k8s.io/gengo v0.0.0-20220902162205-c0856e24416d // indirect
	k8s.io/klog v0.2.0 // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
	k8s.io/utils v0.0.0-20221107191617-1a15be271d1d // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
