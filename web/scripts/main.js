async function fetchPostRequestWithJsonBody(link, data) {
    let response = await fetch(link, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json;charset=utf-8'
        },
        body: JSON.stringify(data)
    });
    if (response.ok) {
        return {
            "data": await response.json(),
            "error": null
        }
    } else {
        return {
            "data": null,
            'error': {
                "statusText": response.statusText,
                "error": response.error,
                "status": response.status
            }
        }
    }
}
