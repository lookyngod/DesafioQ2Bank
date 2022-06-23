
CREATE TABLE Usuarios (ID INT GENERATED ALWAYS AS IDENTITY UNIQUE, 
Nome VARCHAR NOT NULL,
Tipo TEXT CHECK (Tipo in('lojista','comum')), 
CPFCNPJ VARCHAR UNIQUE, 
Email VARCHAR NOT NULL, 
Senha VARCHAR NOT NULL, 
Saldo FLOAT NOT NULL);

CREATE TABLE Transacao (ID INT GENERATED ALWAYS AS IDENTITY,
 IDOrigem int NOT NULL,
  IDDestino int NOT NULL,
   Valor FLOAT NOT NULL,
    Data TIMESTAMP NOT NULL);

ALTER TABLE Transacao 
ADD FOREIGN KEY (IDOrigem) REFERENCES Usuarios(ID);

ALTER TABLE Transacao 
ADD FOREIGN KEY (IDDestino) REFERENCES Usuarios(ID);

INSERT INTO Usuarios (Nome, Tipo, CPFCNPJ, Email, Senha, Saldo) VALUES ('ANDERSON', 'comum', '123.123.123-90','anderson@gmail.com', '123456', 200.10);
INSERT INTO Usuarios (Nome, Tipo, CPFCNPJ, Email, Senha, Saldo) VALUES ('MARIA', 'comum', '444.123.123-90', 'maria@gmail.com', '444444', 800.10);