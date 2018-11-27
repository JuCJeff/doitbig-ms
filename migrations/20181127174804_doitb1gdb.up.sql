CREATE TABLE `doitb1gdb`.`users` (
  `id` INT NOT NULL,
  `firstname` VARCHAR(45) NULL,
  `lastname` VARCHAR(45) NULL,
  `email` VARCHAR(45) NULL,
  `password` VARCHAR(45) NULL,
  `image` VARCHAR(45) NULL,
  PRIMARY KEY (`id`));
  insert into users (id, firstname, lastname, email, password, image) values (12345, "John", "Appleseed"," japple@wisc.edu", "password", "asdf");
