package database

// Таблицы

var CreateGroupsPermissionsTableSQL = `
	CREATE TABLE IF NOT EXISTS remonttiv2.group_permissions(
		permission_id SERIAL NOT NULL,
		component_id integer NOT NULL,
		component_type character varying(50) NOT NULL,
		group_id integer NOT NULL,
		actions integer NOT NULL,
		edit_time integer NOT NULL,
		PRIMARY KEY (permission_id)
	)
	    
	TABLESPACE pg_default;
`

var CreateUsersGroupsTableSQL = `
	CREATE TABLE IF NOT EXISTS remonttiv2.users_groups(
	    user_id integer NOT NULL,
	    group_id integer NOT NULL,
	    UNIQUE (user_id, group_id)                   
	)
	    
	TABLESPACE pg_default;
`

var CreateGroupsTableSQL = `
	CREATE TABLE IF NOT EXISTS remonttiv2.groups(
    	group_id SERIAL NOT NULL,
    	group_name character varying(50) NOT NULL,
		PRIMARY KEY (group_id),
		UNIQUE (group_name)	
	)

	TABLESPACE pg_default;`

var CreateUsersTableSQL = `
	CREATE TABLE IF NOT EXISTS remonttiv2.users(
    	id SERIAL NOT NULL,
    	company_id integer NOT NULL,
		user_name character varying(20) COLLATE pg_catalog."default" NOT NULL,
		password text COLLATE pg_catalog."default" NOT NULL,
		last_login_date integer NOT NULL,
		last_login_error_date integer NOT NULL,
		PRIMARY KEY (id, company_id, user_name),
		UNIQUE (company_id, user_name)	
	)

	TABLESPACE pg_default;`

var SetUserTableOwnerSql = `ALTER TABLE IF EXISTS remonttiv2.users OWNER to %s`

var SetUsersTableCommentSQL = `
	COMMENT ON TABLE remonttiv2.users
		IS 'Contains information about users';
`

var CreateNavigationTableSQL = `
	CREATE TABLE IF NOT EXISTS remonttiv2.navigation(
		id SERIAL PRIMARY KEY NOT NULL,
		item_type integer NOT NULL,
		link character varying(100) COLLATE pg_catalog."default" NOT NULL UNIQUE,
		label character varying(50) COLLATE pg_catalog."default" NOT NULL,
		ordinal_number integer NOT NULL,
		edit_time integer
	)

	TABLESPACE pg_default;
`

var CreateTranslationsTableSQL = `
	CREATE TABLE IF NOT EXISTS remonttiv2.translations(
		id SERIAL PRIMARY KEY NOT NULL,
		name character varying(30) COLLATE pg_catalog."default" NOT NULL,
		label character varying(30) COLLATE pg_catalog."default" NOT NULL,
		ru character varying(150) COLLATE pg_catalog."default",
		en character varying(150) COLLATE pg_catalog."default",
		edit_time integer
	)

	TABLESPACE pg_default;
`

var SetNavigationTableOwnerSQL = `ALTER TABLE IF EXISTS remonttiv2.navigation OWNER to %s;`

var SetNavigationTableCommentSQL = `COMMENT ON TABLE remonttiv2.navigation IS 'Contains information about navigation menu';`

var CreatePermissionsTableSQL = `
	CREATE TABLE IF NOT EXISTS remonttiv2.permissions(
		permission_id SERIAL NOT NULL,
		component_id integer NOT NULL,
		component_type character varying(50) NOT NULL,
		user_id integer NOT NULL,
		actions integer NOT NULL,
		edit_time integer NOT NULL,
		PRIMARY KEY (permission_id)
	)
	TABLESPACE pg_default;
`

var CreateCompaniesTableSQL = `
	CREATE TABLE IF NOT EXISTS remonttiv2.companies(
		company_id SERIAL NOT NULL,
		company_name character varying(50) NOT NULL,
		host_name character varying(100) NOT NULL,
		edit_time integer NOT NULL,
		PRIMARY KEY (company_id, company_name),
		UNIQUE (company_name, host_name)
	)
	TABLESPACE pg_default;
`

var SetPermissionsTableOwnerSQL = `ALTER TABLE IF EXISTS remonttiv2.permissions OWNER to %s;`

var SetPermissionsTableCommentSQL = `COMMENT ON TABLE remonttiv2.permissions IS 'Contains information about user to componens permissions';`
