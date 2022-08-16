async function fetchPostRequestWithJsonBody(data) {
    let response = await fetch('/login', {
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
