<template>
<v-container>
  <v-btn fab dark large color="primary" fixed right bottom @click.native="show_create_game_dialog = true">
    <v-icon>mdi-plus</v-icon>
  </v-btn>
  <v-dialog v-model="show_create_game_dialog" max-width="500">
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
</v-container>
</template>

<script>
import unoService from "../services/unoService";
export default {
  name: "CreateGameFAB",
  data: () => {
    return {
      create_game_err: "",
      create_game_name: "",
      create_game_password: "",
      show_password: false,
      show_create_game_dialog: false,
    };
  },
  methods: {
    async createGame () {
      // Cant create game with no name. Display an error for the user
      if (this.create_game_name === "")
        this.create_game_err = "Please provide a game name";
      else
        this.create_game_err = "";

      // TODO: Call the createGame endpoint
      let res = await unoService.createNewGame(this.create_game_name, this.create_game_password);
      // If successful, close the dialog. If failure, display a message and keep dialog open
      console.log(res.data);

      if (this.create_game_err === "") {
        this.show_create_game_dialog = false;
        this.create_game_name = "";
        this.create_game_password = "";
      }
    },
    resetDialog () {
      this.create_game_err = "";
      this.create_game_name = "";
      this.create_game_password = "";
    }
  },
  watch :
  {
    show_create_game_dialog(new_val) {
      !new_val && this.resetDialog()
    }
  }
};
</script>
