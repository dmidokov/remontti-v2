package database

var CreateUsersTable = `
	CREATE TABLE IF NOT EXISTS public.users(
    	id SERIAL PRIMARY KEY NOT NULL,
    	company_id integer NOT NULL,
		user_name character varying(20) COLLATE pg_catalog."default" NOT NULL,
		password text COLLATE pg_catalog."default" NOT NULL,
		last_login_date timestamp without time zone,
		last_login_error_date timestamp without time zone,
		CONSTRAINT "User_uniq" UNIQUE (user_name)
	)

	TABLESPACE pg_default;

	ALTER TABLE IF EXISTS public.users
		OWNER to %s;

	COMMENT ON TABLE public.users
		IS 'Contains information about users';`

var InsertAdminUser = `
	INSERT INTO public.users 
		(company_id, user_name, password) 
	VALUES 
		(0, 'admin', '%s');
`
var SelectAdminCount = `
	SELECT count(*) FROM public.users WHERE user_name='admin'; 
`
