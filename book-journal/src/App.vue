<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div id="layout">
    <header>
    <div class="wrapper">

      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/library">Library</RouterLink>
      </nav>
    </div>

    </header>
    <RouterView />
  </div>
</template>

<style scoped>
</style>
