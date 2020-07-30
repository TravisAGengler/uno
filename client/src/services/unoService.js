import BaseService from "./baseService";

export default {
    async login(gameId, userName) {
        return BaseService.post(`/login/${gameId}/${userName}`);
    },

    async newGame() {
        return BaseService.get("/newgame");
    },

    async createNewGame(gameName, password) {
        var data = {
            "name" : gameName,
            "password" : password
        };
        return BaseService.post("/games", data);
    },

    update(gameId, userName) {
        return BaseService.get(`/update/${gameId}/${userName}`);
    },

    startGame(gameId, userName) {
        return BaseService.post(`/startgame/${gameId}/${userName}`);
    },

    playCard(gameId, userName, cardNumber, cardColor) {
        return BaseService.post(`/play/${gameId}/${userName}/${cardNumber}/${cardColor}`);
    },

    drawCard(gameId, userName) {
        return BaseService.post(`/draw/${gameId}/${userName}`);
    },

    getGameList() {
        return BaseService.get(`/games`);
    }
}