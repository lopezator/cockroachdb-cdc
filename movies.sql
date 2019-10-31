/* Use database test; */
USE test;

/* Drop table movies if exist, for a clean start */
DROP TABLE IF EXISTS movies;

/* Create table movies */
CREATE TABLE movies (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    title_id STRING(40) NOT NULL,
    ordering INT4 NOT NULL,
    title STRING NOT NULL,
    region STRING(40) NULL,
    language STRING(40) NULL,
    types STRING(200) NULL,
    attributes STRING(200) NULL,
    is_original_title BOOL NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    FAMILY "primary" (id, title_id, ordering, title, region, language, types, attributes, is_original_title)
);

/* Add a fixed movie: */
INSERT INTO movies(id, title_id, "ordering", title, region, "language", types, "attributes", is_original_title)
    VALUES(gen_random_uuid(), 'tt1337', 0, 'Santa with muscles', 'US', 'en', 'original', 'literal English title', true);

/*
/* Add other movies (random): */
INSERT INTO movies(id, title_id, "ordering", title, region, "language", types, "attributes", is_original_title)
    VALUES(gen_random_uuid(), concat('tt', gen_random_uuid()::string), 0, 'Automatically generated title', 'Some region', 'en', 'random', 'auto', false);

/* Update a movie: */
UPDATE movies SET title = 'Santa with muscles - by Hulk Hogan'
    WHERE title_id = 'tt1337';

/* Delete a movie */
DELETE FROM movies WHERE title_id = 'tt1337';
*/