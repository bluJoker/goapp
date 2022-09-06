// 作用：人物对象
class Player extends AcGameObject {
    // x, y: 球中心点坐标; speed:单位高度百分比，避免分辨率不同; is_me:表示是否是自己，自己和非自己的移动方式不同，自己通过鼠标键盘，其他人通过网络传递玩家移动消息
    constructor(playground, x, y, radius, color, speed, is_me) {
        super();
        this.playground = playground;
        this.ctx = this.playground.game_map.ctx; // 画布的引用

        this.x = x;
        this.y = y;
        //vx, vy:速度
        this.vx = 0;
        this.vy = 0;
        this.radius = radius;
        this.color = color;
        this.speed = speed;
        this.is_me = is_me;

        // 精度:浮点运算中小于多少算0
        this.eps = 0.1;
        this.move_length = 0;
    }

    get_dist(x1, y1, x2, y2) { // 两点间欧几里得距离
        let dx = x1 - x2;
        let dy = y1 - y2;
        return Math.sqrt(dx * dx + dy * dy);
    }


    move_to(tx, ty) {
        this.move_length = this.get_dist(this.x, this.y, tx, ty);
        //求角度 反正切
        let angle = Math.atan2(ty - this.y, tx - this.x);
        //vx:横方向速度, vy:纵方向速度
        this.vx = Math.cos(angle);
        this.vy = Math.sin(angle);
    }

    add_listening_events() { // 鼠标点击事件
        let outer = this;
        this.playground.game_map.$canvas.on("contextmenu", function() { // 禁用右键显示菜单事件["contextmenu"]
            return false;
        });
        this.playground.game_map.$canvas.mousedown(function(e) {
            //3:鼠标右键, 1:左键, 2:滚轮
            if (e.which === 3) {
                outer.move_to(e.clientX, e.clientY);
            }
        });
    }

    start() {
        if (this.is_me) {
            this.add_listening_events();
        }
    }

    update() {
        if (this.move_length < this.eps) {
            this.move_length = 0;
            this.vx = this.vy = 0;
        } else {
            // 计算每帧移动距离
            let moved = Math.min(this.move_length, this.speed * this.timedelta / 1000); // 两点间距与update时间内移动距离的小值，因为可能在update时间内就提前到达终点了
            this.x += this.vx * moved; // 加上x方向上需要移动的距离即为此次刷新终点
            this.y += this.vy * moved;
            this.move_length -= moved; // 更新移动距离直到到达上面的if精度条件->表示到达鼠标终点
        }

        this.render();
    }

    render() {
        this.ctx.beginPath();
        this.ctx.arc(this.x, this.y, this.radius, 0, Math.PI * 2, false);
        this.ctx.fillStyle = this.color;
        this.ctx.fill();
    }
}

