
CREATE TABLE IF NOT EXISTS test.orders
(
    id                      BIGINT          NOT NULL AUTO_INCREMENT,
    chain_id                VARCHAR(16)     NOT NULL,
    nonce                   BIGINT          NOT NULL,
    maker_asset             VARCHAR(64)     NOT NULL,
    taker_asset             VARCHAR(64)     NOT NULL,
    maker                   VARCHAR(64)     NOT NULL,
    rate                    DECIMAL(50, 10) NOT NULL,
    PRIMARY KEY (id)
);
