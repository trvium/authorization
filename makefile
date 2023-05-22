.PHONY: run

run:
		docker-compose -f .trvium/docker-compose.yml up --build
