const dropdown = document.getElementById("dropdown");
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

remove.addEventListener("click", (e) => {
    removeBook();
});


div1.style.display = "none";
div3.style.display = "none";
div4.style.display = "none";

console.log("================================")

dropdown.addEventListener('change', function() {

    
    const selectedDivId = dropdown.value;
    console.log(selectedDivId)
    
    if(selectedDivId == "select view mode ")
    {   console.log(selectedDivId)
        div3.style.display = "none";
        div2.style.display = "none";    
        div1.style.display = "block";
        div4.style.display = "none";
    }else if(selectedDivId == "available"){
        div1.style.display = "block";
        div3.style.display = "none";
        div2.style.display = "none";
        div4.style.display = "none";
        console.log(selectedDivId)
    }else if(selectedDivId == "requested"){
        div1.style.display = "none";
        div3.style.display = "none";
        div2.style.display = "block";
        div4.style.display = "none";
        console.log(selectedDivId)
    }else if(selectedDivId == "issued"){
        div1.style.display = "none";
        div2.style.display = "none";
        div3.style.display = "block";
        div4.style.display = "none";
        console.log(selectedDivId)
    }else if(selectedDivId == "issuedBooks"){
        div1.style.display = "none";
        div2.style.display = "none";
        div3.style.display = "none";
        div4.style.display = "block";
    }

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
        if(statuss =="Allow Check-In"){
            reqForm.method = 'POST'
            reqForm.action = "/admin/checkin"
            reqForm.submit();
        }else{
            reqForm.method = 'POST'
            reqForm.action = "/admin/checkout"
            reqForm.submit();
        }
        

    });

