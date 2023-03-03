<template>
  <div class="gomon-editor border-2 border-r-2 border-gray-500 w-12">
    <div class="gomon-header">
      <input type="text" class="gomon-nickname" />
    </div>
    <div class="gomon-body">
      <textarea
        class="gomon-textarea w-full border-2 border-gray-500"
        style="resize: none"
        v-model="rawContent"
        placeholder="支持Markdown语法"
      ></textarea>
    </div>
    <div class="gomon-footer">
      <button>表情</button>
      <button>预览MD</button>
      <button class="bg-blue-500">发送评论</button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
export default defineComponent({
  setup() {
    const rawContent = ref('')
    return {
      rawContent,
    }
  },
})
</script>

<style scoped></style>
