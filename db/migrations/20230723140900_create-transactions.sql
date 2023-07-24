-- migrate:up
create table transactions (
  id BIGINT NOT NULL PRIMARY KEY,
  seller_name VARCHAR(200),
  ref_no VARCHAR(200),
  payment_method VARCHAR(100),
  total_price FLOAT,
  status SMALLINT DEFAULT 0,
  created_at TIMESTAMP WITHOUT TIME ZONE,
  updated_at TIMESTAMP WITHOUT TIME ZONE
);

-- migrate:down
drop table if exists transactions;
