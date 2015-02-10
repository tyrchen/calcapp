CHECK=\033[32m✔\033[39m
DONE="\n$(CHECK) Done.\n"

ECHO=echo
CP=`which cp`
SUDO=`which sudo`
SUPERVISORCTL=`which supervisorctl`
GO=`which go`
RSYNC=`which rsync`

BIN=./bin
SERVER=weixin
DEPLOY_PATH=/home/tchen/calcapp

PACKAGES=network utils calc calcv2 calcv3
PACKAGE_PATHS=$(patsubst %,./%, $(PACKAGES))
TARGETS=$(patsubst %.go,$(BIN)/%,$(wildcard *.go))
ALL_FILES=$(shell find . -type f -name '*.go')

.PHONY: build packages clean dir supervisor nginx remote_deploy copy data

build: data packages $(TARGETS)
	@echo $(DONE)

data:
	@$(RSYNC) -au data $(BIN)

$(TARGETS): $(BIN)/%: %.go
	@echo "building $<..."
	@GOOS=$(GOOS) $(GO) build -o $@ $<

packages:
	@echo "making packages"
	@GOOS=$(GOOS) $(GO) install $(PACKAGE_PATHS)

supervisor:
	@$(ECHO) "\nUpdate supervisor configuration..."
	@$(SUDO) $(SUPERVISORCTL) reread
	@$(SUDO) $(SUPERVISORCTL) update

nginx:
	@$(ECHO) "\nRestart nginx..."
	@$(SUDO) /etc/init.d/nginx restart

clean:
	@rm -f $(TARGETS)
	@echo $(DONE)

copy:
	@$(RSYNC) -avu Makefile $(SERVER):$(DEPLOY_PATH)/
	@$(RSYNC) -avu webclient/_deploy/ $(SERVER):$(DEPLOY_PATH)/
	@$(RSYNC) -avu $(BIN) $(SERVER):$(DEPLOY_PATH)/
	@$(RSYNC) -avu webclient $(SERVER):$(DEPLOY_PATH)/

remote_deploy: GOOS=linux
remote_deploy: build copy
	@ssh -t $(SERVER) "cd $(DEPLOY_PATH); make supervisor; make nginx"
	@echo $(DONE)

cloc:
	@cloc . --exclude-dir=webclient/assets
