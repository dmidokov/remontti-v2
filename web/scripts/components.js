var loginButton = Vue.createApp({
    data() {
        return {
            counter: 0,
        };
    },
    methods: {
        async signIn(event) {
            let user = {
                login: document.getElementById('login').value,
                password: document.getElementById('password').value
            }

            let response = await fetchPostRequestWithJsonBody(user)

            if (response.error == null) {
                console.log("Response data:")
                console.log(response.data)
            }
        },
    },
});
