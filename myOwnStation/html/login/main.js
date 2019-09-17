window.addEventListener("load", function (evt) {
    var passwordTxt = document.getElementById("password");
    var nameTxt = document.getElementById("name");
    var ws;

    // var print = function (message) {
    //     var d = document.createElement("div");
    //     d.innerHTML = message;
    //     output.appendChild(d);
    // };

    function newWebsocekt() {
        ws = new WebSocket(wsUri);

    }
    newWebsocekt()
    document.getElementById("submit").onclick = function (evt) {
        if (!ws) {
            return false
        }
        var name = nameTxt.value
        var password = passwordTxt.value
        var msg = {hi: nameTxt.value}
        print(name, password)


})