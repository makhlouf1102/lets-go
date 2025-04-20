const components  = document.getElementsByClassName("auth-only");

var c;

for (let i = 0; i < components.length; i++) {
    c = components[i];
    c.style.display = "none"
}