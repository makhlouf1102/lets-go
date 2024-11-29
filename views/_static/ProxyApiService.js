import { catchError } from "./catchError";
export class ApiService {

    prefix = "/api/v1"

    async login(email, password) {
        const formData = {
            email: email,
            password: password
        }

        const options = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        }

        const [error, data] = await catchError(fetch(`${this.prefix}/auth/login`, options))

        if (error) {
            throw Error("server error")
        }

        return data
    }
}