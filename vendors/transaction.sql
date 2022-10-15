-- Pemilihan database untuk dimasukkan ke dalam transaction --

use `agent-go`;

INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 1, 1,  "Cash-in & Out");
INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 2, 1,  "Report");
INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 3, 1,  "Setoran Uang");
INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 4, 1,  "Tarik Tunai");
INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 5, 1,  "Isi Ulang Pulsa");
INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 6, 1,  "Belanja Merchant");
INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 7, 2,  "Setoran Pinjaman");
INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 8, 2,  "Setoran Simpanan");
INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 9, 2,  "Tarik Tunai");
INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 10, 3,  "Registrasi Mobile Banking");
INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 11, 3,  "Registrasi Internet Banking");
INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 12, 3,  "Informasi Rekening");
INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 13, 3,  "Transfer Pembayaran");
INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 14, 3,  "Isi Ulang Pulsa");
INSERT INTO transaction_types (created_at, updated_at, id, service_type_transaction_id, name_type_transaction) values (current_time, current_time, 15, 3,  "Setor-Pasti");





