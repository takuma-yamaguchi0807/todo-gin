-- seed users (password is bcrypt for "password")
INSERT INTO users (id, email, password_hash)
VALUES
  (
    '11111111-1111-1111-1111-111111111111',
    'test@example.com',
    crypt('Password1!', gen_salt('bf', 10))
  )
ON CONFLICT (id) DO NOTHING;

-- seed todos
INSERT INTO todos (id, user_id, title, description, status)
VALUES
  ('22222222-2222-2222-2222-222222222222', '11111111-1111-1111-1111-111111111111', '初期タスク1', 'セットアップを完了する', 'doing'),
  ('33333333-3333-3333-3333-333333333333', '11111111-1111-1111-1111-111111111111', '初期タスク2', 'READMEを読む', 'todo')
ON CONFLICT (id) DO NOTHING;

