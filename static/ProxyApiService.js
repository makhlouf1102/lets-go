import { catchError } from './catchError.js';
import { JSON_ERROR, SERVER_ERROR, BAD_CREDENTIALS_ERROR, INVALID_AUTH_ERROR, NEW_TOKEN_ERROR } from './errorsConst.js';
import { apiService } from './ApiService.js';

export const proxyApiService = {
    async protectedRouteJson(request) {
        if (typeof request !== "function") {
            throw new Error("The request parameter must be a function.");
        }

        let [error, response] = await catchError(request());

        if (error) {
            return [error, null];
        }

        if (!response.ok && (response.status === 403 || response.status === 401)) {
            window.sessionStorage.removeItem("accessToken");
            return [new Error(INVALID_AUTH_ERROR), null];
        }

        let [errorJson, responseJson] = await catchError(response.json());

        if (errorJson) {
            return [new Error(JSON_ERROR), null];
        }

        if (responseJson.status === "New Token") {
            window.sessionStorage.setItem('accessToken', responseJson.accessToken);

            return await catchError(request());
        }

        return [null, responseJson];
    },

    async protectedRouteText(request) {
        if (typeof request !== "function") {
            throw new Error("The request parameter must be a function.");
        }

        let [error, response] = await catchError(request());

        if (error) {
            return [error, null];
        }

        if (!response.ok && (response.status === 403 || response.status === 401)) {
            window.sessionStorage.removeItem("accessToken");
            return [new Error(INVALID_AUTH_ERROR), null];
        }

        let [errorText, responseText] = await catchError(response.text());

        if (errorText) {
            return [new Error("Failed to parse text response"), null];
        }

        return [null, responseText];
    },

    async login(email, password) {
        let [error, response] = await catchError(apiService.login(email, password))

        if (error) throw Error(SERVER_ERROR)

        if (!response.ok) {
            const responseStatus = response.status

            if (responseStatus == "400" || responseStatus == "500")
                throw Error(SERVER_ERROR)

            if (responseStatus == "401")
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
        let [error, response] = await this.protectedRouteJson(() => apiService.pingProtected());

        if (error) throw new Error(SERVER_ERROR);

        return response;
    },

    async getProblemCode(programmingLanguage, problemId) {
        let [error, response] = await this.protectedRouteText(() => apiService.getProblemCode(programmingLanguage, problemId))
        
        if (error) throw new Error(SERVER_ERROR);

        return response;
    },

    async runCode(programmingLanguage, code) {
        let [error, response] = await this.protectedRouteText(() => apiService.runCode(programmingLanguage, code))
        if (error) throw new Error(SERVER_ERROR);

        return response;
    } 
}
