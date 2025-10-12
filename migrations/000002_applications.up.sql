CREATE TABLE applications (
    id uuid PRIMARY key default gen_random_uuid(),
    user_id uuid not null,
    company_name VARCHAR(255) NOT NULL,
    position_title VARCHAR(255) NOT NULL,
    job_url TEXT,
    salary_range VARCHAR(100), 
    location VARCHAR(255),
    status VARCHAR(50) DEFAULT 'Wishlist',
    notes TEXT, 
    applied_date TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    constraint fk_user
    	foreign key (user_id)
    	references users(id)
    	on delete cascade
);