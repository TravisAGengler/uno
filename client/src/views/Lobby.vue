<template>
  <v-app id="inspire">
    <v-content>
      <v-container class="fill-height" fluid>
        <v-row align="center" justify="center">
          <v-col>
              <GameList @selectedGame="gameSelected($event)"/>
          </v-col>
        </v-row>
        <v-row align="center" justify="center">
            <LobbyControls v-bind:selectedgameid="selected_game_id"/>
        </v-row>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
// @ is an alias to /src
import GameList from '@/components/GameList'
import LobbyControls from '@/components/LobbyControls'
export default {
  name: 'Lobby',
  components: {
    GameList,
    LobbyControls
  },
  data: () => {
    return {
      selected_game_id: ""
    }
  },
  methods: {
    gameSelected(gameId) {
      this.selected_game_id = gameId;
    } 
  }
}
</script>
