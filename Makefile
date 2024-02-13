REPO_ROOT           := $(shell dirname "$(realpath $(lastword $(MAKEFILE_LIST)))")

#########################################
# Tools                                 #
#########################################
TOOLS_DIR := hack/tools
include $(REPO_ROOT)/hack/tools.mk

##################################################################
# Rules for formatting, generation, linting, build, testing etc. #
##################################################################

.PHONY: fmt
fmt:
	@env GO111MODULE=on go fmt ./...

.PHONY: check
check: $(GOLANGCI_LINT)
	@hack/check.sh --golangci-lint-config=./.golangci.yaml ./core/...

.PHONY: generate
generate: $(CONTROLLER_GEN) $(GEN_CRD_API_REFERENCE_DOCS)
	@hack/generate.sh

.PHONY: add-license-headers
add-license-headers: $(GO_ADD_LICENSE)
	@./hack/add_license_headers.sh

.PHONY: revendor
revendor:
	@env GO111MODULE=on go mod tidy
	@env GO111MODULE=on go mod vendor

.PHONY: test
test:
	@./hack/test.sh

