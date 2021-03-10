# Go parameters
export GOARM=${GOARM:-}

GOOS=linux
BINARY_BASE_NAME=cftsmt

ifeq ($(UNAME_S),Darwin)
  GOOS = darwin
endif


ifeq ($(GOARCH),)
  GOARCH = amd64
  ifneq ($(UNAME_M), x86_64)
    GOARCH = 386
  endif
endif

CUR_D = $(shell pwd)
OUT_D = $(shell echo $${OUT_D:-$(CUR_D)/builds})

.PHONY: all
all: test 
	@make build_all

.PHONY: test
test: 
	go test -v ./...

.PHONY: clean
clean: 
	rm -rf builds

clean_recordings:
	rm -rf recordings
	mkdir recordings

.PHONY: build
build: _build
	@echo "==> build local version"
	@echo ""
	@mv $(OUT_D)/${BINARY_BASE_NAME}_$(GOOS)_$(GOARCH) $(OUT_D)/${BINARY_BASE_NAME}
	@echo "installed as $(OUT_D)/${BINARY_BASE_NAME}"

.PHONY: build_linux_amd64
build_linux_amd64:
build_linux_amd64: GOOS=linux
build_linux_amd64: GOARCH=amd64
build_linux_amd64: GOARM=
build_linux_amd64: _build
	@echo "==> building amd64 with GOOS=linux GOARCH=amd64 GOARM="
	@mv $(OUT_D)/${BINARY_BASE_NAME}_$(GOOS)_$(GOARCH) $(OUT_D)/${BINARY_BASE_NAME}_amd64
	@echo "installed as $(OUT_D)/${BINARY_BASE_NAME}_amd64"

.PHONY: build_linux_arm6
build_linux_arm6:
build_linux_arm6: GOOS=linux
build_linux_arm6: GOARCH=arm
build_linux_arm6: GOARM=6
build_linux_arm6: _build
	@echo "==> building arm6 with GOOS=linux GOARCH=arm6 GOARM="
	@mv $(OUT_D)/${BINARY_BASE_NAME}_$(GOOS)_$(GOARCH) $(OUT_D)/${BINARY_BASE_NAME}_arm6
	@echo "installed as $(OUT_D)/${BINARY_BASE_NAME}_arm6"

.PHONY: _build
_build:
	@echo "=> building ${BINARY_BASE_NAME} via go build"
	@echo ""
	@OUT_D=${OUT_D} GOOS=${GOOS} GOARCH=${GOARCH} GOARM=${GOARM} ./scripts/build.sh
	@echo "built $(OUT_D)/${BINARY_BASE_NAME}_$(GOOS)_$(GOARCH)"

.PHONY: build_all
build_all: clean
	@make build 
	@make build_linux_arm6
	@make build_linux_amd64

build_dry_run:
	 goreleaser --snapshot --skip-publish --rm-dist
