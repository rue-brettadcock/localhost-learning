CREATE DATABASE userInformation;
use userInformation;

CREATE TABLE users (
    id int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username varchar(50) DEFAULT NULL,
    password varchar(120) DEFAULT NULL
)


