CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS albums (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	title VARCHAR(100),
	artist VARCHAR(100),
	price DECIMAL(10, 2),
	image_url VARCHAR(255)
);

INSERT INTO albums (id, title, artist, price, image_url)
VALUES
	(DEFAULT, 'Evergreen', 'After The Burial', 12.99, 'https://go-learning-albums.s3.us-east-2.amazonaws.com/after-the-burial_evergreen.jpg'),
	(DEFAULT, 'WLFGRL', 'Machine Girl', 10.99, 'https://go-learning-albums.s3.us-east-2.amazonaws.com/machine-girl_wlfgrl.jfif'),
	(DEFAULT, 'Master of Puppets', 'Metallica', 11.99, 'https://go-learning-albums.s3.us-east-2.amazonaws.com/metallica_master-of-puppets.jpg'),
	(DEFAULT, 'The Black Album', 'Metallica', 11.99, 'https://go-learning-albums.s3.us-east-2.amazonaws.com/metallica_the-black-album.jpg'),
	(DEFAULT, 'Selected Ambient Works 85-92', 'Aphex Twin', 13.99, 'https://go-learning-albums.s3.us-east-2.amazonaws.com/aphex-twin_selected-ambient-works-85-92.png'),
	(DEFAULT, 'The Afterman: Descension', 'Coheed and Cambria', 9.99, 'https://go-learning-albums.s3.us-east-2.amazonaws.com/coheed-and-cambria_the-afterman-descension.jpg'),
	(DEFAULT, 'SMILE! :D', 'Porter Robinson', 15.99, 'https://go-learning-albums.s3.us-east-2.amazonaws.com/porter-robinson_smile!.jpg');
