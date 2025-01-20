import { proxyApiService } from "./ProxyApiService.js";

document.getElementById("admin-ping-btn").addEventListener("click", async (e) => {
    proxyApiService.pingProtected()
    .then((data) => {
        console.log(data)
        alert("Well done! You are connected to the server.")
    })
    .catch((error) => {
        console.error("Error", error)
    })

})