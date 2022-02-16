build:
	@echo "\033[0;32mBuilding binary...\033[m"
	go build -o urlShortener/urlShortener -v ./urlShortener/cmd/urlShortener/main.go

dir4db:
	@echo "\033[0;32mCreating folder for database volume at $${HOME}/db-data...\033[m"
	@if [ ! -d "$${HOME}/db-data" ]; then mkdir $${HOME}/db-data; fi

run: dir4db
	docker-compose up -d --force-recreate
logs:
	docker-compose logs

clean:
	docker-compose down
	docker volume rm $$(docker volume ls -q)
	docker rmi -f $$(docker images -aq)

exec:
	docker exec -it postgresql bash
status:
	docker ps -a

.PHONY: all lib clean fclean re

.DEFAULT_GOAL := build
