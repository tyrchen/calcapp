CHECK=\033[32m✔\033[39m
DONE="\n$(CHECK) Done.\n"

ECHO=echo
SUDO=`which sudo`
SUPERVISORCTL=`which supervisorctl`
GO=`which go`
RSYNC=`which rsync`

BIN=./bin
SERVER=zoneke
DEPLOY_PATH=/home/tyr/calcapp

PACKAGES=network utils calc
PACKAGE_PATHS=$(patsubst %,./%, $(PACKAGES))
TARGETS=$(patsubst %.go,$(BIN)/%,$(wildcard *.go))
ALL_FILES=$(shell find . -type f -name '*.go')

.PHONY: build packages clean dir supervisor nginx remote_deploy copy

build: clean packages $(TARGETS)
	@echo $(DONE)

$(TARGETS): $(BIN)/%: %.go
	@echo "building $<..."
	@$(GO) build -o $@ $< 

packages:
	@echo "making packages"
	@$(GO) install $(PACKAGE_PATHS)

supervisor:
	@$(ECHO) "\nUpdate supervisor configuration..."
	@$(SUDO) $(SUPERVISORCTL) reread
	@$(SUDO) $(SUPERVISORCTL) update
	@$(SUDO) $(SUPERVISORCTL) restart calc-*

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

remote_deploy: build copy
	@$(SSH) -t $(SERVER) make supervisor; make nginx"
	@echo $(DONE)

cloc:
	@cloc . --exclude-dir=webclient/assets
