create table if not exists users (
    id uuid primary key default gen_random_uuid(),
    email varchar(255) not null,
    first_name varchar(50),
    last_name varchar(50),
    password_hash varchar(255) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz,
    deleted_at timestamptz
);

create unique index if not exists idx_users_email_lower on users(lower(email)) where deleted_at is null;


create index if not exists idx_users_deleted_at on users(deleted_at) where deleted_at is null;

