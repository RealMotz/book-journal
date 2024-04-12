<script setup>
import Card from 'primevue/card';

defineProps({
  book: {
    type: Object,
    required: true,
  },
})
</script>

<template>
  <div class="book">
    <RouterLink :to="{ name: 'book-details', params: { id: book.id } }">
      <Card>
        <template #title>{{ book.title }}</template>
        <template #content>
          <p class="m-0">
            {{ book.description }}
          </p>
        </template>
      </Card>
    </RouterLink>
  </div>
</template>

<style scoped>
.book {
  text-decoration: none !important;
  color: black;
  width: 200px;
  margin-bottom: 18px;
  padding: 10px;
}

.p-card:hover {
  transform: scale(1.01);
  box-shadow: 0 3px 12px 0 rgba(0, 0, 0, 0.2);
}

a {
  all: unset;
  cursor: pointer;
}
</style>
