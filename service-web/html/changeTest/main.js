window.addEventListener("load", function (evt) {
    var nameTxt = document.getElementById("userId");
    var submit = document.getElementById("submit");


    newWebsocekt()
    document.getElementById("submit").onclick = function (evt) {
        nameTxt.value = "test"
        submit.value = "yes"
    }
})