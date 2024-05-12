<script setup>
import { ref, onMounted } from 'vue'
import { MdPreview } from 'md-editor-v3';
import apiClient from "@/services/BooksService.js"
import Divider from 'primevue/divider';
import EditBookModal from "@/components/ModalEditBook.vue";
import 'md-editor-v3/lib/style.css';

const book = ref(null)
const props = defineProps({
    id: {
        required: true,
    }
})

onMounted(async () => {
    fetchBookData()
})

const fetchBookData = async function () {
    try {
        const response = await apiClient.get(`/books/${props.id}`);
        book.value = response.data;
    } catch (err) {
        console.log(err);
    }
}

async function toggleReading() {
    if (book.value.isReading) {
        book.value.isReading = false;
        book.value.read = true;
    } else {
        book.value.isReading = true;
    }

    try {
        await apiClient.put(`/books/${book.value.id}`, book.value);
    } catch (err) {
        console.log(err);
    }
}

</script>

<template>
    <main>
        <div class="book-info" v-if="book">
            <img class="book-image" :src="book.thumbnail" alt="Image" />
            <span>
                <h1>{{ book.title }}</h1>
                <div class="options">
                    <EditBookModal :book="book" @book-updated="fetchBookData" />
                    <Button :label="book.isReading ? 'Finish Reading' : book.read ? 'Read' : 'Start Reading'"
                        :severity="book.isReading ? 'danger' : 'contrast'" rounded @click="toggleReading"
                        :disabled="book.read" />
                </div>
            </span>
            <div>{{ book.subtitle }}</div>
            <div>{{ book.authors }}</div>
            <div>{{ book.published_date }}</div>
            <div><a :href="book.info_link" target="_blank">Google book info page</a></div>
            <Divider />

            <h2>Description</h2>
            <p>{{ book.description }}</p>
            <Divider />

            <h2>Notes</h2>
            <MdPreview v-model="book.memo" scrollAuto="false" />
        </div>
    </main>
</template>

<style scoped>
img {
    width: 100px;
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

.md-editor {
    height: 100%;
}
</style>