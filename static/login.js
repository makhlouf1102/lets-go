import { proxyApiService } from "./ProxyApiService.js";

document.getElementById('loginForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const email = document.getElementById('email').value
    const password = document.getElementById('password').value

    proxyApiService.login(email, password)
        .then((data)=> {
            window.sessionStorage.setItem('accessToken', data.accessToken);
            window.location.href = '/';
        })
        .catch((error) => {
            console.error('Error:', error);
            alert(error);
        })
});