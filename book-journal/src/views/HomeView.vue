<script setup>
import Book from "@/components/Book.vue"
import apiClient from "@/services/EventService.js"
import {ref, onMounted} from 'vue'

const books = ref(null)

onMounted(() => {
    apiClient.get("/books")
      .then((res) => {
        books.value = res.data.result
        console.log('books:', res.data)})
      .catch((error) => {
        console.log(error) })
})
</script>

<template>
  <main>
    <h2>Currently Reading</h2>
    <div class="books">
      <Book v-for="book in books" :key="book.id" :book="book"/>
    </div>
    <h2>To Read</h2>
    <div class="books">
      <Book v-for="book in books" :key="book.id" :book="book"/>
      <div class="queue">Queue</div>
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
  flex-direction:row;
  align-items: stretch;
  flex-wrap: wrap;
  justify-content: space-evenly;
}

.books:hover {
  cursor: pointer;
}

.queue {
  width: 150px;
  height: 150px;
  border: 1px solid #39495c;
  padding: 10px;
  font-size: 20px;
  font-weight: 700;
  text-align: center;
  line-height: 130px;
}
.queue:hover {
  transform: scale(1.01);
  box-shadow: 0 3px 12px 0 rgba(0, 0, 0, 0.2);
}
</style>
