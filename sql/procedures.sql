DELIMITER //


DROP PROCEDURE IF EXISTS `theworlddata`.`resigter_account` //
CREATE PROCEDURE `theworlddata`.`resigter_account`
(
   param_email VARCHAR(50),
   param_password VARCHAR(32)
)
BEGIN
   INSERT INTO `theworlddata`.`account`
   (
      email,
      `password`
   )
   VALUES
   (
      param_email,
      MD5(param_password)
   );
END //


DROP PROCEDURE IF EXISTS `theworlddata`.`validate_account` //
CREATE PROCEDURE `theworlddata`.`validate_account`
(
   param_email VARCHAR(50),
   param_password VARCHAR(32)
)
BEGIN
   IF NOT EXISTS (
      SELECT 1
      FROM `theworlddata`.`account` a
      WHERE a.email = param_email AND
         a.password = MD5(param_password)
      LIMIT 1
   )
   THEN
      SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Login invalid. Try to again, please!', MYSQL_ERRNO = 1001;
   END IF;
END //