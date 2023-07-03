
OS_ID = generic
MACHINE = generic

UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)

ifeq ($(UNAME_S),Linux)
	OS_ID = Linux_$(UNAME_M)
	GID = god
endif
ifeq ($(UNAME_S),Darwin)
	OS_ID = Darwin_$(UNAME_M)
	GID = staff
endif

CODE_NAME = deployer

SOURCES = $(CODE_NAME).go \
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

all:	release/$(CODE_NAME)_$(OS_ID) \
		release/$(CODE_NAME)_$(OS_ID).tar.gz \
		release/$(CODE_NAME)_$(OS_ID).sha256

release/$(CODE_NAME)_$(OS_ID): $(BUILT_SOURCES)
	@echo "build the $(CODE_NAME)_$(OS_ID) binary..."
	@go build -o release/$(CODE_NAME)_$(OS_ID) $(CODE_NAME).go
	@echo "set owner and strip the binary"
	@chown luc:$(GID) release/$(CODE_NAME)_$(OS_ID)
	@strip release/$(CODE_NAME)_$(OS_ID)

release/$(CODE_NAME)_$(OS_ID).tar.gz: release/$(CODE_NAME)_$(OS_ID)
	@echo "create the $(CODE_NAME)_$(OS_ID).tar.gz archive..."
	@(cd release ; tar zcf $(CODE_NAME)_$(OS_ID).tar.gz $(CODE_NAME)_$(OS_ID))

release/$(CODE_NAME)_$(OS_ID).sha256: release/$(CODE_NAME)_$(OS_ID).tar.gz
	@echo "create the sha256 information file..."
	@sha256sum release/$(CODE_NAME)_$(OS_ID).tar.gz | awk '{print $$1}' > release/$(CODE_NAME)_$(OS_ID).sha256
	@echo "SHA256: $$(cat release/$(CODE_NAME)_$(OS_ID).sha256)"

install: release/$(CODE_NAME)_$(OS_ID)
	@echo "Installing the new $(CODE_NAME) binary..."
	@sudo cp release/$(CODE_NAME)_$(OS_ID) /usr/local/sbin/$(CODE_NAME)
	@sudo chmod 0755 /usr/local/sbin/$(CODE_NAME)
	@sudo chown 0:0 /usr/local/sbin/$(CODE_NAME)

clean:
	@rm -f release/*$(OS_ID)*

changelog:
	@echo "version built $(TOOL_VERSION)"
