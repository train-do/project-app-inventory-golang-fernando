-- Insert data into Category
INSERT INTO "Category" (name, created_at) VALUES
('Electronics', now()),
('Furniture', now()),
('Clothing', now()),
('Food', now()),
('Toys', now()),
('Books', now()),
('Sports', now()),
('Health', now()),
('Beauty', now()),
('Automotive', now());

-- Insert data into Location
INSERT INTO "Location" (warehouse, rack, created_at) VALUES
('Warehouse A', 'Rack 1', now()),
('Warehouse A', 'Rack 2', now()),
('Warehouse B', 'Rack 1', now()),
('Warehouse B', 'Rack 2', now()),
('Warehouse C', 'Rack 1', now());

-- Insert data into Goods
INSERT INTO "Goods" (name, stock, created_at) VALUES
('Laptop', 10, now()),
('Sofa', 5, now()),
('T-Shirt', 20, now()),
('Pasta', 50, now()),
('Action Figure', 15, now()),
('Novel', 30, now()),
('Basketball', 12, now()),
('Vitamins', 25, now()),
('Lipstick', 40, now()),
('Car Battery', 7, now()),
('Headphones', 14, now()),
('Desk', 3, now()),
('Jeans', 18, now()),
('Cookies', 60, now()),
('Doll', 8, now()),
('Board Game', 22, now()),
('Smartphone', 25, now()),
('Coffee', 35, now()),
('Jacket', 10, now()),
('Bike', 4, now());

-- Insert data into Conjuction
INSERT INTO "Conjuction" (good_id, category_id, location_id) VALUES
(1, 1, 1),
(2, 2, 2),
(3, 3, 1),
(4, 4, 3),
(5, 5, 2),
(6, 6, 3),
(7, 7, 4),
(8, 8, 1),
(9, 9, 5),
(10, 10, 2),
(11, 1, 4),
(12, 2, 1),
(13, 3, 3),
(14, 4, 5),
(15, 5, 2),
(16, 1, 4),
(17, 6, 5),
(18, 7, 3),
(19, 8, 1),
(20, 9, 2);

-- Insert data into Log
INSERT INTO "Log" (good_id, information, qty, created_at) VALUES
(1, 'in', 10, now()),
(2, 'in', 5, now()),
(3, 'in', 20, now()),
(4, 'in', 50, now()),
(5, 'in', 15, now()),
(6, 'in', 30, now()),
(7, 'in', 12, now()),
(8, 'in', 25, now()),
(9, 'in', 40, now()),
(10, 'in', 7, now()),
(11, 'in', 14, now()),
(12, 'in', 3, now()),
(13, 'in', 18, now()),
(14, 'in', 60, now()),
(15, 'in', 8, now()),
(16, 'in', 22, now()),
(17, 'in', 25, now()),
(18, 'in', 35, now()),
(19, 'in', 10, now()),
(20, 'in', 4, now()),
(1, 'out', 3, now()),
(2, 'out', 1, now());
