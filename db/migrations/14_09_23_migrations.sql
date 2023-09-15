-- Criação da tabela para armazenar informações sobre categorias de mídia
CREATE TABLE IF NOT EXISTS Categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

-- Adição de um índice à coluna "name" para melhorar a performance em consultas de busca
CREATE INDEX IF NOT EXISTS idx_name ON Categories (name);



