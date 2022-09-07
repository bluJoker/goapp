class Settings {
    constructor(root) {
        this.root = root;
        this.username = "";
        this.photo = "";

        this.start();
    }

    start() {
        this.getinfo(); // 创建时从服务器端获取用户信息
    }

    login() {  // 打开登录界面
        //this.$register.hide();
        //this.$login.show();
    }

    register() {  // 打开注册界面
        //this.$login.hide();
        //this.$register.show();
    }

    getinfo() {
        let outer = this;

        $.ajax({
            url: "http://118.31.14.70:8087/settings/getinfo",
            type: "GET",
            success: function(resp) {
                console.log(resp);
                if (resp.result === "success") {
                    outer.username = resp.username;
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
        //this.$settings.hide();
    }

    show() {
        //this.$settings.show();
    }
}

