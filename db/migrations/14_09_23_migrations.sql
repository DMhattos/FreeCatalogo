-- Tabela para armazenar informações sobre categorias de mídia
CREATE TABLE IF NOT EXISTS Categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL COMMENT 'Nome da categoria de mídia'
    INDEX idx_name (name)
);

