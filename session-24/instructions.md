docker run -d --name docker-mysql -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 mysql:latest

docker run -d --name docker-mysql -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 mysql:latest

docker pull mysql:latest

show databases;
create database users;

show tables;

CREATE TABLE Users (FirstName VARCHAR(20),SecondName VARCHAR(20),UserName Varchar(40) ,password Varchar(40));
