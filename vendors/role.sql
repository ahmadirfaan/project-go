-- Pemilihan database untuk dimasukkan ke dalam role ---
use `backend-b-antar-jemput`;

INSERT INTO roles (created_at, updated_at, role) values (current_time, current_time, "Agent");
INSERT INTO roles (created_at, updated_at, role) values (current_time, current_time, "Customer");
