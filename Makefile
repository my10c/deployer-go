
OS_ID = generic
MACHINE = generic

UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)

ifeq ($(UNAME_S),Linux)
	OS_ID = Linux_$(UNAME_M)
endif
ifeq ($(UNAME_S),Darwin)
	OS_ID = Darwin_$(UNAME_M)
endif

SOURCES = deployer.go \
	mod/initialize/initialize.go \
	mod/utils/utils.go \
	mod/vars/vars.go \
	mod/msg/msg.go \
	mod/logs/logs.go \
	mod/help/help.go \
	mod/config/config.go \
	mod/api/api.go

BUILT_SOURCES = $(SOURCES)
TOOL_VERSION := $(shell cat mod/vars/vars.go | grep MyVersion | egrep -v MyProgname | awk '{print $$3}')

all:	release/deployer_$(OS_ID) \
		release/deployer_$(OS_ID).tar.gz \
		release/deployer_$(OS_ID).sha256

release/deployer_$(OS_ID): $(BUILT_SOURCES)
	@echo "build the deployer_$(OS_ID) binary..."
	@go build -o release/deployer_$(OS_ID) deployer.go

release/deployer_$(OS_ID).tar.gz: release/deployer_$(OS_ID)
	@echo "create the deployer_$(OS_ID).tar.gz archive..."
	@(cd release ; tar zcf deployer_$(OS_ID).tar.gz deployer_$(OS_ID))

release/deployer_$(OS_ID).sha256: release/deployer_$(OS_ID).tar.gz
	@echo "create the sha256 information file..."
	@sha256sum release/deployer_$(OS_ID).tar.gz | awk '{print $$1}' > release/deployer_$(OS_ID).sha256
	@echo "SHA256: $$(cat release/deployer_$(OS_ID).sha256)"

install: release/deployer_$(OS_ID)
	@echo "Installing the new deployer binary..."
	@sudo cp release/deployer_$(OS_ID) /usr/local/sbin/deployer
	@sudo chmod 0755 /usr/local/sbin/deployer
	@sudo chown 0:0 /usr/local/sbin/deployer

clean:
	@rm -f release/*$(OS_ID)*

changelog:
	@echo "version built $(TOOL_VERSION)"
