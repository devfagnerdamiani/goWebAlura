-- AULA 4 - Instalando Postgreesql 
-- Este arquivo é para ser inserido no POSTGREESQL

CREATE TABLE produtos(
id serial primary key,
nome varchar,
descricao varchar,
preco decimal,
quantidade integer
);


INSERT INTO produtos (nome,descricao,preco,quantidade) VALUES ('Camiseta','Preta',19,10);
INSERT INTO produtos (nome,descricao,preco,quantidade) VALUES ('Fone','Muito bom',99,5);