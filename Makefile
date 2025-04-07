build:
	sudo docker-compose build

up-app:
	sudo docker-compose up -d app

up: up-app

down:
	sudo docker-compose down

restart: down up

app-shell:
	sudo docker-compose exec app sh

