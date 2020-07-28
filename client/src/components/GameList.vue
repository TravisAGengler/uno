<template>
  <v-card class="elevation-12">
    <v-toolbar color="primary" dark flat>
      <v-toolbar-title>Game List</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-tooltip bottom></v-tooltip>
    </v-toolbar>
    <v-data-table
      :headers="headers"
      :items="game_list"
      :items-per-page="5"
      :single-select="true"
    >
    <template v-slot:body="{ items }">
        <tbody>
          <tr v-for="item in items" :key="item.Name" 
          @click="selectItem(item)" 
          :class="{'selected': item === selected}">
            <td>{{ item.Name }}</td>
            <td>{{ item.NumPlayers }}</td>
            <td>{{ item.MaxPlayers }}</td>
            <td>{{ item.HasStarted }}</td>
            <td>
              <v-icon v-if="item.HasPassword">mdi-lock</v-icon>
            </td>
          </tr>
        </tbody>
    </template>
    </v-data-table>
    <v-btn test-id="game-list-refresh" @click.native="getGameList" color="primary">
      <v-icon>mdi-refresh</v-icon>
    </v-btn>
  </v-card>
</template>

<script>
import unoService from "../services/unoService";
export default {
  name: "GameList",
  data: () => {
    return {
      headers: [
        {
          text: "Game ID",
          value: "Name"
        },
        {
          text: "Cur Players",
          value: "NumPlayers"
        },
        {
          text: "Max Players",
          value: "MaxPlayers"
        },
        {
          text: "In Progress",
          value: "HasStarted"
        },
        {
          text: "Password",
          value: "HasPassword"
        },
      ],
      selected: null,
      game_list: [],
      sim: true // Only true while debugging
    };
  },
  created: function () {
      this.getGameList();
  },
  methods: {
    async getGameList() {
      let res = await unoService.getGameList();
      this.game_list = res.data;
      this.selected = null;
    },
    selectItem (item) {
      //console.log('Item selected: ' + item.Name)
      if (item === this.selected)
        this.selected = null;
      else
        this.selected = item;
    }
  },
  watch: {
    selected: {
      handler: function() {
        // TODO: Change this to a joinable game ID later
        var id = this.selected ? this.selected.Name : "";
        this.$emit('selectedGame', id);
      }
    }
  }
};
</script>

<style scoped>
.selected {
    background-color: #8c8c8c !important;
    font-weight: bold;
}
</style>