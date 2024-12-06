let prefix = "/api/v1";
export const apiService = {

    async request(endpoint, options = {}) {
        const defaultHeaders = {
            "Content-Type": "application/json",
        };

        const token = localStorage.getItem("authToken");
        if (token) {
            defaultHeaders["Authorization"] = `Bearer ${token}`;
        }

        options.headers = { ...defaultHeaders, ...options.headers };

        return await fetch(`${prefix}${endpoint}`, options);
    },
    
    async ping() {
        return await this.request("/ping");
    },
    
    async pingProtected() {
        return await this.request("/ping-protected");
    },
    
    async login(email, password) {
        const formData = { email, password };

        const options = {
            method: "POST",
            body: JSON.stringify(formData),
        };

        return await this.request("/auth/login", options);
    }
}
