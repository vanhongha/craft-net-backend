-- Table: public.users

-- DROP TABLE IF EXISTS public.users;

CREATE TABLE IF NOT EXISTS public.users
(
    id integer NOT NULL,
    last_name "char" NOT NULL,
    fist_name "char" NOT NULL,
    date_of_birth date,
    email "char",
    phone_number "char",
    created_at time with time zone,
    updated_at time with time zone,
    status "char" NOT NULL DEFAULT '0'::"char",
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_fkey_accounts FOREIGN KEY (id)
        REFERENCES public.accounts (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to root;

COMMENT ON COLUMN public.users.id
    IS 'Unique identifier for the user.';

COMMENT ON COLUMN public.users.last_name
    IS 'The last name of the user.';

COMMENT ON COLUMN public.users.fist_name
    IS 'The first name of the user.';

COMMENT ON COLUMN public.users.date_of_birth
    IS 'The date of birth of the user.';

COMMENT ON COLUMN public.users.email
    IS 'The email address of the user.';

COMMENT ON COLUMN public.users.phone_number
    IS 'The phone number of the user.';

COMMENT ON COLUMN public.users.created_at
    IS 'The timestamp when the user record was created.';

COMMENT ON COLUMN public.users.updated_at
    IS 'The timestamp when the user record was last updated.';

COMMENT ON COLUMN public.users.status
    IS 'The current status of the user, 0 for inactive and other values for specific states.';