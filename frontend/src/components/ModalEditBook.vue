<script setup>
import { ref } from 'vue';
import Dialog from 'primevue/dialog';
import Editor from 'primevue/editor';
import Textarea from 'primevue/textarea';
import ConfirmDialog from 'primevue/confirmdialog';
import { useConfirm } from "primevue/useconfirm";
import Toast from 'primevue/toast';
import { useToast } from "primevue/usetoast";
import router from "@/router";
import apiClient from "@/services/BooksService.js";
const emit = defineEmits(['bookUpdated', 'bookDeleted']);

const props = defineProps({
    book: {
        type: Object,
        required: true,
    }
});


const confirm = useConfirm();
const toast = useToast()

const visible = ref(false);
const description = ref(props.book.description);
const memo = ref(props.book.memo);

async function updateBook() {
    try {
        await apiClient.put(`/books/${props.book.id}`, {
            description: description.value,
            memo: memo.value,
            isReading: props.book.isReading,
        });
        emit('bookUpdated')
        toggleEditBook();
    } catch (err) {
        console.log(err);
    }
}

async function deleteBook() {
    try {
        await apiClient.delete(`/books/${props.book.id}`);
        emit('bookDeleted')
    } catch (err) {
        console.log(err);
    }
}

function toggleEditBook() {
    visible.value = !visible.value;
}

const confirmDelete = () => {
    confirm.require({
        message: 'Do you want to delete this book?',
        header: 'Danger Zone',
        icon: 'pi pi-info-circle',
        rejectLabel: 'Cancel',
        acceptLabel: 'Delete',
        rejectClass: 'p-button-secondary p-button-outlined',
        acceptClass: 'p-button-danger',
        accept: async () => {
            deleteBook()
            toggleEditBook()
            await router.push({ name: 'library' });

            toast.add({
                severity: 'info',
                summary: 'Confirmed',
                detail: 'Book deleted',
                life: 3000
            });
        },
        reject: () => { }
    });
}
</script>

<template>
    <ConfirmDialog></ConfirmDialog>
    <Toast></Toast>
    <Button label="edit" severity="secondary" rounded @click="toggleEditBook"></Button>
    <Dialog v-model:visible="visible" modal header="Edit book">
        <span class="p-text-secondary block mb-5"></span>
        <form action="post" @submit.prevent="updateBook">
            <div class="flex align-items-center gap-3 mb-3">
                <label for="description" class="font-semibold w-6rem">Description</label>
                <Textarea v-model="description" id="description" class="flex-auto" rows="8" />
            </div>
            <div class="flex align-items-center gap-3 mb-5">
                <label for="memo" class="font-semibold w-6rem">Memo</label>
                <Editor v-model="memo" id="memo" editorStyle="height: 320px" />
            </div>
            <Button type="button" severity="danger" class="delete-book" @click="confirmDelete()">Delete</Button>
            <div class="flex justify-content-end gap-2">
                <Button type="button" severity="secondary" @click="visible = false">Cancel</Button>
                <Button type="submit">Update</Button>
            </div>
        </form>
    </Dialog>
</template>

<style>
Button {
    margin-bottom: 10px;
}

h4 {
    margin: 0;
    padding: 0;
}

.book-info-block {
    display: block;
    width: 100%;
}

.book-info span {
    font-size: small;
}

.auto-complete-input {
    padding-bottom: 10px;
    border-bottom: 1px solid #d4d4d4;
}

.delete-book {
    float: left;
}
</style>