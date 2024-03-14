<script setup>
import Book from "@/components/Book.vue"
import {ref} from 'vue'

const books = ref([
  {
    id: 1,
    title: "Awesome Book"
  }
])
</script>

<template>
  <main>
    <div class="books">
      <Book v-for="book in books" :key="book.id" :book="book"/>
    </div>
  </main>
</template>

<style scoped>
.books {
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>