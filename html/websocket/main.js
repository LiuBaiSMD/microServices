window.addEventListener("load", function (evt) {
    var defaultIP= "localhost"
    var wsUri = "ws://localhost:8080/websocket/hi"
    var output = document.getElementById("output");
    var nameTxt = document.getElementById("name");
    var ws;

    var print = function (message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };

    function newWebsocekt() {
        ws = new WebSocket(wsUri);

        ws.onopen = function (evt) {
            print('<span style="color: green;">Connection Open</span>');
        }
        ws.onclose = function (evt) {
            print('<span style="color: red;">Connection Closed</span>');
            ws = null;
        }
        ws.onmessage = function (evt) {
            //print('<span style="color: blue;">i get you messenge: </span>' + evt.data);
            print('<span style="color: blue;">i get you messenge! </span>');
        }
        ws.onerror = function (evt) {
            print('<span style="color: red;">Error: </span>' + evt.data);
        }
    }
    newWebsocekt()
    document.getElementById("send").onclick = function (evt) {
        if (!ws) {
            return false
        }

        var msg = {hi: nameTxt.value}
        if (!nameTxt.value){
            return false
        }
        req = JSON.stringify(msg)
        print('<span style="color: blue;">Sent request: </span>' + req);
        ws.send(JSON.stringify(msg));
        nameTxt.value = ""
        return false;
    };

    document.getElementById("cancel").onclick = function (evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        print('<span style="color: red;">Request Canceled</span>');
        return false;
    };

    document.getElementById("open").onclick = function (evt) {
        if (!ws) {
            newWebsocekt()
        }
        return false;
    };
})
