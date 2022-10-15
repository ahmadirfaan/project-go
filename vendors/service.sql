-- Pemilihan database untuk dimasukkan ke dalam service --

use `agent-go`;

INSERT INTO service_type_transactions (created_at, updated_at, id, name_service) values (current_time, current_time, 1,  "Laku Pandai");
INSERT INTO service_type_transactions (created_at, updated_at, id,  name_service) values (current_time, current_time, 2, "Tunai");
INSERT INTO service_type_transactions (created_at, updated_at, id, name_service) values (current_time, current_time, 3, "Mini ATM BRI");

