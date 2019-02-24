CREATE DATABASE content4godb;
CREATE USER 'content4gousr'@'localhost' IDENTIFIED BY 'content4gopasswd';
GRANT SELECT,INSERT,UPDATE,DELETE,CREATE,DROP ON content4godb.* TO 'content4gousr'@'localhost';