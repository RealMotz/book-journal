<script setup>
import Dialog from 'primevue/dialog';
import Editor from 'primevue/editor';
import apiClient from "@/services/BooksService.js"
import { ref } from 'vue';

const emit = defineEmits(['bookCreated'])
const visible = ref(false)
const title = ref('')
const description = ref('')
const notes = ref('')

async function createBook() {
    try {
        await apiClient.post("/books", {
            title: title.value,
            description: description.value,
            notes: notes.value,
        });
        emit('bookCreated')
        visible.value = false;
    } catch (error) {
        console.log(error);
    }
}
</script>

<template>
    <Button type="button" @click="visible = true">Add book</Button>
    <Dialog v-model:visible="visible" modal header="Add a book">
        <span class="p-text-secondary block mb-5"></span>
        <form action="post" @submit.prevent="createBook">
            <div class="flex align-items-center gap-3 mb-3">
                <label for="title" class="font-semibold w-6rem">Title</label>
                <InputText v-model="title" id="title" class="flex-auto" autocomplete="off" />
            </div>
            <div class="flex align-items-center gap-3 mb-3">
                <label for="description" class="font-semibold w-6rem">Description</label>
                <InputText v-model="description" id="description" class="flex-auto" autocomplete="off" />
            </div>
            <div class="flex align-items-center gap-3 mb-5">
                <label for="email" class="font-semibold w-6rem">Notes</label>
                <Editor v-model="notes" editorStyle="height: 320px" />
            </div>
            <div class="flex justify-content-end gap-2">
                <Button type="button" severity="secondary" @click="visible = false">Cancel</Button>
                <Button type="submit">Create</Button>
            </div>
        </form>
    </Dialog>
</template>

<style scoped>
Button {
    margin-bottom: 10px;
}
</style>