INSERT INTO test.orders (id, chain_id, nonce, maker_asset, taker_asset, maker, rate) VALUES 
(1, '1', 0, '0xf4d2888d29d722226fafa5d9b24f9164c092421e', '0xf4d2888d29d722226fafa5d9b24f9164c092421f', '0xf4d2888d29d722226fafa5d9b24f9164c09eeeee', 0.1),
(2, '5', 1, '0xf4d2888d29d722226fafa5d9b24f9164c092421f', '0xf4d2888d29d722226fafa5d9b24f9164c092421e', '0xf4d2888d29d722226fafa5d9b24f9164c09eeeee', 0.2);

UPDATE test.orders set nonce = 12 where id = 1;

INSERT INTO test.orders (id, chain_id, nonce, maker_asset, taker_asset, maker, rate) VALUES 
(3, '137', 0, '0xf4d2888d29d722226fafa5d9b24f9164c092421e', '0xf4d2888d29d722226fafa5d9b24f9164c092421f', '0xf4d2888d29d722226fafa5d9b24f9164c09eeeee', 0.5),
(4, '5', 1, '0xf4d2888d29d722226fafa5d9b24f9164c092421f', '0xf4d2888d29d722226fafa5d9b24f9164c092421e', '0xf4d2888d29d722226fafa5d9b24f9164c09eeeee', 1);

UPDATE test.orders set chain_id = '4421' where id = 4;