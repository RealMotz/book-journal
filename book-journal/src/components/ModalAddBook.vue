<script setup>
import Dialog from 'primevue/dialog';
import Editor from 'primevue/editor';
import { ref } from 'vue';

const visible = ref(false)
const value = ref('')
</script>

<template>
    <Button type="button" @click="visible = true">Add book</Button>
    <Dialog v-model:visible="visible" modal header="Add a book">
        <span class="p-text-secondary block mb-5"></span>
        <div class="flex align-items-center gap-3 mb-3">
            <label for="username" class="font-semibold w-6rem">Title</label>
            <InputText id="username" class="flex-auto" autocomplete="off" />
        </div>
        <div class="flex align-items-center gap-3 mb-5">
            <label for="email" class="font-semibold w-6rem">Notes</label>
            <Editor v-model="value" editorStyle="height: 320px" />
        </div>
        <div class="flex justify-content-end gap-2">
            <Button type="button" label="Cancel" severity="secondary" @click="visible = false"></Button>
            <Button type="button" label="Create" @click="visible = false"></Button>
        </div>
    </Dialog>
</template>

<style scoped>
Button {
    margin-bottom: 10px;
}
</style>