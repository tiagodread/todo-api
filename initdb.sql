CREATE SCHEMA IF NOT EXISTS public;


CREATE TABLE IF NOT EXISTS public.task (
	id SERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	description text,
	created_at DATE NOT NULL,
	completed_at DATE,
	is_completed bool NOT NULL,
	reward_in_sats integer NOT NULL,
	due_date DATE
);