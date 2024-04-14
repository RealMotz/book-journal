<script setup>
import Dialog from 'primevue/dialog';
import Editor from 'primevue/editor';
import apiClient from "@/services/BooksService.js"
import { ref } from 'vue';

const emit = defineEmits(['bookCreated'])
const visible = ref(false)
const name = ref('')
const description = ref('')
const memo = ref('')

async function createBook() {
    try {
        await apiClient.post("/books", {
            name: name.value,
            description: description.value,
            memo: memo.value,
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
                <label for="name" class="font-semibold w-6rem">Name</label>
                <InputText v-model="name" id="name" class="flex-auto" autocomplete="off" />
            </div>
            <div class="flex align-items-center gap-3 mb-3">
                <label for="description" class="font-semibold w-6rem">Description</label>
                <InputText v-model="description" id="description" class="flex-auto" autocomplete="off" />
            </div>
            <div class="flex align-items-center gap-3 mb-5">
                <label for="memo" class="font-semibold w-6rem">Memo</label>
                <Editor v-model="memo" id="memo" editorStyle="height: 320px" />
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