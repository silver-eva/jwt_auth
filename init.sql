create schema if not exists auth;

create extension if not exists "uuid-ossp";

create table if not exists auth.users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    uname VARCHAR(64) not null unique,
    name varchar(128),
    second varchar(128),
    password text not null
);

create table if not exists auth.sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES auth.users(id) ON DELETE CASCADE,
    created timestamp default now(),
    refreshed timestamp default now(),
    terminated timestamp,
    token text not null,
    active boolean default true
);

create table if not exists auth.properties (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES auth.users(id) ON DELETE CASCADE,
    max_sessions integer default 5,
    session_expires integer default 7 -- in days
);

create table if not exists auth.accesses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES auth.users(id) ON DELETE CASCADE,
    endpoint text not null
);