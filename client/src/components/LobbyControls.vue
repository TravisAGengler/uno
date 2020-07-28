<template>
  <v-card class="elevation-12">
    <v-dialog v-model="show_create_game_dialog" max-width="500">
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          color="primary"
          dark
          v-bind="attrs"
          v-on="on"
        >
          Create Game
        </v-btn>
      </template>
      <v-card>
        <v-card-title class="headline">Create New Game</v-card-title>
        <v-form>
          <v-col cols="12" sm="12" md="12">
            <v-text-field
              v-model="create_game_name"
              label="Game Name">
            </v-text-field>
          </v-col>
          <v-col cols="12" sm="12" md="12">
            <v-text-field
              v-model="create_game_password"
              :append-icon="show_password ? 'mdi-eye' : 'mdi-eye-off'"
              :type="show_password ? 'text' : 'password'"
              name="password"
              label="Password (Optional)"
              class="input-group--focused"
              @click:append="show_password = !show_password"
            ></v-text-field>
          </v-col>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" text @click="createGame">Create</v-btn>
          </v-card-actions>
          <v-card v-if="create_game_err != ''">{{ create_game_err }}</v-card>
        </v-form>
      </v-card>
    </v-dialog>

    <v-btn test-id="game-list-join" @click.native="joinGame" color="primary" :disabled="!selectedgameid" >Join Game</v-btn>
  </v-card>
</template>

<script>
//import unoService from "../services/unoService";
export default {
  name: "LobbyControls",
  data: () => {
    return {
      create_game_err: "",
      create_game_name: "",
      create_game_password: "",
      show_password: false,
      show_create_game_dialog: false,
    };
  },
  props:
  [
    'selectedgameid'
  ],
  methods: {
    async joinGame() {
      // TODO: Join a game
      // We may want to show a modal here to prompt for passwords if required, before sending the request
      console.log("Join Game");
      console.log(this.selectedgameid);
    },
    async createGame () {
      this.clearErrMsg();
      // Cant create game with no name. Display an error for the user?
      if (this.create_game_name === "") {
        this.create_game_err = "Please provide a game name";
      }

      // TODO: Call the createGame endpoint
      // If successful, close the dialog. If failure, display a message and keep dialog open
      if (this.create_game_err === "") {
        this.show_create_game_dialog = false;
        this.create_game_name = "";
        this.create_game_password = "";
      }
    },
    clearErrMsg () {
      this.create_game_err = "";
    }
  },
  watch :
  {
    show_create_game_dialog(new_val) {
      !new_val && this.clearErrMsg()
    }
  }
};
</script>
