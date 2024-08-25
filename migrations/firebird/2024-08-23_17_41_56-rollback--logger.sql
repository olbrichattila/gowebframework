EXECUTE BLOCK AS
    BEGIN
    IF (EXISTS (SELECT 1 FROM rdb$relations WHERE rdb$relation_name = 'logs')) THEN
    BEGIN
        EXECUTE STATEMENT 'DROP TABLE "logs"';
    END
END