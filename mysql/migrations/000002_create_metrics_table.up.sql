CREATE TABLE IF NOT EXISTS metrics(
    id INT NOT NULL UNIQUE AUTO_INCREMENT,
    name VARCHAR (255) UNIQUE,
    high_treshold INT ,
    low_treshold INT ,
    current INT ,
    service_id INT ,
    FOREIGN KEY (service_id) REFERENCES services(service_id) ,
    PRIMARY KEY (id)
)