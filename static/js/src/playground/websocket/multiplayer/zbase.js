// 实现前端与server的wss连接
class MultiPlayerSocket {
    constructor(playground) {
        this.playground = playground;
        this.ws = new WebSocket("ws://118.31.14.70:8087/wss/multiplayer"); // 向后端发起建立连接请求

        this.start();
    }

    start() {
        this.receive();
    }

    send_create_player(email, photo) {
        let outer = this;
        console.log("send: ", outer.uuid, email, photo)
        this.ws.send(JSON.stringify({ // 将json转为字符串
            'event': "create_player",
            'uuid': outer.uuid, // playground/zbase.js 中的 this.mps.uuid
            'email': email,
            'photo': photo,
        }));
    }

    receive_create_player(uuid, email, photo) {
        console.log("receive_create_player: ", email, photo)
        let player = new Player(
            this.playground,
            this.playground.width / 2 / this.playground.scale,
            0.5,
            0.04,
            "white",
            0.2,
            "enemy",
            email,
            photo,
        );

        player.uuid = uuid; // 等于创建它的窗口的uuid
        this.playground.players.push(player);
    }

    receive() {
        let outer = this;

        // 前端接收wss信息
        this.ws.onmessage = function(e) {
            let data = JSON.parse(e.data); // 将字符串转换为字典
            console.log("receive: ", data);
            let uuid = data.uuid;
            if (uuid === outer.uuid) {
                console.log("self: ", uuid)
                return false;
            }
            let event = data.event;
            if (event === "create_player") {
                outer.receive_create_player(uuid, data.email, data.photo);
            }
        };
    }

}

