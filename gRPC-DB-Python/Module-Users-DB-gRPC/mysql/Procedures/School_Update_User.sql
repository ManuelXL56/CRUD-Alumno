CREATE PROCEDURE `Update_User`(IN setOldMatricule TEXT, IN setNewMatricule TEXT, IN setPassword TEXT, IN setFull_Name TEXT, IN setMail TEXT, IN SetDirection TEXT, IN SetPhone TEXT, IN SetRole TEXT)
BEGIN
	DECLARE _ID_Role INT DEFAULT 0;
    DECLARE _ID_DataUser INT DEFAULT 0;
    DECLARE _ID_Account INT DEFAULT 0;
    SET _ID_Role = (SELECT `ID_Role` FROM `Roles` WHERE `Role` IN (SetRole));
    IF ((SELECT count(`Matricule`) FROM `Accounts` WHERE `Matricule` IN (setOldMatricule)) = 1) AND (_ID_Role IS NOT NULL) THEN

        SET _ID_Account = (SELECT `ID_Account` FROM `Accounts` WHERE `Matricule` IN (setOldMatricule));
        UPDATE `Accounts` SET `Matricule` = setNewMatricule,  `Password` = setPassword,  `ID_Role` = _ID_Role
		WHERE `Matricule` IN (setOldMatricule);
        
		SET _ID_DataUser = (SELECT `ID_DataUser` FROM `Users` WHERE `ID_Account` IN (_ID_Account));
		UPDATE `DataUser` SET `Full_Name` = setFull_Name, `Mail` = setMail, `Direction`= SetDirection, `Phone` = SetPhone
        WHERE `ID_DataUser` IN (_ID_DataUser);
        
		UPDATE `Users` SET `ID_DataUser` = _ID_DataUser, `ID_Account` = _ID_Account WHERE `ID_DataUser` IN (_ID_DataUser);
        
        SELECT 'Update executed' AS `result`;
	ELSEIF (_ID_Role IS NULL) THEN
		SELECT 'Rol No valido' AS `result`;
	ELSE
		SELECT 'User not exist' AS `result`;
    END IF;
END;