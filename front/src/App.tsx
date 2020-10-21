import React from 'react';
import './App.css';

var socket = new WebSocket("ws://172.21.0.4:8074/echo");

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


function sendData() {
  let el : string = (document.getElementById("new_value") as HTMLInputElement).value;
  socket.send(el);
}


function App() {
  return (
    <div className="App">
      <header className="App-header">
        <p>
          <input id={"new_value"} type="text"/>

          <button onClick={sendData}> Клик </button>

        </p>
      </header>
      <div className={"data"}>
      ttt

      </div>
    </div>
  );
}

export default App;
