CREATE TABLE users (
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null,
    password varchar(255) not null
);

CREATE TABLE user_sessions (
    id serial not null unique,
    user_id int not null,
    foreign key (user_id) references users (id),
    refresh_token varchar(255) not null,
    refresh_token_ttl int not null
);