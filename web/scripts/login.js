const loginButton = Vue.createApp({
    data() {
        return {
            counter: 0,
            message: '',
        };
    },
    methods: {
        async signIn(event) {
            let user = {
                login: document.getElementById('login').value,
                password: document.getElementById('password').value
            }

            let response = await fetchPostRequestWithJsonBody('/login', user)

            if (response.error == null) {

                if (response.data.status === "error") {
                    this.message = response.data.message
                } else {
                    window.location.href = window.location.href
                }
            }
        },
    },
});

loginButton.mount('#login-form')

