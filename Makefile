CHECK=\033[32m✔\033[39m
DONE="\n$(CHECK) Done.\n"

SUDO=/usr/bin/sudo
SUPERVISORCTL=/usr/bin/supervisorctl
GO=go
BIN=./bin
DATA=/var/data/calcapp
PACKAGES=network utils calc
PACKAGE_PATHS=$(patsubst %,./%, $(PACKAGES))
TARGETS=$(patsubst %.go,$(BIN)/%,$(wildcard *.go))
ALL_FILES=$(shell find . -type f -name '*.go')

.PHONY: build packages clean dir supervisor nginx deploy config

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
	@$(ECHO) "\nRestart $(PROJECT)..."
	@$(SUDO) $(SUPERVISORCTL) restart $(PROJECT)

nginx:
	@$(ECHO) "\nRestart nginx..."
	@$(SUDO) /etc/init.d/nginx restart

dir:
	@mkdir -p $(DATA)/bp/origin $(DATA)/bp/new $(DATA)/mp
	@mkdir -p $(DATA)/mac

clean:
	@rm -f $(TARGETS)
	@echo $(DONE)

config:
	@$(SUDO) cp -r webclient/_deploy/etc/. /etc/.

deploy: config supervisor nginx
	@echo $(DONE)
