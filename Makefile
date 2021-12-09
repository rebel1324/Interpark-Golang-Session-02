OUTPUT_DIR := output/
OUTPUT_NAME := server
OUTPUT_PATH := ${OUTPUT_DIR}${OUTPUT_NAME}
SRC_DIR := cmd

build:
	$(MAKE) -C ${SRC_DIR} build output_path="../${OUTPUT_PATH}"
run:
	./${OUTPUT_PATH}
docker:
	docker-compose build
docker-run:
	docker-compose up --build
docker-daemon:
	docker-compose up --build -d
