CREATE TABLE IF NOT EXISTS posts(
    id SERIAL PRIMARY KEY, -- Уникальный идентификатор.
    title VARCHAR(255) NOT NULL, -- Заголовок.
    description TEXT NOT NULL, -- Описание.
    created_at TIMESTAMP NOT NULL DEFAULT NOW(), -- Дата создания записи.
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() -- Дата обновления записи.
);