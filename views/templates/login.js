import { catchError } from "../_static/catchError";
import { ApiService } from "../_static/ApiService";

document.getElementById('loginForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const email = document.getElementById('email').value
    const password = document.getElementById('password').value

    const [error, response] = await catchError(ApiService.login(email, password))

    if (error) {
        console.error('Error:', error);
        alert('An error occurred. Please try again.');
    }

    if (response.ok) {
        const data = await response.json();
        window.sessionStorage.setItem('refreshToken', data.refreshToken);
        window.sessionStorage.setItem('accessToken', data.accessToken);
        window.location.href = '/';
    } else {
        alert('Login failed. Please try again.');
    }
});