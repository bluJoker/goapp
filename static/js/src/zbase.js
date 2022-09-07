export class AcGame {
    constructor(id) {
        this.id = id; //传进i来的id为div-id
        this.$ac_game = $('#' + id); //jquery根据id找div元素

        this.settings = new Settings(this);
        this.menu = new AcGameMenu(this);
        this.playground = new AcGamePlayground(this);
    }
}

