start:
	docker-compose up -d
	docker-compose up -d --build app
build:
	docker-compose up -d --build app