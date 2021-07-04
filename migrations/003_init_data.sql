INSERT INTO inventories (sku, name, price, qty) VALUES ('120P90', 'Google Home', 49.99, 10);
INSERT INTO inventories (sku, name, price, qty) VALUES ('43N23P', 'Macbook Pro', 5399.99, 5);
INSERT INTO inventories (sku, name, price, qty) VALUES ('A304SD', 'Alexa Speaker', 109.5, 10);
INSERT INTO inventories (sku, name, price, qty) VALUES ('234234', 'Raspberry Pi B', 30, 2);


INSERT INTO promotions (item, promotion_type, minimum_qty, promotion_data, is_active, created_at, created_by, updated_at, updated_by) VALUES ('43N23P', 'free-item', 1, '234234', 1, DEFAULT, null, null, null);
INSERT INTO promotions (item, promotion_type, minimum_qty, promotion_data, is_active, created_at, created_by, updated_at, updated_by) VALUES ('120P90', 'buy-3-pay-2', 3, '2', 1, DEFAULT, null, null, null);
INSERT INTO promotions (item, promotion_type, minimum_qty, promotion_data, is_active, created_at, created_by, updated_at, updated_by) VALUES ('A304SD', 'discount', 3, '10', 1, DEFAULT, null, null, null);