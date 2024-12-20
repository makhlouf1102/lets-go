import { ACCESS_TOKEN_NAME, HAS_REFRESH_TOKEN } from "./constants";
import { proxyApiService } from "./ProxyApiService";

// find the the right cookie
const cookies = document.cookie.split(";")
const cookieExists = cookies.find((row) => row.startsWith(HAS_REFRESH_TOKEN))

if (!sessionStorage.getItem(ACCESS_TOKEN_NAME) && cookieExists) {
    proxyApiService.pingProtected()
        .then((data) => {
            window.sessionStorage.setItem(ACCESS_TOKEN_NAME, data.accessToken);
            window.location.reload()
        })
        .catch((error) => {
            console.error('Error:', error);
        })
} 