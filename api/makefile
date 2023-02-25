TARGET = dockertainer
BUILD_PATH ?= build

VERSION := $(shell git describe --tags --always)
PACKAGE_NAME := dockertainer-$(VERSION)

all: $(TARGET)

.PHONY: build_prepare
build_prepare:
	@mkdir -p $(BUILD_PATH)/$(PACKAGE_NAME)

.PHONY: $(TARGET)
$(TARGET): build_prepare
	go build -gcflags="all=-N -l" -o $(TARGET) main.go

	@mv ./$(TARGET) $(BUILD_PATH)/$(PACKAGE_NAME)/
	@echo "Build successfully"

.PHONY: pack
pack:
	@if [ -e out ] ; then rm -rf out; fi
	@mkdir out
	@cp $(TARGET) ./out/$(TARGET)
	tar -C out -zcf $(TARGET)-v$(VERSION).tar.gz .

.PHONY: clean
clean:
	@rm -rf ./bin
	@rm -rf ./out
	@rm -rf ./$(TARGET)
	go clean -i .
