this app is built with Go-Lang (clean architecture), React JS, postgres

- create database db_bengkel

CREATE TABLE histories
(
id SERIAL PRIMARY KEY,
no_rangka VARCHAR(32) NOT NULL,
merek VARCHAR(100) NOT NULL,
updated_at TIMESTAMP NOT NULL
)

CREATE TABLE customers
(
id SERIAL PRIMARY KEY,
name VARCHAR(100) NOT NULL,
phone VARCHAR(100) NOT NULL,
created_at timestamp NOT NULL
)

CREATE TABLE history_details
(
id SERIAL PRIMARY KEY,
history_id INT NOT NULL,
pic VARCHAR(100) NOT NULL,
created_at TIMESTAMP NOT NULL,
notes TEXT,
customer_id INT NOT NULL
)

ALTER TABLE history_details ADD plat_nomor VARCHAR(10) NOT NULL

- dependencies
  go get github.com/gofiber/fiber/v2
  go get -u github.com/doug-martin/goqu/v9
  go get github.com/lib/pq
  go get github.com/joho/godotenv
