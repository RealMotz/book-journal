<script setup>
import { ref, onMounted } from 'vue'
import apiClient from "@/services/BooksService.js"
import Divider from 'primevue/divider';

const book = ref(null)
const props = defineProps({
    id: {
        required: true,
    }
})

onMounted(async () => {
    try {
        const response = await apiClient.get(`/books/${props.id}`);
        book.value = response.data;
    } catch (err) {
        console.log(err);
    }
})

async function toggleReading() {
    book.value.isReading = !book.value.isReading;
    try {
        const response = await apiClient.put(`/books`, book.value);
        console.log(response);
    } catch (err) {
        console.log(err);
    }
}

</script>
<template>
    <div v-if="book">
        <div class="book-info">
            <img class="book-image" src="https://placehold.co/600x400" alt="Image" />
            <span>
                <h1>Title</h1>
                <div class="options">
                    <Button :label="book.isReading ? 'Stop Reading' : 'Start reading'"
                        :severity="book.isReading ? 'danger' : 'contrast'" rounded @click="toggleReading" />
                </div>
            </span>
            <p>{{ book.title }}</p>
            <Divider />

            <h2>Description</h2>
            <p>{{ book.description }}</p>
            <Divider />

            <h2>Notes</h2>
            <div v-html="book.notes"></div>
        </div>
    </div>
</template>

<style scoped>
img {
    width: 100%;
    text-align: center;
}

span {
    position: relative;
}

.options {
    position: absolute;
    right: 0;
    top: 40px;
}

.options button {
    width: 150px;
}
</style>