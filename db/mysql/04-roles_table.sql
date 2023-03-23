CREATE TABLE roles (
  id BIGINT NOT NULL AUTO_INCREMENT,
  alias VARCHAR(255) NOT NULL UNIQUE,

  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (id)
);
