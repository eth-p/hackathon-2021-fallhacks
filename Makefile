PROJECT = github.com/eth-p/hackathon-2021-fallhacks
OUT_DIR = bin

# Target: build
# Build the game for the current platform.
.PHONY: build
build:
	@mkdir -p $(OUT_DIR) || true
	go build -o $(OUT_DIR)/fallsands $(PROJECT)

# Target: build-wasm
# Build the game for webassembly.
build-wasm:
	GOOS=js GOARCH=wasm go build -o "$(OUT_DIR)/wasm/game.wasm" $(PROJECT)
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" "$(OUT_DIR)/wasm/wasm_exec.js"
	cp "main.html" "$(OUT_DIR)/wasm/index.html"

# Target: run
# Run the game.
.PHONY: run
run: build
	$(OUT_DIR)/fallsands
