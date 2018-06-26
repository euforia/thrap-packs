
NAME = ${stack.id}
BUILDTIME = $(shell date +%Y-%m-%dT%T%z)

# These options are to allow for cross-platform support
DIST_OPTS = -a -tags netgo -installsuffix netgo
# This sets the version and build time in the binary
LD_OPTS = -ldflags="-X main._version=$(VERSION) -X main._buildtime=$(BUILDTIME) -w"
# CGO_ENABLED is set to be zero to cross-platform support
BUILD_CMD = CGO_ENABLED=0 go build $(LD_OPTS)

SOURCE_FILES = $(shell ls ./cmd/*.go | grep -v _test.go)

clean:
	rm -rf dist
	rm -f $(NAME)

test:
	go test -cover $(shell go list ./... | grep -v /vendor/)

$(NAME):
	$(BUILD_CMD) -o $(NAME) $(SOURCE_FILES)

dist:
	mkdir dist
	GOOS=linux $(BUILD_CMD) $(DIST_OPTS) -o dist/$(NAME)-linux $(SOURCE_FILES)
	GOOS=darwin $(BUILD_CMD) $(DIST_OPTS) -o dist/$(NAME)-darwin $(SOURCE_FILES)