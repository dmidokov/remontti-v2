const BASE_URL = import.meta.env.VITE_DEV_API_BASE_URL

/**
 * @param {String} link The date
 */
export async function get(link) {

    if (import.meta.env.DEV) {
        link = BASE_URL + link
    }

    let response = await fetch(link, {
        method: 'GET', headers: {
            'Content-Type': 'application/json;charset=utf-8'
        }
    });
    if (response.ok) {
        return {
            "data": await response.json(), "error": null
        }
    } else {
        return {
            "data": null, 'error': {
                "statusText": response.statusText, "error": response.error, "status": response.status
            }
        }
    }
}

/**
 * @param {String} link The URI
 * @param {Object} data The json request body
 */
export async function post(link, data) {

    if (import.meta.env.DEV) {
        link = BASE_URL + link
    }

    let response = await fetch(link, {
        method: 'POST', headers: {
            'Content-Type': 'application/json;charset=utf-8'
        }, body: JSON.stringify(data)
    });
    if (response.ok) {
        return {
            "data": await response.json(), "error": null
        }
    } else {
        return {
            "data": null,
            'error': {
                "statusText": response.statusText,
                "error": response.error,
                "status": response.status,
                "data": await response.json()
            }
        }
    }
}

/**
 * @param {String} link The URI
 * @param {Object} data The json request body
 */
export async function put(link, data) {

    if (import.meta.env.DEV) {
        link = BASE_URL + link
    }

    let response = await fetch(link, {
        method: 'POST', headers: {
            'Content-Type': 'application/json;charset=utf-8'
        }, body: JSON.stringify(data)
    });
    if (response.ok) {
        return {
            "data": await response.json(), "error": null
        }
    } else {
        return {
            "data": null,
            'error': {
                "statusText": response.statusText,
                "error": response.error,
                "status": response.status,
                "data": await response.json()
            }
        }
    }
}