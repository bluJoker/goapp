class Settings {
    constructor(root) {
        this.root = root;
        this.email = "";
        this.photo = "";

        this.$settings = $(`
<div class="ac-game-settings">
    <div class="ac-game-settings-login">
        <div class="ac-game-settings-title">
            登录
        </div>
        <div class="ac-game-settings-email">
            <div class="ac-game-settings-item">
                <input type="text" placeholder="Email address">
            </div>
        </div>
        <div class="ac-game-settings-password">
            <div class="ac-game-settings-item">
                <input type="password" placeholder="Password">
            </div>
        </div>
        <div class="ac-game-settings-submit">
            <div class="ac-game-settings-item">
                <button>登录</button>
            </div>
        </div>
        <div class="ac-game-settings-error-message">
        </div>
        <div class="ac-game-settings-option">
            注册
        </div>
        <br>
        <div class="ac-game-settings-acwing">
            <img width="50" src="http://118.31.14.70:8087/static/image/menu/shenle_xiao.jpg">
            <br>
            <div>
                w2一键登录
            </div>
        </div>
    </div>
    <div class="ac-game-settings-register">
        <div class="ac-game-settings-title">
            注册
        </div>
        <div class="ac-game-settings-username">
            <div class="ac-game-settings-item">
                <input type="text" placeholder="Username">
            </div>
        </div>
        <div class="ac-game-settings-email">
            <div class="ac-game-settings-item">
                <input type="text" placeholder="Email address">
            </div>
        </div>
        <div class="ac-game-settings-password ac-game-settings-password-first">
            <div class="ac-game-settings-item">
                <input type="password" placeholder="Password">
            </div>
        </div>
        <div class="ac-game-settings-password ac-game-settings-password-second">
            <div class="ac-game-settings-item">
                <input type="password" placeholder="Confirm password">
            </div>
        </div>
        <div class="ac-game-settings-submit">
            <div class="ac-game-settings-item">
                <button>Register</button>
            </div>
        </div>
        <div class="ac-game-settings-error-message">
        </div>
        <div class="ac-game-settings-option">
            登录
        </div>
        <br>
        <div class="ac-game-settings-acwing">
            <img width="50" src="https://app2672.acapp.acwing.com.cn/static/image/menu/shenle_xiao.jpg">
            <br>
            <div>
                w2一键登录
            </div>
        </div>

    </div>
</div>
    `);

        this.$login = this.$settings.find(".ac-game-settings-login");

        this.$login_email = this.$login.find(".ac-game-settings-email input");
        this.$login_password = this.$login.find(".ac-game-settings-password input");
        this.$login_submit = this.$login.find(".ac-game-settings-submit button");
        this.$login_error_message = this.$login.find(".ac-game-settings-error-message");
        this.$login_register = this.$login.find(".ac-game-settings-option");

        this.$login.hide();

        this.$register = this.$settings.find(".ac-game-settings-register");
        this.$register_username = this.$register.find(".ac-game-settings-username input");
        this.$register_email = this.$register.find(".ac-game-settings-email input");
        this.$register_password = this.$register.find(".ac-game-settings-password-first input");
        this.$register_password_confirm = this.$register.find(".ac-game-settings-password-second input");
        this.$register_submit = this.$register.find(".ac-game-settings-submit button");
        this.$register_error_message = this.$register.find(".ac-game-settings-error-message");
        this.$register_login = this.$register.find(".ac-game-settings-option");

        this.$register.hide();

        this.root.$ac_game.append(this.$settings);

        this.start();
    }

    start() {
        this.getinfo(); // 创建时从服务器端获取用户信息
        this.add_listening_events();
    }

    login() {  // 打开登录界面
        this.$register.hide();
        this.$login.show();
    }

    register() {  // 打开注册界面
        this.$login.hide();
        this.$register.show();
    }

    add_listening_events() {
        this.add_listening_events_login();
        this.add_listening_events_register();
    }

    add_listening_events_login() {
        let outer = this;
        this.$login_register.click(function() {
            outer.register();
        });
        this.$login_submit.click(function() {
            outer.login_on_remote();
        });
    }

    add_listening_events_register() {
        let outer = this;
        this.$register_login.click(function() {
            outer.login();
        });
    }

    login_on_remote() {  // 在远程服务器上登录
        let outer = this;
        let email = this.$login_email.val(); //val: 取出input的值
        let password = this.$login_password.val();
        console.log("login_on_remote", email, password)
        this.$login_error_message.empty(); //每次登陆清空上次的errmsg

        $.ajax({
            url: "http://118.31.14.70:8087/settings/login",
            type: "POST",
            data: {
                email: email,
                password: password,
            },
            success: function(resp) {
                console.log("login_on_remote-resp: ", resp)
                if (resp.result === "success") {
                    location.reload(); //逻辑: 登录成功就刷新页面, 调用getinfo可以获取cookie的用户信息, 登录成功会显示菜单页面[getinfo: outer.root.menu.show();]
                } else {
                    outer.$login_error_message.html(resp.result);
                }
            }
        });
    }

    getinfo() {
        let outer = this;

        $.ajax({
            url: "http://118.31.14.70:8087/settings/getinfo",
            type: "GET",
            success: function(resp) {
                console.log("getinfo", resp);
                if (resp.result === "success") {
                    outer.email = resp.email;
                    outer.photo = resp.photo;
                    // 获取用户信息成功，表示已登录。将当前页面隐藏，打开菜单页面
                    outer.hide();
                    outer.root.menu.show();
                } else {
                    // 获取用户信息失败，打开登录界面
                    outer.login();
                }
            }
        });
    }

    hide() {
        this.$settings.hide();
    }

    show() {
        this.$settings.show();
    }
}

