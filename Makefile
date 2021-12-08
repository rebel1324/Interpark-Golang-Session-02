OUTPUT_DIR := ./output/
OUTPUT_NAME := main.out
OUTPUT_PATH := ${OUTPUT_DIR}${EXECUTABLE_NAME}
GO_SRC_PATH := cmd

build:
	$(MAKE) -C "$(GO_SRC_PATH)" build output_dir="$(OUTPUT_DIR)/$(OUTPUT_NAME)"
run:
	"$(OUTPUT_DIR)/$(OUTPUT_NAME)"
docker:
	docker-compose build
