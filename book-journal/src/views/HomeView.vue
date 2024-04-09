<script setup>
import Book from "@/components/Book.vue"
import apiClient from "@/services/BooksService.js"
import { ref, onMounted } from 'vue'

const books = ref(null)

onMounted(() => {
  apiClient.get("/books")
    .then((res) => {
      books.value = res.data
      console.log('books:', res.data)
    })
    .catch((error) => {
      console.log(error)
    })
})
</script>

<template>
  <main>
    <h2>Currently Reading</h2>
    <div class="books">
      <Book v-for="book in books" :key="book.id" :book="book" />
    </div>
  </main>
</template>

<style scoped>
h2 {
  font-size: 20px;
  border-bottom: 1px solid black;
  margin-bottom: 10px;
}

.books {
  display: flex;
  flex-direction: row;
  align-items: stretch;
  flex-wrap: wrap;
  justify-content: space-evenly;
}

.books:hover {
  cursor: pointer;
}
</style>
@/services/BooksService.js