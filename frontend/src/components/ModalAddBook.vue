<script setup>
import Dialog from 'primevue/dialog';
import Editor from 'primevue/editor';
import AutoComplete from 'primevue/autocomplete';
import apiClient from "@/services/BooksService.js"
import { ref } from 'vue';
import axios from 'axios';

const emit = defineEmits(['bookCreated'])

const visible = ref(false)
const title = ref('')
const description = ref('')
const memo = ref('')

const books = ref([])
const selectedBook = ref({})

async function createBook() {
    if (!Boolean(title.value)) {
        return;
    }

    const data = fetchBookData(selectedBook.value);

    try {
        await apiClient.post("/books", {
            title: title.value,
            subtitle: data.subtitle,
            authors: data.authors,
            publication_date: data.publication_date,
            thumbnail: data.thumbnail,
            info_link: data.info_link,
            description: description.value,
            memo: memo.value,
        });
        emit('bookCreated')
        visible.value = false;
    } catch (error) {
        console.log(error);
    }
}

const fetchBookData = function (bookData) {
    let volumeInfo = bookData.hasOwnProperty("volumeInfo") ? bookData.volumeInfo : null
    return {
        subtitle: volumeInfo ? volumeInfo.subtitle : "",
        authors: volumeInfo && volumeInfo.authors ? volumeInfo.authors.join(",") : "",
        publication_date: volumeInfo ? volumeInfo.publishedDate : "",
        thumbnail: volumeInfo.imageLinks && volumeInfo.imageLinks.thumbnail ?
            volumeInfo.imageLinks.thumbnail : "https://placehold.co/100x80",
        info_link: volumeInfo.infoLink ? volumeInfo.infoLink : ""
    }
}

async function searchBook(event) {
    if (!event.query.trim().length) {
        books.value = [];
        return;
    }

    const query = event.query.toLocaleLowerCase().replace(" ", "+")
    try {
        const res = await axios.get(`https://www.googleapis.com/books/v1/volumes?q=${query}&maxResults=10`);
        const filteredImg = res.data.items.filter((book) => {
            return book.volumeInfo.hasOwnProperty("imageLinks");
        })
        books.value = filteredImg;
    } catch (error) {
        console.log(error);
    }
}

const getAuthors = function (option) {
    const volumeInfo = option.hasOwnProperty("volumeInfo") ? option.volumeInfo : null
    return (volumeInfo && volumeInfo.authors) ? volumeInfo.authors.join(", ") : ""
}

const onItemSelect = () => {
    title.value = selectedBook.value.volumeInfo.title;
    description.value = selectedBook.value.volumeInfo.description;
}
</script>

<template>
    <Button type="button" @click="visible = true">Add book</Button>
    <Dialog v-model:visible="visible" modal header="Add a book">
        <span class="p-text-secondary block mb-5"></span>
        <form action="post" @submit.prevent="createBook">

            <div class="flex align-items-center gap-3 mb-3 auto-complete-input">
                <label for="title" class="font-semibold w-6rem">Search</label>
                <AutoComplete v-model="selectedBook" optionLabel="volumeInfo.title" :suggestions="books"
                    @complete="searchBook" @item-select="onItemSelect" placeholder="Type book title here">
                    <template #option="slotProps">
                        <div class="flex align-options-center">
                            <img :alt="slotProps.option.volumeInfo.title"
                                :src="slotProps.option.volumeInfo.imageLinks.smallThumbnail" class="mr-2"
                                style="width: 50px" />
                            <div class="book-info">
                                <h4 class="book-info-block">{{ slotProps.option.volumeInfo.title }}</h4>
                                <span class="book-info-block">{{ getAuthors(slotProps.option)
                                    }}</span>
                            </div>
                        </div>
                    </template>
                </AutoComplete>
            </div>


            <div class="flex align-items-center gap-3 mb-3">
                <label for="title" class="font-semibold w-6rem">Title</label>
                <InputText v-model="title" id="title" class="flex-auto" autocomplete="off" />
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

/* primevue styles */

.p-autocomplete {
    width: 85%;
}

.p-autocomplete>input {
    width: 100%;
}

.p-autocomplete-panel .p-autocomplete-items {
    min-width: 400px;
    max-width: 630px;
}
</style>