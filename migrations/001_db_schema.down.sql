-- Down Migration (Rollback)
DROP TABLE IF EXISTS subscriptions;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS payment_methods;
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS merchants;