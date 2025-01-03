Тестовое задание: сервис валютчика (Golang + MySQL)

Запуск MySQL-контейнера:
docker run --name mysql-local \
    -e MYSQL_ROOT_PASSWORD=root \
    -p 3306:3306 \
    -d mysql:8.0

Создаем базу:
docker exec -it mysql-local mysql -u root -p
CREATE DATABASE currency_db;

В коде (NewMySQLConnection) подставьте следующие параметры:
	•	host: localhost (или 127.0.0.1)
	•	port: 3306
	•	user: root
	•	password: root
	•	dbName: например, currency_db

Запускаем Go-приложение

Для получения данных:
    •	GET http://localhost:8080/rates - получить все записи
    •	GET http://localhost:8080/rates/day?date=YYYY-MM-DD - получить записи за конкретный день, где YYYY-MM-DD это дата.