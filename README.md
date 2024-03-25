
1.  Buat database postgresql
2.  Masuk ke database yang sudah di buat dan lakukan query ini:
    BEGIN;
    CREATE TYPE role_type AS ENUM(
      'admin',
      'user'
    );
    COMMIT;
