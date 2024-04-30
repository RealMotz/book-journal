<script setup>
import { ref } from 'vue';
import { MdEditor } from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
const emit = defineEmits(['editorTextChanged']);

const props = defineProps({
    memo: {
        type: String,
        required: true,
    }
});

const memo = ref(props.memo);

const toolbars = [
    'bold',
    'underline',
    'italic',
    '-',
    'title',
    'strikeThrough',
    'quote',
    'unorderedList',
    'orderedList',
    'task',
    '-',
    'codeRow',
    'code',
    'link',
    'image',
    'table',
    'mermaid',
    '=',
    'preview',
    'previewOnly',
];

const onChanged = function (text) {
    emit('editorTextChanged', text)
}
</script>

<template>
    <MdEditor v-model="memo" language="us" :toolbars="toolbars" height="300px" @onChange="onChanged" />
</template>

<style>
.md-editor-footer-item {
    display: none;
}

.md-editor {
    height: 400px;
    width: auto;
}
</style>
