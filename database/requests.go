package database

// Таблицы

var CreateUsersTableSQL = `
	CREATE TABLE IF NOT EXISTS public.users(
    	id SERIAL PRIMARY KEY NOT NULL,
    	company_id integer NOT NULL,
		user_name character varying(20) COLLATE pg_catalog."default" NOT NULL,
		password text COLLATE pg_catalog."default" NOT NULL,
		last_login_date integer NOT NULL,
		last_login_error_date integer NOT NULL,
		CONSTRAINT "User_uniq" UNIQUE (user_name)
	)

	TABLESPACE pg_default;`

var SetUserTableOwnerSql = `ALTER TABLE IF EXISTS public.users OWNER to %s`

var SetUsersTableCommentSQL = `
	COMMENT ON TABLE public.users
		IS 'Contains information about users';`

var CreateNavigationTableSQL = `
	CREATE TABLE IF NOT EXISTS public.navigation(
		id SERIAL PRIMARY KEY NOT NULL,
		item_type integer NOT NULL,
		link character varying(100) COLLATE pg_catalog."default" NOT NULL UNIQUE,
		label character varying(50) COLLATE pg_catalog."default" NOT NULL,
		edit_time integer
	)

	TABLESPACE pg_default;
`

var CreateTranslationsTableSQL = `
	CREATE TABLE IF NOT EXISTS public.translations(
		id SERIAL PRIMARY KEY NOT NULL,
		name character varying(30) COLLATE pg_catalog."default" NOT NULL,
		label character varying(30) COLLATE pg_catalog."default" NOT NULL,
		ru character varying(150) COLLATE pg_catalog."default",
		en character varying(150) COLLATE pg_catalog."default",
		edit_time integer
	)

	TABLESPACE pg_default;
`

var SetNavigationTableOwnerSQL = `ALTER TABLE IF EXISTS public.navigation OWNER to %s;`

var SetNavigationTableCommentSQL = `COMMENT ON TABLE public.navigation IS 'Contains information about navigation menu';`

var CreatePermissionsTableSQL = `
	CREATE TABLE IF NOT EXISTS public.permissions(
		permission_id SERIAL NOT NULL,
		component_id integer NOT NULL,
		user_id integer NOT NULL,
		actions integer NOT NULL,
		edit_time integer NOT NULL
	)
	TABLESPACE pg_default;
`
var SetPermissionsTableOwnerSQL = `ALTER TABLE IF EXISTS public.permissions OWNER to %s;`

var SetPermissionsTableCommentSQL = `COMMENT ON TABLE public.permissions IS 'Contains information about user to componens permissions';`
