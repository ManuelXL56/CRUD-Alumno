CREATE PROCEDURE `Delete_User`(IN setMatricule TEXT)
BEGIN
    DECLARE _ID_DataUser INT DEFAULT 0;
    DECLARE _ID_Account INT DEFAULT 0;
    IF ((SELECT count(`Matricule`) FROM `Accounts` WHERE `Matricule` IN (setMatricule)) > 0) THEN
        SET _ID_Account = (SELECT `ID_Account` FROM `Accounts` WHERE `Matricule` IN (setMatricule));
        SET _ID_DataUser = (SELECT `ID_DataUser` FROM `Users` WHERE `ID_Account` IN (_ID_Account));
        DELETE FROM `Users` WHERE `ID_Account` IN (_ID_Account) AND `ID_DataUser` IN (_ID_DataUser);
        DELETE FROM `DataUser` WHERE `ID_DataUser` IN (_ID_DataUser);
        DELETE FROM `Accounts` WHERE `Matricule` IN (setMatricule);
        SELECT 'Delete executed' AS `result`;
	ELSE
		SELECT 'User not exist' AS `result`;
    END IF;
END;