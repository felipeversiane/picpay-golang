CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    amount NUMERIC(10, 2) NOT NULL,
    payee UUID NOT NULL,
    payer UUID NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    CONSTRAINT fk_payee FOREIGN KEY(payee) REFERENCES users(id),
    CONSTRAINT fk_payer FOREIGN KEY(payer) REFERENCES users(id)
);