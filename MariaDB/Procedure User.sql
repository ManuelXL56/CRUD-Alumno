USE School;

#----------------------------------------------------Procedure Login-------------------------------------------------------------------------------------------------------#
DROP PROCEDURE `Login`;
DELIMITER //
CREATE PROCEDURE `Login`(IN SetMatricule TEXT, IN SetPassword TEXT)
BEGIN
	DECLARE _ID_Role INT DEFAULT 0;
    DECLARE _Role TEXT DEFAULT "";
    
    SET _ID_Role = (SELECT `ID_Role` FROM `Accounts` WHERE `Matricule` IN (SetMatricule) AND `Password` IN (SetPassword));
    IF (_ID_Role > 0) THEN
    IF (_ID_Role > 0) THEN
		SET _Role = (SELECT `Role` FROM `Roles` WHERE `ID_Role` IN (_ID_Role));
        SELECT 'true' AS `result`, _Role AS `role`;
	ELSE
		SELECT 'false' AS `result`;
    END IF;
END
DELIMITER ;

call `Login`("liesITAHmOvq-urAS05q4M-tM-M=", "1XrwYfSxaLjfzTi5kQXw-zTqS3YuwRkY");

SELECT * FROM `Accounts` WHERE `Username` = 'Manu2001' AND `Password` = '123456';

SELECT * FROM `Users`;
SELECT * FROM `DataUser`;
SELECT * FROM `Accounts`;

#----------------------------------------------------Procedure Register User-------------------------------------------------------------------------------------------------------#
DROP PROCEDURE `Register_User`;
DELIMITER //
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
END//
DELIMITER ;

call `Register_User`('liesITAHmOvq-urAS05q4M-tM-M=', '1XrwYfSxaLjfzTi5kQXw-zTqS3YuwRkY', '', '', '', '', 'pSyuPK-8EUNPAs3EfgFCvWLxrtI2');

CALL `Register_User`("1njyZ_m1b7JmRysfXS5WPWderKbSP4cY-y0=", "pAr2dqLDZ-d3EiPOuZ2KCfICt3rVuQEvBqf2SQ==", "qSmtIKTrWrhH6cFXnBiKh3ifxdB1ew==", "qSmtIKTrH-U3WgiXT6Vjb-mYv7ZoB7iKTvI0XBs=", "gi-mMabjORhsbt7b7nMG930C5tp6jUA=", "0X36ZPGwa7BqRT6S_XExPCEGDqQnMHkSzbE=", "pSyuPK-8EUNPAs3EfgFCvWLxrtI2");

SELECT `Matricule`, `Password`, `Full_Name`, `Mail`, `Direction`, `Phone`, `Role` FROM `Users`
INNER JOIN `Accounts` ON `Users`.`ID_Account` = `Accounts`.`ID_Account`
INNER JOIN `DataUser` ON `Users`.`ID_DataUser` = `DataUser`.`ID_DataUser`
INNER JOIN `Roles` ON `Accounts`.`ID_Role` = `Roles`.`ID_Role`;

SELECT * FROM `Users`;
SELECT * FROM `DataUser`;
SELECT * FROM `Accounts`;

DELETE FROM `Users`;
DELETE FROM `DataUser`;
DELETE FROM `Accounts`;

#----------------------------------------------------Procedure Search User-------------------------------------------------------------------------------------------------------#

DROP PROCEDURE `Search_User`;
DELIMITER //
CREATE PROCEDURE `Search_User`(IN setMatricule TEXT)
BEGIN
	SELECT `Matricule`, `Password`, `Full_Name`, `Mail`, `Direction`, `Phone`, `Role` FROM `Users`
	INNER JOIN `Accounts` ON `Users`.`ID_Account` = `Accounts`.`ID_Account`
	INNER JOIN `DataUser` ON `Users`.`ID_DataUser` = `DataUser`.`ID_DataUser`
	INNER JOIN `Roles` ON `Accounts`.`ID_Role` = `Roles`.`ID_Role` 
	WHERE `Matricule` IN (setMatricule);
END//
DELIMITER ;

call `Search_User`('liesITAHmOvq-urAS05q4M-tM-M=');

#----------------------------------------------------Procedure Update User-------------------------------------------------------------------------------------------------------#
DROP PROCEDURE `Update_User`;
DELIMITER //
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
END//
DELIMITER ;

call `Update_User`('liesITAHmOvq-urAS05q4M-tM-M=', '1XrwYfSxaLjfzTi5kQXw-zTqS3YuwRkY', 'lDq2MKPmbi9c5AEV2YEtximQYiLLkg0=', 'qSmtIKTrf8cyBgiRQ8wdmv3-t_J0iCsLjtEi6Mkc6T6Q0fPw', 
'iSmtLJ7qOOIMRlq4SoMrm_35uLOt1mXatSqiyGKipGHWJyTdZcaA', 'piewJLTiLKA3EUu5ToU6hQuNKMoummOm2xMc8HPcTMs=', '0X36ZPGwa7BqRT6S_XExPCEGDqQnMHkSzbE=', 
'Admin');

SELECT `Matricule`, `Password`, `Full_Name`, `Mail`, `Direction`, `Phone`, `Role` FROM `Users`
INNER JOIN `Accounts` ON `Users`.`ID_Account` = `Accounts`.`ID_Account`
INNER JOIN `DataUser` ON `Users`.`ID_DataUser` = `DataUser`.`ID_DataUser`
INNER JOIN `Roles` ON `Accounts`.`ID_Role` = `Roles`.`ID_Role`;

#----------------------------------------------------Procedure Delete User-------------------------------------------------------------------------------------------------------#
DELETE FROM `Users`;
DELETE FROM `DataUser`;
DELETE FROM `Accounts`;

DROP PROCEDURE `Delete_User`;
DELIMITER //
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
END//
DELIMITER ;

call `Delete_User`('liesITAHmOvq-urAS05q4M-tM-M=');