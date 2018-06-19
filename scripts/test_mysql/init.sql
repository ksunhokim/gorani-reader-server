CREATE DATABASE gorani_test;
CREATE USER 'gorani'@'%' IDENTIFIED BY 'gorani';
GRANT ALL PRIVILEGES ON `gorani_test`.* TO 'gorani'@'%';
