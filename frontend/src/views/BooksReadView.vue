<script setup>
import Book from "@/components/Book.vue"
import apiClient from "@/services/BooksService.js"
import { ref, onMounted } from 'vue'

const books = ref(null)

onMounted(async () => {
    try {
        const response = await apiClient.get("/books", {
            params: {
                read: true
            }
        });
        books.value = response.data
    } catch (error) {
        console.log(error);
    }
})
</script>

<template>
    <main>
        <h1>Books Read</h1>
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
    justify-content: flex-start;
}

.books:hover {
    cursor: pointer;
}
</style>