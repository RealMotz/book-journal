<script setup>
import Book from "@/components/Book.vue"
import {ref} from 'vue'

const books = ref([
  {
    id: 1,
    title: "Awesome Book"
  },
  {
    id: 2,
    title: "Awesome Book 2"
  },
  {
    id: 3,
    title: "Awesome Book 3"
  },
  {
    id: 4,
    title: "Awesome Book 4"
  }
])
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