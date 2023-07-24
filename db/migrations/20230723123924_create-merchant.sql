-- migrate:up
create table merchants (
  id BIGINT NOT NULL PRIMARY KEY,
  name VARCHAR(200),
  store_name VARCHAR(200),
  status SMALLINT DEFAULT 0,
  phone_no VARCHAR(150),
  joined_at TIMESTAMP WITHOUT TIME ZONE,
  created_at TIMESTAMP WITHOUT TIME ZONE,
  updated_at TIMESTAMP WITHOUT TIME ZONE
);

-- migrate:down
drop table if exists merchants;
