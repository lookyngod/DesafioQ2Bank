
CREATE TABLE Usuarios (ID SERIAL PRIMARY KEY UNIQUE, 
Nome VARCHAR NOT NULL,
Tipo TEXT CHECK (Tipo in('lojista','comum')), 
CPFCNPJ VARCHAR UNIQUE, 
Email VARCHAR NOT NULL, 
Senha VARCHAR NOT NULL, 
Saldo FLOAT NOT NULL);

CREATE TABLE Transacao (ID SERIAL PRIMARY KEY UNIQUE,
 IDOrigem int NOT NULL,
  IDDestino int NOT NULL,
   Valor FLOAT NOT NULL,
    Data Date NOT NULL);

ALTER TABLE Transacao 
ADD FOREIGN KEY (IDOrigem) REFERENCES Usuarios(ID);

ALTER TABLE Transacao 
ADD FOREIGN KEY (IDDestino) REFERENCES Usuarios(ID);

INSERT INTO Usuarios (Nome, Tipo, CPFCNPJ, Email, Senha, Saldo) VALUES ('ANDERSON', 'comum', '123.123.123-90','anderson@gmail.com', '123456', 200.10);
INSERT INTO Usuarios (Nome, Tipo, CPFCNPJ, Email, Senha, Saldo) VALUES ('MARIA', 'comum', '444.123.123-90', 'maria@gmail.com', '444444', 800.10);
INSERT INTO Usuarios (Nome, Tipo, CPFCNPJ, Email, Senha, Saldo) VALUES ('Q2Bank', 'lojista', '47.263.110/0001-00', 'q2bank@q2bank.com', '555555', 20.00);
INSERT INTO Usuarios (Nome, Tipo, CPFCNPJ, Email, Senha, Saldo) VALUES ('Q2Pay', 'lojista', '55.451.532/0001-37', 'q2pay@q2pay.com', '666666', 20.00);