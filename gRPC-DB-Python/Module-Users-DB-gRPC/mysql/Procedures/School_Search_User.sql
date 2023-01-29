CREATE PROCEDURE `Search_User`(IN setMatricule TEXT)
BEGIN
	SELECT `Matricule`, `Password`, `Full_Name`, `Mail`, `Direction`, `Phone`, `Role` FROM `Users`
	INNER JOIN `Accounts` ON `Users`.`ID_Account` = `Accounts`.`ID_Account`
	INNER JOIN `DataUser` ON `Users`.`ID_DataUser` = `DataUser`.`ID_DataUser`
	INNER JOIN `Roles` ON `Accounts`.`ID_Role` = `Roles`.`ID_Role` 
	WHERE `Matricule` IN (setMatricule);
END;