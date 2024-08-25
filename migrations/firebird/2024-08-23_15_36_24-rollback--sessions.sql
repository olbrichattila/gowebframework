EXECUTE BLOCK AS
    BEGIN
    IF (EXISTS (SELECT 1 FROM rdb$relations WHERE rdb$relation_name = 'sessions')) THEN
    BEGIN
        EXECUTE STATEMENT 'DROP TABLE "sessions"';
    END
END