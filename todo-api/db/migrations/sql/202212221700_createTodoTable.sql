CREATE TABLE IF NOT EXISTS todos(
  id uuid DEFAULT gen_random_uuid(),
  title VARCHAR NOT NULL,
  created_at TIMESTAMP NOT NULL,
  PRIMARY KEY (id)
)
