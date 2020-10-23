import React from 'react';


import './App.css';

let socket = new WebSocket("ws://172.20.0.2:8074/echo");

socket.onopen = function () {
    window.alert("Соединение установлено.");
};

socket.onclose = function (event) {
    if (event.wasClean) {
        window.alert('Соединение закрыто чисто');
    } else {
        window.alert('Обрыв соединения'); // например, "убит" процесс сервера
    }
    window.alert('Код: ' + event.code + ' причина: ' + event.reason);
};

socket.onmessage = function (event) {
    console.log("Получены данные " + event.data);
    let el: HTMLDivElement = (document.getElementById("data") as HTMLDivElement);
    if (el) {
        el.innerHTML += " <br>" + event.data;
    }
};

socket.onerror = function (error) {
    window.alert(error);
};


function sendData() {
    let el: HTMLInputElement = (document.getElementById("new_value") as HTMLInputElement);
    if (el.value) {
        let val: string = el.value;
        socket.send(val);
    }
}

export default class TechView4 extends React.Component<{}, {}> {
    render() {
        return (
            <div className="Appr">rrrrrrrrrrrrrrrrrr
                <header className="App-header">
                    <p>
                        <input id={"new_value"} type="text"/>
                        <button onClick={sendData}> Клик</button>
                    </p>
                </header>
                <div id={"data"}>

                </div>
            </div>
        );
    }
}
