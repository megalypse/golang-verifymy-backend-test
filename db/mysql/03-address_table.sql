CREATE TABLE addresses (
  id BIGINT NOT NULL AUTO_INCREMENT,
  alias VARCHAR(255)
  zipcode VARCHAR(255),
  street_name VARCHAR(255),
  number VARCHAR(255),
  state VARCHAR(255),
  country VARCHAR(255),
  user_id BIGINT NOT NULL,

  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME,

  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);
