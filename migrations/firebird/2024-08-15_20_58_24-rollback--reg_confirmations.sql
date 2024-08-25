EXECUTE BLOCK AS
    BEGIN
    IF (EXISTS (SELECT 1 FROM rdb$relations WHERE rdb$relation_name = 'reg_confirmations')) THEN
    BEGIN
        EXECUTE STATEMENT 'DROP TABLE "reg_confirmations"';
    END
END