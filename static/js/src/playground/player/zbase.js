// 作用：人物对象
class Player extends AcGameObject {
    // x, y: 球中心点坐标; speed:单位高度百分比，避免分辨率不同; is_me:表示是否是自己，自己和非自己的移动方式不同，自己通过鼠标键盘，其他人通过网络传递玩家移动消息
    constructor(playground, x, y, radius, color, speed, is_me) {
        super();
        this.playground = playground;
        this.ctx = this.playground.game_map.ctx; // 画布的引用

        this.x = x;
        this.y = y;
        // vx, vy:速度
        this.vx = 0;
        this.vy = 0;

        //被击中伤害后的速度等参数
        this.damage_x = 0;
        this.damage_y = 0;
        this.damage_speed = 0;
        //被击中后的摩擦力
        this.friction = 0.9;

        this.radius = radius;
        this.color = color;
        this.speed = speed;
        this.is_me = is_me;

        // 精度:浮点运算中小于多少算0
        this.eps = 0.1;
        this.move_length = 0;

        //当前选择的技能
        this.cur_skill = null;

        this.spent_time = 0;
    }

    get_dist(x1, y1, x2, y2) { // 两点间欧几里得距离
        let dx = x1 - x2;
        let dy = y1 - y2;
        return Math.sqrt(dx * dx + dy * dy);
    }


    move_to(tx, ty) {
        this.move_length = this.get_dist(this.x, this.y, tx, ty);
        // 求角度 反正切
        let angle = Math.atan2(ty - this.y, tx - this.x);
        // vx:横方向速度, vy:纵方向速度
        this.vx = Math.cos(angle);
        this.vy = Math.sin(angle);
    }

    shoot_fireball(tx, ty) {
        let x = this.x, y = this.y;
        let radius = this.playground.height * 0.01;
        let angle = Math.atan2(ty - this.y, tx - this.x);
        let vx = Math.cos(angle), vy = Math.sin(angle);
        let color = "orange";
        let speed = this.playground.height * 0.3;
        let move_length = this.playground.height * 1;
        new FireBall(this.playground, this, x, y, radius, vx, vy, color, speed, move_length, this.playground.height * 0.005);
    }

    add_listening_events() { // 鼠标点击事件
        let outer = this;
        this.playground.game_map.$canvas.on("contextmenu", function() { // 禁用右键显示菜单事件["contextmenu"]
            return false;
        });
        this.playground.game_map.$canvas.mousedown(function(e) {
            // 3:鼠标右键, 1:左键, 2:滚轮
            if (e.which === 3) {
                outer.move_to(e.clientX, e.clientY);
            } else if (e.which === 1) {
                if (outer.cur_skill === "fireball") { // 当前技能为火球，表示按了q键
                    outer.shoot_fireball(e.clientX, e.clientY);
                }
                outer.cur_skill = null; //点完左键释放掉技能
            }
        });

        // 此处获取键盘的事件不能使用canvas:因为canvas不能聚焦？HTML5 Canvas本身不支持键盘事件监听与获取
        $(window).keydown(function(e) {
            if (e.which === 81) { // keycode中q键:81
                outer.cur_skill = "fireball";
                return false;
            }
        });
    }

    // 玩家被火球击中，所有被作用的函数写在对应的类中
    is_attacked(angle, damage) {
        // 碰撞粒子效果
        for (let i = 0; i < 20 + Math.random() * 10; i ++ ) {
            let x = this.x, y = this.y;
            let radius = this.radius * Math.random() * 0.1;
            let angle = Math.PI * 2 * Math.random();
            let vx = Math.cos(angle), vy = Math.sin(angle);
            let color = this.color;
            let speed = this.speed * 10;
            let move_length = this.radius * Math.random() * 5;
            new Particle(this.playground, x, y, radius, vx, vy, color, speed, move_length);
        }

        //玩家血量[半径]减去伤害值
        this.radius -= damage;
        if (this.radius < 5) { // 半径小于5像素认为死亡
            this.destroy();
            return false;
        }
        this.damage_x = Math.cos(angle);
        this.damage_y = Math.sin(angle);
        this.damage_speed = damage * 45;
        //血量减少，但速度变快
        this.speed *= 1.25;
    }

    start() {
        if (this.is_me) {
            this.add_listening_events();
        } else {
            let tx = Math.random() * this.playground.width;
            let ty = Math.random() * this.playground.height;
            this.move_to(tx, ty);
        }
    }

    update() {
        //单机模式中其他玩家随机攻击
        this.spent_time += this.timedelta / 1000;
        if (!this.is_me && this.spent_time > 4 && Math.random() < 1 / 300.0) { // 每秒刷新60次，概率1/300相当于5s发射一次
            let player = this.playground.players[Math.floor(Math.random() * this.playground.players.length)]; // 随机向一名玩家发射炮弹
            let tx = player.x + player.speed * this.vx * this.timedelta / 1000 * 0.3;
            let ty = player.y + player.speed * this.vy * this.timedelta / 1000 * 0.3;
            this.shoot_fireball(tx, ty);
        }

        //伤害消失。10表示被撞后速度<=10就不管它了，让其再次随机运动。
        if (this.damage_speed > 10) {
            this.vx = this.vy = 0;
            this.move_length = 0;
            this.x += this.damage_x * this.damage_speed * this.timedelta / 1000;
            this.y += this.damage_y * this.damage_speed * this.timedelta / 1000;
            this.damage_speed *= this.friction;
        } else {
            if (this.move_length < this.eps) {
                this.move_length = 0;
                this.vx = this.vy = 0;

                if (!this.is_me) { // robot停下来需要继续随机移动
                    let tx = Math.random() * this.playground.width;
                    let ty = Math.random() * this.playground.height;
                    this.move_to(tx, ty);
                }
            } else {
                // 计算每帧移动距离
                let moved = Math.min(this.move_length, this.speed * this.timedelta / 1000); // 两点间距与update时间内移动距离的小值，因为可能在update时间内就提前到达终点了
                this.x += this.vx * moved; // 加上x方向上需要移动的距离即为此次刷新终点
                this.y += this.vy * moved;
                this.move_length -= moved; // 更新移动距离直到到达上面的if精度条件->表示到达鼠标终点
            }
        }
        this.render();
    }

    render() {
        this.ctx.beginPath();
        this.ctx.arc(this.x, this.y, this.radius, 0, Math.PI * 2, false);
        this.ctx.fillStyle = this.color;
        this.ctx.fill();
    }

    on_destroy() {
        for (let i = 0; i < this.playground.players.length; i++) {
            if (this.playground.players[i] === this) {
                this.playground.players.splice(i, 1);
            }
        }
    }
}

