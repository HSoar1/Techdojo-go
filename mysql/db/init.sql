CREATE DATABASE api_database;
use api_database;

CREATE TABLE users (
    id integer(10), 
    username varchar(255), 
    token varchar(255)); 


INSERT INTO users (id, username, token) VALUES ( 1, 'soar' , 'soarstoken');
INSERT INTO users (id, username, token) VALUES ( 2, 'shizuku' , 'shizukustoken');