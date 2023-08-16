const available = document.getElementById("available");
const requested = document.getElementById("requested");
const issued = document.getElementById("issued");
const issuedBooks = document.getElementById("issuedBooks");

const div1 = document.getElementById("div1");
const div2 = document.getElementById("div2");
const div3 = document.getElementById("div3");
const div4 = document.getElementById("div4");
let chooseForm = document.getElementById("chooseForm");
let accept = document.getElementById("accept");
let deny = document.getElementById("deny");
let remove = document.getElementById("remove");
let reqForm = document.getElementById("reqForm");
let reqbtn = document.getElementById("reqbtn");
function removeBook() {
    window.location.href = `/admin/remove`
}
div3.style.display = "none";
div2.style.display = "none";
div1.style.display = "block";
div4.style.display = "none";

remove.addEventListener("click", (e) => {
    removeBook();
});

available.addEventListener("click", (e) => {
    div1.style.display = "block";
    div3.style.display = "none";
    div2.style.display = "none";
    div4.style.display = "none";
});

requested.addEventListener("click", (e) => {
    div1.style.display = "none";
    div3.style.display = "none";
    div2.style.display = "block";
    div4.style.display = "none";
});

issued.addEventListener("click", (e) => {
    div1.style.display = "none";
    div2.style.display = "none";
    div3.style.display = "block";
    div4.style.display = "none";
});

issuedBooks.addEventListener("click", (e) => {
    div1.style.display = "none";
    div2.style.display = "none";
    div3.style.display = "none";
    div4.style.display = "block";
});

accept.addEventListener("click", (e) => {
    chooseForm.method = 'POST'
    chooseForm.action = "/admin/choose/accept"
    chooseForm.submit();
});

deny.addEventListener("click", (e) => {
    chooseForm.method = 'POST'
    chooseForm.action = "/admin/choose/deny"
    chooseForm.submit();
});


reqbtn.addEventListener("click", (e) => {
    statuss = reqbtn.innerHTML;
    console.log(statuss);
    if (statuss == "Allow Check-In") {
        reqForm.method = 'POST'
        reqForm.action = "/admin/checkin"
        reqForm.submit();
    } else {
        reqForm.method = 'POST'
        reqForm.action = "/admin/checkout"
        reqForm.submit();
    }


});

