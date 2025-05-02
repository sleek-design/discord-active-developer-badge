BINARY_NAME=bot
PLATFORMS=darwin linux windows
ARCHITECTURES=amd64 arm64

all: clean build package

build: $(foreach PLATFORM,$(PLATFORMS),$(foreach ARCH,$(ARCHITECTURES),build-$(PLATFORM)-$(ARCH)))

build-%-amd64:
	@echo "Building $(BINARY_NAME)-$*-amd64"
	@GOOS=$* GOARCH=amd64 go build -o dist/$(BINARY_NAME)-$*-amd64 .
	@echo "$(BINARY_NAME)-$*-amd64: $$(sha256sum dist/$(BINARY_NAME)-$*-amd64 | awk '{print $$1}')"

build-%-arm64:
	@echo "Building $(BINARY_NAME)-$*-arm64"
	@GOOS=$* GOARCH=arm64 go build -o dist/$(BINARY_NAME)-$*-arm64 .
	@echo "$(BINARY_NAME)-$*-arm64: $$(sha256sum dist/$(BINARY_NAME)-$*-arm64 | awk '{print $$1}')"

package: $(foreach PLATFORM,$(PLATFORMS),$(foreach ARCH,$(ARCHITECTURES),package-$(PLATFORM)-$(ARCH)))

package-%-amd64:
	@echo "Packaging $(BINARY_NAME)-$*-amd64"
	@if [ "$*" = "windows" ]; then \
    		mv dist/$(BINARY_NAME)-$*-amd64 dist/$(BINARY_NAME)-$*-amd64.exe; \
			zip -j dist/$(BINARY_NAME)-$*-amd64.zip dist/$(BINARY_NAME)-$*-amd64.exe; \
		else \
			chmod +x dist/$(BINARY_NAME)-$*-amd64; \
			tar -zcvf dist/$(BINARY_NAME)-$*-amd64.tar.gz dist/$(BINARY_NAME)-$*-amd64; \
	fi

package-%-arm64:
	@echo "Packaging $(BINARY_NAME)-$*-arm64"
	@if [ "$*" = "windows" ]; then \
    		mv dist/$(BINARY_NAME)-$*-arm64 dist/$(BINARY_NAME)-$*-arm64.exe; \
			zip -j dist/$(BINARY_NAME)-$*-arm64.zip dist/$(BINARY_NAME)-$*-arm64.exe; \
		else \
			chmod +x dist/$(BINARY_NAME)-$*-arm64; \
			tar -zcvf dist/$(BINARY_NAME)-$*-arm64.tar.gz dist/$(BINARY_NAME)-$*-arm64; \
	fi

clean:
	@rm -f dist/$(BINARY_NAME)-*

.PHONY: all build clean package