let moreButton;

const addCompanyButton = Vue.createApp({
    data() {
        return {
            message: ""
        }
    },
    methods: {
        openAddCompanyModal() {
            document
                .getElementById("companies-modal")
                .classList
                .remove("invisible")
        }
    }
});

const addCompanyModal = Vue.createApp({
    data() {
        return {
            message: ""
        }
    },
    methods: {
        async addCompany() {

            let company = {
                name: document.getElementById("company-name").value,
                host: document.getElementById("company-host").value
            }

            document.getElementById("company-name").classList.remove("wrong-field")
            document.getElementById("company-host").classList.remove("wrong-field")

            if (company.name.trim() == "") {
                document.getElementById("company-name").classList.add("wrong-field")
            }
            if (company.host.trim() == "") {
                document.getElementById("company-host").classList.add("wrong-field")
            }

            let response = await fetchPostRequestWithJsonBody("/companies", company)

            if (response.error == null) {
                if (response.data.status == "error") {
                    this.message = response.data.message
                } else {
                    
                }
            }
        },

        close() {
            document
                .getElementById("companies-modal")
                .classList
                .add("invisible")
        }
    }
});


addCompanyModal.mount("#companyAddModal")
addCompanyButton.mount("#addNewCompany")