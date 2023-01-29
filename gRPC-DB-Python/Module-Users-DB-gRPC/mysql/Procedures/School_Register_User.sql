CREATE PROCEDURE `Register_User`(IN setMatricule TEXT, IN setPassword TEXT, IN setFull_Name TEXT, IN setMail TEXT, IN SetDirection TEXT, IN SetPhone TEXT, IN SetRole TEXT)
BEGIN
	DECLARE _ID_Role INT DEFAULT 0;
    DECLARE _ID_DataUser INT DEFAULT 0;
    DECLARE _ID_Account INT DEFAULT 0;
    SET _ID_Role = (SELECT `ID_Role` FROM `Roles` WHERE `Role` = SetRole);
    IF ((SELECT count(`Matricule`) FROM `Accounts` WHERE `Matricule` IN (SetMatricule)) = 0)  AND (_ID_Role IS NOT NULL) THEN
		ALTER TABLE `DataUser` AUTO_INCREMENT = 1;
        ALTER TABLE `Accounts` AUTO_INCREMENT = 1;
        ALTER TABLE `Users` AUTO_INCREMENT = 1;
        
        INSERT INTO `Accounts` (`Matricule`, `Password`, `ID_Role`) 
		VALUES (setMatricule, setPassword, _ID_Role);
        
        INSERT INTO `DataUser` (`Full_Name`, `Mail`, `Direction`, `Phone`) 
		VALUES (setFull_Name, setMail, SetDirection, SetPhone);
        
        SET _ID_DataUser = (SELECT LAST_INSERT_ID());
        SET _ID_Account = (SELECT `ID_Account` FROM `Accounts` WHERE `Matricule` IN (setMatricule) AND `Password` IN (SetPassword));
        INSERT INTO `Users` (`ID_DataUser`, `ID_Account`) 
		VALUES (_ID_DataUser, _ID_Account);
        
        SELECT 'Register executed' AS `result`;
	ELSEIF (_ID_Role IS NULL) THEN
		SELECT 'Rol No valido' AS `result`;
	ELSE
		SELECT 'User already exists' AS `result`;
    END IF;
END;