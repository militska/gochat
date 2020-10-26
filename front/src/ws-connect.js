

let socket = new WebSocket("ws://172.20.0.2:8074/echo");

socket.onopen = function() {
    alert("Соединение установлено.");
};

socket.onclose = function(event) {
    if (event.wasClean) {
        alert('Соединение закрыто чисто');
    } else {
        alert('Обрыв соединения'); // например, "убит" процесс сервера
    }
    alert('Код: ' + event.code + ' причина: ' + event.reason);
};

socket.onmessage = function(event) {
    window.console.log("Получены данные " + event.data);
    var elements = document.getElementsByClassName("data");

    elements[0].innerHTML += " <br>" + event.data;
};

socket.onerror = function(error) {
    alert(error);
};

