build:
	@echo "\033[0;32mBuilding binary...\033[m"
	@$(MAKE) -s -C urlShortener

all: build
	@$(MAKE) run -s -C urlShortener
inmem:
	sed -i -e 's/SOLUTION=0/SOLUTION=0/g' .env
	sed -i -e 's/SOLUTION=1/SOLUTION=0/g' .env
	docker-compose build --no-cache
	docker-compose up -d --force-recreate
psql:
	sed -i -e 's/SOLUTION=0/SOLUTION=1/g' .env
	sed -i -e 's/SOLUTION=1/SOLUTION=1/g' .env
	docker-compose build --no-cache
	docker-compose up -d --force-recreate

run_inmem:
	@$(MAKE) run_inmem -s -C urlShortener

dir4db:
	@echo "\033[0;32mCreating folder for database volume at $${HOME}/db-data...\033[m"
	@if [ ! -d "$${HOME}/db-data" ]; then mkdir $${HOME}/db-data; fi

run: dir4db
	docker-compose up -d --build
logs:
	docker-compose logs
vol :

clean:
	docker-compose down
	docker volume rm $$(docker volume ls -q)
exec:
	docker exec -it urlshortener bash
url:
	docker exec -it urlshortener bash

status:
	docker ps -a
test:
	@$(MAKE) test -s -C urlShortener
conn:
	psql -h localhost -p 5432 -U deedsbaron urlshort
.PHONY: all lib clean fclean re

.DEFAULT_GOAL := all
