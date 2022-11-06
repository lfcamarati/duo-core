CREATE TABLE client (
  id INT NOT NULL AUTO_INCREMENT,
  address VARCHAR(255) NOT NULL,
  email VARCHAR(100) NOT NULL,
  phone VARCHAR(20) NOT NULL,
  type VARCHAR(2) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE client_pf (
  id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  cpf VARCHAR(11) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (id) REFERENCES client(id)
);

CREATE TABLE client_pj (
  id INT NOT NULL,
  corporate_name VARCHAR(255) NOT NULL,
  cnpj VARCHAR(14) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (id) REFERENCES client(id)
);

CREATE TABLE service (
  id INT NOT NULL AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255),
  price DECIMAL(10,2) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE user (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (username)
);