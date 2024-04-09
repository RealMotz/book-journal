<script setup>
import Book from "@/components/Book.vue"
import router from "@/router/index.js"
import apiClient from "@/services/BooksService.js"
import Modal from "@/components/ModalAddBook.vue"
import { ref, onMounted } from 'vue'

const books = ref(null)

onMounted(() => {
  apiClient.get("/books")
    .then((res) => {
      books.value = res.data
    })
    .catch((error) => {
      console.log(error)
    })
})

async function addBook() {
  try {
    const newBook = await apiClient.post("/books");
    router.push({
      name: 'book-details',
      params: { id: newBook.data.id }
    })
  } catch (error) {
    console.log(error);
  }
}
</script>

<template>
  <div class="about">
    <h1>Library</h1>
    <Modal />
    <div class="books">
      <Book v-for="book in books" :key="book.id" :book="book" />
    </div>
  </div>
</template>

<style scoped>
button {
  border-style: none;
  padding: 10px;
  border-radius: 10px;
  background-color: #c47e7e;
  color: #fff;
  margin-bottom: 10px;
}

button:active {
  transform: translateY(4px);
  box-shadow: 0 1px #ffffff;
}
</style>
