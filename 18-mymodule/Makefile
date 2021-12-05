#Docker-compose

.PHONY: help

help:
	    @echo "Makefile arguments:"
	    @echo "Makefile commands:"
		@echo "make"
	    @echo "make destroy"
	    @echo "make dev"

.DEFAULT_GOAL := run

run:
	    @docker-compose up -d --build
destroy:
		@docker-compose down
dev:
		@docker build --target dev . -t web
		@docker run -v ${PWD}:/work -p 8080:8080 -it web sh
