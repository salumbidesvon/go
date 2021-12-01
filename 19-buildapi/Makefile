#Docker-compose

.PHONY: help

help:
	    @echo "Makefile arguments:"
	    @echo "Makefile commands:"
		@echo "make"
	    @echo "make destroy"

.DEFAULT_GOAL := run

run:
	    @docker-compose up -d
destroy:
		@docker-compose down
dev:
		@docker-compose up -d
		@docker exec -it web /bin/sh
