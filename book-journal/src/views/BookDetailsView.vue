<script setup>
import { ref, onMounted } from 'vue'
import apiClient from "@/services/BooksService.js"

const props = defineProps({
    id: {
        required: true,
    }
})
const book = ref(null)

onMounted(async () => {
    try {
        const response = await apiClient.get(`/books/${props.id}`);
        book.value = response.data;
    } catch (err) {
        console.log(err);
    }
})
</script>
<template>
    <div v-if="book">
        <h1>Book View</h1>
        <div class="book-info">
            <p>Name: {{ book.name }}</p>
            <p>Description: {{ book.description }}</p>
        </div>
    </div>
</template>