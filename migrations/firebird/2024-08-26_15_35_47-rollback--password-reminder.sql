EXECUTE BLOCK AS
    BEGIN
    IF (EXISTS (SELECT 1 FROM rdb$relations WHERE rdb$relation_name = 'password_reminders')) THEN
    BEGIN
        EXECUTE STATEMENT 'DROP TABLE "password_reminders"';
    END
END