CREATE TABLE user (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (username)
);

CREATE TABLE client (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  cpfCnpj VARCHAR(14) NOT NULL,
  address VARCHAR(255),
  email VARCHAR(100),
  phone VARCHAR(20),
  type VARCHAR(2) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE service_client (
  id INT NOT NULL AUTO_INCREMENT,
  id_client INT NOT NULL,
  description VARCHAR(255) NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  period_type VARCHAR(15) NOT NULL,
  week_days VARCHAR(30),
  specific_date DATE,
  PRIMARY KEY (id),
  FOREIGN KEY (id_client) REFERENCES client(id),
  CONSTRAINT chk_period_type CHECK (period_type IN ('WEEKLY', 'SPECIFIC_DATE')),
  CONSTRAINT chk_week_days CHECK (
  	(period_type = 'WEEKLY' AND week_days IS NOT NULL) OR (period_type <> 'WEEKLY' AND week_days IS NULL)
	),
  CONSTRAINT chk_specific_date CHECK (
  	(period_type = 'SPECIFIC_DATE' AND specific_date IS NOT NULL) OR (period_type <> 'SPECIFIC_DATE' AND specific_date IS NULL)
	)
);

CREATE TABLE service (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  type VARCHAR(30) NOT NULL,
  description VARCHAR(255),
  price DECIMAL(10,2) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE social_media_management (
  id INT NOT NULL,
  week_frequency INT NOT NULL,
  plan_type VARCHAR(10) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (id) REFERENCES service(id)
);

-- ! REAVALIAR
-- CREATE TABLE contracted_service (
--   id INT NOT NULL AUTO_INCREMENT,
--   id_contract INT NOT NULL,
--   id_service INT NOT NULL,
--   title VARCHAR(255) NOT NULL,
--   description VARCHAR(255),
--   price DECIMAL(10,2) NOT NULL,
  
--   PRIMARY KEY (id),
--   FOREIGN KEY (id_contract) REFERENCES contract(id)
--   FOREIGN KEY (id_service) REFERENCES service(id)
-- );