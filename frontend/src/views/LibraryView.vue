<script setup>
import Book from "@/components/Book.vue";
import apiClient from "@/services/BooksService.js";
import AddBookModal from "@/components/ModalAddBook.vue";
import Toast from 'primevue/toast';
import { ref, onMounted } from 'vue';

const books = ref(null)

onMounted(async () => {
  fetchAllBooks()
})

async function fetchAllBooks() {
  try {
    const response = await apiClient.get("/books");
    books.value = response.data;
  } catch (error) {
    console.log(error);
  }
}
</script>

<template>
  <main>
    <Toast></Toast>
    <div class="about">
      <h1>Library</h1>
      <AddBookModal @book-created="fetchAllBooks" />
      <div class="books">
        <Book v-for="book in books" :key="book.id" :book="book" />
      </div>
    </div>
  </main>
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

.books {
  display: flex;
  align-items: baseline;
}
</style>
