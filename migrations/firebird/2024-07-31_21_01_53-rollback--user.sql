EXECUTE BLOCK AS
    BEGIN
    IF (EXISTS (SELECT 1 FROM rdb$relations WHERE rdb$relation_name = 'users')) THEN
    BEGIN
        EXECUTE STATEMENT 'DROP TABLE "users"';
    END
END