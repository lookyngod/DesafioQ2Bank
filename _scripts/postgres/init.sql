CREATE TABLE Usuarios (ID int Primary Key NOT NULL, 
Nome VARCHAR NOT NULL, Tipo TEXT CHECK (Tipo in('lojista','comum')), 
CPFCNPJ VARCHAR UNIQUE, 
Email VARCHAR NOT NULL, 
Senha VARCHAR NOT NULL, 
Saldo FLOAT NOT NULL);

CREATE TABLE IF NOT EXISTS Transacao (ID int Primary Key NOT NULL,
 IDUsuario int NOT NULL,
  IDDestino int NOT NULL,
   Valor FLOAT NOT NULL,
    Data TIMESTAMP NOT NULL);