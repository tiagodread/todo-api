create table task (
	id SERIAL primary key,
	title varchar(255) not null,
	description text,
	created_at date not null,
	completed_at date,
	is_completed bool not null,
	reward_in_sats integer not null
);

