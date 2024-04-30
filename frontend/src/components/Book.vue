<script setup>
import Card from 'primevue/card';

const props = defineProps({
  book: {
    type: Object,
    required: true,
  },
})

const shortenStr = function (str) {
  return str.length > 40 ? `${str.slice(0, 40)}...` : str;
}
</script>

<template>
  <div class="book">
    <RouterLink :to="{ name: 'book-details', params: { id: book.id } }">
      <Card>
        <template #header>
          <img alt="user header" :src="book.thumbnail" />
        </template>
        <template #title>{{ shortenStr(book.title) }}</template>
        <template #subtitle>
          <p class="m-0">{{ shortenStr(book.subtitle) }}</p>
        </template>
        <template #content>
          <p class="m-0">{{ book.authors }}</p>
        </template>
      </Card>
    </RouterLink>
  </div>
</template>

<style>
.book {
  text-decoration: none !important;
  color: black;
  width: 200px;
  margin-bottom: 18px;
  padding: 10px 10px 0 10px;
}

a {
  all: unset;
  cursor: pointer;
}

.p-card:hover {
  transform: scale(1.01);
  box-shadow: 0 3px 12px 0 rgba(0, 0, 0, 0.2);
}

.p-card-body {
  padding: 10px;
  flex: 150px;
  position: relative;
}

.p-card-header {
  text-align: center;
}

.p-card-title {
  font-size: medium;
  line-height: normal;
  max-height: 100px;
}

.p-card-subtitle {
  font-size: small;
}

.p-card-subtitle p {
  text-wrap: pretty;
}

.p-card-content {
  position: absolute;
  bottom: 10px;
}
</style>
