const available = document.getElementById("available");
const requested = document.getElementById("requested");

const div1 = document.getElementById("div1");
const div2 = document.getElementById("div2");

div2.style.display = "none";
div1.style.display = "block";

available.addEventListener("click", (e) => {
    div1.style.display = "block";
    div2.style.display = "none";
});

requested.addEventListener("click", (e) => {
    div1.style.display = "none";
    div2.style.display = "block";
});
