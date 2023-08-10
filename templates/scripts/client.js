const dropdown = document.getElementById("dropdown");
const div1 = document.getElementById("div1");
const div2 = document.getElementById("div2");
const div3 = document.getElementById("div3");


div1.style.display = "none";
// div2.style.display = "none";

dropdown.addEventListener('change', function() {

    
    const selectedDivId = dropdown.value;
    console.log(selectedDivId)
    
    if(selectedDivId == "select view mode ")
    {   console.log(selectedDivId)
        div2.style.display = "none";    
        div1.style.display = "block";
    }else if(selectedDivId == "available"){
        div1.style.display = "block";
        div2.style.display = "none";
        console.log(selectedDivId)
    }else if(selectedDivId == "requested"){
        div1.style.display = "none";
        div2.style.display = "block";
        console.log(selectedDivId)
    }
});

