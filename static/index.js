import { proxyApiService } from "./ProxyApiService.js";

document.getElementById('testtest').addEventListener('click', async (e) => {
    e.preventDefault();


    proxyApiService.pingProtected()
        .then(()=> {
            alert("success")
        })
        .catch((error) => {
            console.error('Error:', error);
            alert(error);
        })
});