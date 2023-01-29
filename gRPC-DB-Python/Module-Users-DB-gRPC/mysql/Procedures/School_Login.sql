CREATE PROCEDURE `Login`(IN SetMatricule TEXT, IN SetPassword TEXT)
BEGIN
	DECLARE _ID_Role INT DEFAULT 0;
    DECLARE _Role TEXT DEFAULT "";
    
    SET _ID_Role = (SELECT `ID_Role` FROM `Accounts` WHERE `Matricule` IN (SetMatricule) AND `Password` IN (SetPassword));
    IF (_ID_Role > 0) THEN
		SET _Role = (SELECT `Role` FROM `Roles` WHERE `ID_Role` IN (_ID_Role));
        SELECT 'true' AS `result`, _Role AS `role`;
	ELSE
		SELECT 'false' AS `result`;
    END IF;
END;