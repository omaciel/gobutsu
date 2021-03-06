CREATE TABLE IF NOT EXISTS testcases (
	id bigserial PRIMARY KEY,
	classname varchar(255) NOT NULL,
	filename varchar(255) NOT NULL,
	linenumber integer NOT NULL,
	testcasename varchar(255) NOT NULL,
	duration double precision NOT NULL
);