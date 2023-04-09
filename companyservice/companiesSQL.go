package companyservice

const getAll = `SELECT * FROM remonttiv2.companies;`

const getCompanyByName = `SELECT * FROM remonttiv2.companies WHERE company_name=$1;`

const getCompanyById = `SELECT * FROM remonttiv2.companies WHERE company_id=$1;`

const getCompanyByHostName = `SELECT * FROM remonttiv2.companies WHERE host_name=$1;`

const insertCompany = "INSERT INTO remonttiv2.companies (company_name, host_name, edit_time) VALUES($1, $2, $3)"

const selectCompanyByNameHostTime = "SELECT * FROM remonttiv2.companies WHERE company_name=$1 AND host_name=$2 AND edit_time=$3"

const deleteByCompanyId = "DELETE FROM remonttiv2.companies WHERE company_id = $1"

const getCompaniesForUser = `SELECT 
    			remonttiv2.companies.company_id, remonttiv2.companies.company_name, 
    			remonttiv2.companies.host_name, remonttiv2.companies.edit_time 
			FROM 
			    remonttiv2.companies, remonttiv2.permissions 
			WHERE
			    remonttiv2.companies.company_id = remonttiv2.permissions.component_id AND
				remonttiv2.permissions.component_type = 'company' AND
			    (remonttiv2.permissions.actions & $1) = $1 AND
			    remonttiv2.permissions.user_id = $2;`

const getUserCompanyName = `SELECT
    			remonttiv2.companies.company_name
			FROM
    			remonttiv2.companies, remonttiv2.users
			WHERE
				remonttiv2.users.id = $1 AND
				remonttiv2.users.company_id = remonttiv2.companies.company_id
			`

const deletePermissionByComponentIdAndType = "DELETE FROM remonttiv2.permissions WHERE component_id = $1 AND component_type = $2"
