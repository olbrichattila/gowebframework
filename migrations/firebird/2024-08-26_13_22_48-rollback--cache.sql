EXECUTE BLOCK AS
    BEGIN
    IF (EXISTS (SELECT 1 FROM rdb$relations WHERE rdb$relation_name = 'caches')) THEN
    BEGIN
        EXECUTE STATEMENT 'DROP TABLE "caches"';
    END
END