CREATE TABLE book_golang.book(
     `id` INT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(100) NULL,
	CONSTRAINT book_pk PRIMARY KEY(id)
)

ALTER TABLE book
ADD 'description' VARCHAR(200);
ADD 'author' VARCHAR(200);
ADD 'image' VARCHAR(200);
ADD 'genre' VARCHAR(200);
ADD 'public_date' VARCHAR(200);