CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS albums (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	title VARCHAR(100),
	artist VARCHAR(100),
	price DECIMAL(10, 2)
);

INSERT INTO albums (id, title, artist, price)
VALUES
	(DEFAULT, 'Evergreen', 'After The Burial', 12.99),
	(DEFAULT, 'WLFGRL', 'Machine Girl', 10.99),
	(DEFAULT, 'Master of Puppets', 'Metallica', 11.99),
	(DEFAULT, 'Selected Ambient Works 85-92', 'Aphex Twin', 13.99),
	(DEFAULT, 'The Afterman: Descension', 'Coheed and Cambria', 9.99),
	(DEFAULT, 'SMILE! :D', 'Porter Robinson', 15.99);
