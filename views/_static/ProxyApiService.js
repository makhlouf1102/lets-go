import { catchError } from "./catchError";
import { JSON_ERROR, SERVER_ERROR } from "./errorsConst";
export class ProxyApiService {

    prefix = "/api/v1"
    api = ApiService
    

    async protectedRoute(request) {
        // TODO : manage refresh token
        // Do stuff to verify the token
        request()
    }

    async login(email, password) {
        [error, response] = await catchError(this.api.login(email, password))

        if (error) throw Error(SERVER_ERROR)
        
        if (!response.ok) {
            // TODO : Manage the bad login responses
        }
        
        [error, responseJson] = await catchError(response.JSON())

        if (error) throw Error(JSON_ERROR)
        
        return responseJson
    }
}