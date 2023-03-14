<script setup lang="ts">
import CommentEditor from './components/CommentEditor.vue'
import CommentItem from './components/CommentItem.vue'
import CommentList from './components/CommentList.vue'


</script>

<template>
  <CommentEditor></CommentEditor>
  <CommentList :topic-id="'topic01'"></CommentList>
</template>

<style scoped></style>
