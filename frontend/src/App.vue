<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div id="layout">
    <header>
      <div class="wrapper">
        <nav>
          <RouterLink :to="{ name: 'home' }">Reading</RouterLink>
          <RouterLink :to="{ name: 'books-read' }">Read</RouterLink>
          <RouterLink :to="{ name: 'library' }">Library</RouterLink>
        </nav>
      </div>
    </header>

    <RouterView />
  </div>
</template>

<style scoped>
a {
  text-decoration: none;
  color: hsla(160, 100%, 37%, 1);
  transition: 0.4s;
  padding: 3px;
}

.wrapper {
  margin-bottom: 20px;
}
</style>