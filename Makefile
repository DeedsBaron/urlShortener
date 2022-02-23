build:
	@echo "\033[0;32mBuilding binary...\033[m"
	@$(MAKE) -s -C urlShortener

all: build
	@$(MAKE) run -s -C urlShortener

inmem:
	@echo "\033[0;32mChanging SOLUTION value to 0 in .env file\033[m"
	sed -i -e 's/SOLUTION=0/SOLUTION=0/g' .env
	sed -i -e 's/SOLUTION=1/SOLUTION=0/g' .env
	docker-compose build --no-cache urlshortener
	docker-compose up -d --force-recreate urlshortener

psql: clean
	@echo "\033[0;32mChanging SOLUTION value to 1 in .env file\033[m"
	sed -i -e 's/SOLUTION=0/SOLUTION=1/g' .env
	sed -i -e 's/SOLUTION=1/SOLUTION=1/g' .env
	docker-compose build --no-cache
	docker-compose up -d --force-recreate

clean:
	- docker-compose down
	- docker volume rm $$(docker volume ls -q)

logs:
	docker-compose logs
status:
	docker ps -a

test_inmem:
	@$(MAKE) test_inmem -s -C urlShortener
test_psql: clean
	docker-compose build --no-cache postgresql
	docker-compose up -d --force-recreate postgresql
	@$(MAKE) test_psql -s -C urlShortener
testall: test_inmem test_psql

.PHONY: all build inmem psql clean re logs status test_inmem test_psql testall

.DEFAULT_GOAL := inmem
