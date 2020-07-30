<template>
  <v-container>
    <GameList/>
    <CreateGameFAB/>
  </v-container>
</template>

<script>
// @ is an alias to /src
import GameList from '@/components/GameList'
import CreateGameFAB from '@/components/CreateGameFAB'
export default {
  name: 'Lobby',
  components: {
    GameList,
    CreateGameFAB
  }
}
</script>
