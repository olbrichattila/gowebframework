EXECUTE BLOCK AS
    BEGIN
    IF (EXISTS (SELECT 1 FROM rdb$relations WHERE rdb$relation_name = 'jobs')) THEN
    BEGIN
        EXECUTE STATEMENT 'DROP TABLE "jobs"';
    END
END