# TODO: Remplacez "template" par le nom de votre plugin (ex: "discord", "aws")
PLUGIN_NAME=template

# Noms des binaires et des dossiers
PLUGIN_BINARY=orkestra-plugin-$(PLUGIN_NAME)
DIST_DIR=dist

.PHONY: all build clean rebuild

all: build

build:
	@echo "Building plugin: $(PLUGIN_BINARY)..."
	@mkdir -p $(DIST_DIR)
	go build -o $(DIST_DIR)/$(PLUGIN_BINARY) .
	@echo "Plugin built successfully in $(DIST_DIR)/"

clean:
	@echo "Cleaning up..."
	@rm -rf $(DIST_DIR)
	@echo "Cleanup complete."

rebuild: clean all

