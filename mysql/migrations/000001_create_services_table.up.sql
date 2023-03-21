CREATE TABLE IF NOT EXISTS services(
    service_id INT NOT NULL UNIQUE AUTO_INCREMENT,
    user_id VARCHAR (255) UNIQUE,
    name VARCHAR (255),
    region VARCHAR (255),
    PRIMARY KEY (service_id)
)