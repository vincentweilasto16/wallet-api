-- Remove the seeded users
DELETE FROM users
WHERE email IN (
        'andi.wijaya@example.com',
        'siti.rahma@example.com',
        'budi.santoso@example.com'
    );