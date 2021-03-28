create database job;
use job;
create table User(
   id VARCHAR(100) NOT NULL,
   email VARCHAR(100) NOT NULL,
   password VARCHAR(100) NOT NULL,
   role VARCHAR(100) NOT NULL,
   PRIMARY KEY ( id )
);