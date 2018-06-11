var url = "http://localhost:9090/note/add";
xmlhttp = new XMLHttpRequest();

xmlhttp.onreadystatechange = state_Change;
xmlhttp.open("POST", url, true);
xmlhttp.setRequestHeader("Content-Type", "application/json; charset=utf-8");
var params = {
    userId: 1,
    username: "alan",
    content: "this is a note",
};
xmlhttp.send(JSON.stringify(params));
var state_Change = function () {
    console.log(xmlhttp.readyState, xmlhttp.status)
    if (xmlhttp.readyState == 4) {// 4 = "loaded"
        if (xmlhttp.status == 200) {// 200 = OK
        }
        else {
            console.log("Problem retrieving XML data");
        }
    }
}
//---------------------------------------------------------

var url = "http://localhost:9090/note/add";
xmlhttp = new XMLHttpRequest();

xmlhttp.onreadystatechange = state_Change;
xmlhttp.open("POST", url, true);
xmlhttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded; charset=utf-8");

var params = "userID=1&content=thisisanote"
xmlhttp.send(params);
var state_Change = function () {
    console.log(xmlhttp.readyState, xmlhttp.status)
    if (xmlhttp.readyState == 4) {// 4 = "loaded"
        if (xmlhttp.status == 200) {// 200 = OK
        }
        else {
            console.log("Problem retrieving XML data");
        }
    }
}
