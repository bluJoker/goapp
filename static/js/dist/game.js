class AcGameMenu {
    constructor(root) { //root: 即web.html中的js:acgame对象
        this.root = root;
        this.$menu = $(`
<div class="ac-game-menu">
    <div class="banner-middle">
        <h4>一款现代多人游戏杰作-《DESTRUCTOID》</h4>
    </div>
    <div class="ac-game-menu-field">
        <div class="ac-game-menu-field-item ac-game-menu-field-item-single">
        单人模式
    </div>
    <br>
    <div class="ac-game-menu-field-item ac-game-menu-field-item-multi">
        多人模式
    </div>
    <br>
    <div class="ac-game-menu-field-item ac-game-menu-field-item-settings">
        退出
    </div>
    </div>
</div>
    `);//html对象前加$，普通对象不加$
        this.root.$ac_game.append(this.$menu); //向ac_game对象加入该menu页面对象
        this.$single = this.$menu.find('.ac-game-menu-field-item-single'); //class前加. id前加#
        this.$multi = this.$menu.find('.ac-game-menu-field-item-multi'); //class前加. id前加#
        this.$settings = this.$menu.find('.ac-game-menu-field-item-settings'); //class前加. id前加#

        this.start();
    }

    start() {
        this.add_listening_events();
    }

    add_listening_events() {
        let outer = this;
        this.$single.click(function(){
            outer.hide();
            outer.root.playground.show();
        });
        this.$multi.click(function(){
            console.log("click mulit mode");
        });
        this.$settings.click(function(){
            console.log("click settings");
        });
    }

    show() {  // 显示menu界面
        this.$menu.show();
    }

    hide() {  // 关闭menu界面
        this.$menu.hide();
    }
}

class AcGamePlayground {
    constructor(root) {
        this.root = root;
        this.$playground = $(`<div>游戏界面</div>`);

        this.hide(); // 一开始显示的是菜单界面, 点击 "单人模式" 才打开游戏界面

        this.root.$ac_game.append(this.$playground);
        this.start();
    }

    start() {
    }

    show() {  // 打开playground界面
        this.$playground.show();
    }

    hide() {  // 关闭playground界面
        this.$playground.hide();
    }
}

class AcGame {
    constructor(id) {
        this.id = id; //传进i来的id为div-id
        this.$ac_game = $('#' + id); //jquery根据id找div元素

        this.menu = new AcGameMenu(this);
        this.playground = new AcGamePlayground(this);
    }
}

