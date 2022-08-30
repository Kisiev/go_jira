start:
	docker-compose up -d
	docker-compose up -d --no-deps --build app
build:
	docker-compose up -d --no-deps --build app