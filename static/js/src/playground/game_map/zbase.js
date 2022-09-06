// 作用：游戏地图对象
class GameMap extends AcGameObject {
    constructor(playground) {
        super(); // 调用基类构造函数；将自己注册到AC_GAME_OBJECT数组中
        this.playground = playground;
        this.$canvas = $(`<canvas></canvas>`);
        this.ctx = this.$canvas[0].getContext('2d');
        this.ctx.canvas.width = this.playground.width;
        this.ctx.canvas.height = this.playground.height;
        this.playground.$playground.append(this.$canvas);
    }

    start() {
    }

    update() {
        this.render(); //每帧画一次canvas
    }

    render() { //渲染
        // choose color first, then draw rectangle
        this.ctx.fillStyle = "rgba(0, 0, 0, 0.5)"; //黑色, 半透明
        this.ctx.fillRect(0, 0, this.ctx.canvas.width, this.ctx.canvas.height);
    }
}

