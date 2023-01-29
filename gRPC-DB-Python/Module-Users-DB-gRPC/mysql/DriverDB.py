def configure_db(setConection):
    connection = setConection
    cursor = connection.cursor()

    sql = "DROP PROCEDURE IF EXISTS Delete_User"
    cursor.execute(sql)
    connection.commit()
    sql = "DROP PROCEDURE IF EXISTS Login"
    cursor.execute(sql)
    connection.commit()
    sql = "DROP PROCEDURE IF EXISTS Register_User"
    cursor.execute(sql)
    connection.commit()
    sql = "DROP PROCEDURE IF EXISTS Update_User"
    cursor.execute(sql)
    connection.commit()
    sql = "DROP PROCEDURE IF EXISTS Search_User"
    cursor.execute(sql)
    connection.commit()
    with open('/app/mysql/Procedures/School_Delete_User.sql', 'r') as f:
        sql = f.read()
    cursor.execute(sql)
    connection.commit()
    with open('/app/mysql/Procedures/School_Login.sql', 'r') as f:
        sql = f.read()
    cursor.execute(sql)
    connection.commit()
    with open('/app/mysql/Procedures/School_Register_User.sql', 'r') as f:
        sql = f.read()
    cursor.execute(sql)
    connection.commit()
    with open('/app/mysql/Procedures/School_Update_User.sql', 'r') as f:
        sql = f.read()
    cursor.execute(sql)
    connection.commit()
    with open('/app/mysql/Procedures/School_Search_User.sql', 'r') as f:
        sql = f.read()
    cursor.execute(sql)
    connection.commit()

    sql = """
        CREATE TABLE IF NOT EXISTS `DataUser` (
            `ID_DataUser` INT PRIMARY KEY AUTO_INCREMENT,
            `Full_Name` TEXT NOT NULL,
            `Mail` TEXT NOT NULL,
            `Direction` TEXT NOT NULL,
            `Phone` TEXT NOT NULL
        );
    """
    cursor.execute(sql)
    connection.commit()
    sql = """
        CREATE TABLE IF NOT EXISTS `Users` (
            `ID_User` INT PRIMARY KEY AUTO_INCREMENT,
            `ID_DataUser` INT NOT NULL,
            `ID_Account` INT NOT NULL
        );
    """
    cursor.execute(sql)
    connection.commit()
    sql = """
        CREATE TABLE IF NOT EXISTS `Accounts` (
            `ID_Account` INT PRIMARY KEY AUTO_INCREMENT,
            `Matricule` TEXT NOT NULL,
            `Password` TEXT NOT NULL,
            `ID_Role` INT NOT NULL
        );
    """
    cursor.execute(sql)
    connection.commit()
    sql = """
        CREATE TABLE IF NOT EXISTS `Roles` (
            `ID_Role` INT PRIMARY KEY AUTO_INCREMENT,
            `Role` TEXT NOT NULL
        );
    """
    cursor.execute(sql)
    connection.commit()

    sql = """
        ALTER TABLE `Users` ADD FOREIGN KEY (`ID_DataUser`) REFERENCES `DataUser` (`ID_DataUser`);
    """
    cursor.execute(sql)
    connection.commit()

    sql = """
        ALTER TABLE `Users` ADD FOREIGN KEY (`ID_Account`) REFERENCES `Accounts` (`ID_Account`);
    """
    cursor.execute(sql)
    connection.commit()

    sql = """
        ALTER TABLE `Accounts` ADD FOREIGN KEY (`ID_Role`) REFERENCES `Roles` (`ID_Role`);
    """
    cursor.execute(sql)
    connection.commit()

    sql = """
        INSERT INTO `Roles` (`Role`) SELECT * FROM (SELECT "pSyuPK-8EUNPAs3EfgFCvWLxrtI2" AS `Role`) AS temp WHERE NOT EXISTS (SELECT `Role` FROM `Roles` WHERE `Role` = "pSyuPK-8EUNPAs3EfgFCvWLxrtI2") LIMIT 1;
    """
    cursor.execute(sql)
    connection.commit()

    sql = """
        INSERT INTO `Roles` (`Role`) SELECT * FROM (SELECT "pSS2OK_oIxI0E4iakILWETrKxxdrRA==" AS `Role`) AS temp WHERE NOT EXISTS (SELECT `Role` FROM `Roles` WHERE `Role` = "pSS2OK_oIxI0E4iakILWETrKxxdrRA==") LIMIT 1;
    """
    cursor.execute(sql)
    connection.commit()

    sql = """
        call `Register_User`('liesITAHmOvq-urAS05q4M-tM-M=', '1XrwYfSxaLjfzTi5kQXw-zTqS3YuwRkY', '', '', '', '', 'pSyuPK-8EUNPAs3EfgFCvWLxrtI2');
    """
    cursor.execute(sql)
    connection.commit()

    connection.close()