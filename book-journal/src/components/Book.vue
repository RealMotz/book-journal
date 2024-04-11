<script setup>
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
      <h2>{{ book.title }}</h2>
      <p>{{ book.description }}</p>
    </RouterLink>
  </div>
</template>

<style scoped>
.book {
  width: 200px;
  border: 1px solid #39495c;
  margin-bottom: 18px;
  padding: 10px;
}

.book:hover {
  transform: scale(1.01);
  box-shadow: 0 3px 12px 0 rgba(0, 0, 0, 0.2);
}
</style>
