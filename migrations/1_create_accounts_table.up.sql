-- Table: public.accounts

-- DROP TABLE IF EXISTS public.accounts;

CREATE TABLE IF NOT EXISTS public.accounts
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    user_id integer NOT NULL,
    username "char" NOT NULL,
    password_hash "char" NOT NULL,
    created_at time with time zone,
    status "char" NOT NULL DEFAULT '0'::"char",
    CONSTRAINT accounts_pkey PRIMARY KEY (id),
    CONSTRAINT user_id UNIQUE (user_id),
    CONSTRAINT username UNIQUE (username)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.accounts
    OWNER to root;

COMMENT ON COLUMN public.accounts.id
    IS 'Unique identifier for the account, auto-generated.';

COMMENT ON COLUMN public.accounts.user_id
    IS 'Links the account to a specific user.';

COMMENT ON COLUMN public.accounts.username
    IS 'User-selected name for logging in.';

COMMENT ON COLUMN public.accounts.password_hash
    IS 'Hashed password for security.';

COMMENT ON COLUMN public.accounts.created_at
    IS 'Timestamp of account creation.';

COMMENT ON COLUMN public.accounts.status
    IS 'Account status.';