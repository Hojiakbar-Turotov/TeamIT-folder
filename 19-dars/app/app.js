const startSnow = document.getElementById("startSnowing");
const closedSnow = document.getElementById("closedSnowing");
let harakat = null;
function startSnowing() {
    harakat = setInterval(createSnow, 100);
    startSnow.classList.add("d-none");
    closedSnow.classList.remove("d-none");
}
function closedSnowing() {
    clearTimeout(harakat);
    startSnow.classList.remove("d-none");
    closedSnow.classList.add("d-none");
}
function createSnow() {
    const snow = document.createElement("i");
    snow.classList.add("fa-snowflake");
    snow.classList.add("fas");
    snow.style.left = Math.random() * window.innerWidth + "px";
    snow.style.animationDuration = Math.random() * 3 + 2 + "s";
    snow.style.opacity = Math.random();
    snow.style.fontSize = Math.random() * 10 + 10 + "px";
    document.body.append(snow);
}