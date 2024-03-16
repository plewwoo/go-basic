CREATE TABLE book_golang.book(
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(200) NULL,
	`description` VARCHAR(200) NULL,
	`author` VARCHAR(200) NULL,
	`image` VARCHAR(200) NULL,
	`genre` VARCHAR(200) NULL,
	`public_date` VARCHAR(200) NULL,
	CONSTRAINT book_pk PRIMARY KEY(id);
)

ALTER TABLE book
ADD 'description' VARCHAR(200);
ADD 'author' VARCHAR(200);
ADD 'image' VARCHAR(200);
ADD 'genre' VARCHAR(200);
ADD 'public_date' VARCHAR(200);

INSERT INTO book (name, description, author, image, genre, public_date) VALUES ('Dragon Ball', 'The series follows the adventures of Goku, a young boy with a monkey tail and exceptional strength who battles evil-doers in hand-to-hand combat.', 'Akira Toriyama', 'https://www.norli.no/media/catalog/product/9/7/9781421555645_1_2.jpg?auto=webp&format=pjpg&width=728&height=910&fit=cover', 'Adventure', '1986');
INSERT INTO book (name, description, author, image, genre, public_date) VALUES ('Naruto', 'Naruto is a ninja-in-training whose wild antics amuse his teammates.', 'Masashi Kishimoto', 'https://m.media-amazon.com/images/I/91XCgBHoO8L._AC_UF1000,1000_QL80_.jpg', 'Adventure', '1999');
INSERT INTO book (name, description, author, image, genre, public_date) VALUES ('One Piece', 'ONE PIECE is a legendary high-seas quest unlike any other.', 'Eiichiro Oda', 'https://du.lnwfile.com/mm24ip.jpg', 'Adventure', '1997');
INSERT INTO book (name, description, author, image, genre, public_date) VALUES ('Bleach', 'Bleach (stylized in all caps) is a Japanese manga series', 'Tite Kubo', 'https://upload.wikimedia.org/wikipedia/en/3/3f/Bleach_%28manga%29_1.png', 'Shonen manga', '2002');