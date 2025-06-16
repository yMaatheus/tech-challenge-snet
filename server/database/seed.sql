INSERT INTO establishments (number, name, corporate_name, address, city, state, zip_code, address_number)
VALUES
  ('E001', 'Establishment One', 'Establishment One Ltd', '123 Main St', 'CityA', 'CA', '12345-678', '10'),
  ('E002', 'Establishment Two', 'Establishment Two Inc', '456 Side Ave', 'CityB', 'CB', '23456-789', '20');

INSERT INTO stores (number, name, corporate_name, address, city, state, zip_code, address_number, establishment_id)
VALUES
  ('S001', 'Store Alpha', 'Alpha Ltd', '111 Store Rd', 'CityA', 'CA', '12345-000', '1', 1),
  ('S002', 'Store Beta', 'Beta Inc', '222 Store Ave', 'CityA', 'CA', '12345-111', '2', 1),
  ('S003', 'Store Gamma', 'Gamma LLC', '333 Shop St', 'CityB', 'CB', '23456-000', '3', 2);
