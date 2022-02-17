build:
	@echo "\033[0;32mBuilding binary...\033[m"
<<<<<<< HEAD
	go build -o urlShortener/urlShortener -v ./urlShortener/cmd/urlShortener/main.go
=======
	@$(MAKE) -s -C urlShortener
>>>>>>> ed8f4a1 (postgresql container is configured and working)

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
<<<<<<< HEAD

=======
>>>>>>> ed8f4a1 (postgresql container is configured and working)
exec:
	docker exec -it postgresql bash
status:
	docker ps -a
<<<<<<< HEAD
=======
test:
	@$(MAKE) test -s -C urlShortener
conn:
	psql -h localhost -p 5432 -U deedsbaron urlshort
>>>>>>> ed8f4a1 (postgresql container is configured and working)

.PHONY: all lib clean fclean re

.DEFAULT_GOAL := build
