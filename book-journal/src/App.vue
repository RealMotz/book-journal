<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div id="layout">
    <header>
      <div class="wrapper">

        <nav>
          <RouterLink :to="{ name: 'home' }">Home</RouterLink>
          <RouterLink :to="{ name: 'library' }">Library</RouterLink>
        </nav>
      </div>
    </header>

    <RouterView />
  </div>
</template>
