import React from 'react';


import './App.css';

let socket = new WebSocket("ws://172.21.0.3:8074/chat");

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

type Message = {
    Username: string;
    Message: string;
};

function sendData() {
    let elName: HTMLInputElement = (document.getElementById("name") as HTMLInputElement);
    let elMessage: HTMLInputElement = (document.getElementById("message") as HTMLInputElement);
    if (elName.value && elMessage.value) {
        let msg: Message = {Username: elName.value, Message: elMessage.value};

        socket.send(JSON.stringify(msg));
    }
}

export default class TechView extends React.Component<{}, {}> {
    render() {
        return (
            <div className="Appr">
                <header className="App-header">
                    <p>
                        <input id={"name"} type="text"/>
                        <input id={"message"} type="text"/>
                        <button onClick={sendData}> Клик</button>
                    </p>
                </header>
                <div id={"data"}>

                </div>
            </div>
        );
    }
}
