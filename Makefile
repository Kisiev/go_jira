build:
	git pull
	docker-compose down
	docker-compose build
	cd app && openssl req -newkey rsa:2048 -sha256 -nodes -keyout key.pem -x509 -days 365 -out cert.pem -subj "/C=RU/ST=New York/L=Brooklyn/O=Example Brooklyn Company/CN=62.113.98.178"
	#docker-compose up -d --build app