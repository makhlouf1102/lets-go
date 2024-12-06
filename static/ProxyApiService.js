import { catchError } from './catchError.js';
import { JSON_ERROR, SERVER_ERROR, BAD_CREDENTIALS_ERROR, INVALID_AUTH_ERROR, NEW_TOKEN_ERROR } from './errorsConst.js';
import { apiService } from './ApiService.js';


export const proxyApiService = {
    async protectedRoute(request) {
        if (typeof request !== "function") {
            throw new Error("The request parameter must be a function.");
        }

        let [error, response] = await catchError(request());

        if (error) throw Error(SERVER_ERROR)

        if (!response.ok && (response.status == "403" || response.status == "401")) {
            window.sessionStorage.removeItem("accessToken")
            throw Error(INVALID_AUTH_ERROR)
        } else {
            let [errorJson, responseJson] = await response.json()

            if (errorJson) throw Error(JSON_ERROR)

            if (responseJson.status == "New Token") {
                window.sessionStorage.setItem('accessToken', data.accessToken);

                return await catchError(request())

            }
        }

        return response
    },

    async login(email, password) {
        let [error, response] = await catchError(apiService.login(email, password))

        if (error) throw Error(SERVER_ERROR)

        if (!response.ok) {
            const responseStatus = response.status

            if (responseStatus == "400" || responseStatus == "500")
                throw Error(SERVER_ERROR)

            if (responseStatus == "401") throw Error()
            throw Error(BAD_CREDENTIALS_ERROR)
        }

        let [errorJson, responseJson] = await catchError(response.json())

        if (errorJson) throw Error(JSON_ERROR)

        return responseJson
    },

    async ping() {
        let [error, response] = await catchError(apiService.ping())

        if (error) throw Error(SERVER_ERROR)

        if (!response.ok) {
            throw Error(response.status)
        }

        return true
    },

    async pingProtected() {
        let [error, response] = await (this.protectedRoute(() => apiService.pingProtected()))

        if (error) throw Error(SERVER_ERROR)
        
        if (!response.ok) throw Error(response.status)

        return true
    }
}