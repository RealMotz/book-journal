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

    <h1>Book Journal</h1>
    <RouterView />
  </div>
</template>

<style>
h2 {
  font-size: 20px;
}
</style>
