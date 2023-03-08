<template>
  <div class="gomon-commenteditor-container">
    <div class="gomon-header">
      <input
        type="text"
        class="gomon-input gomon-nickname"
        placeholder="昵称"
      />
      <input type="text" class="gomon-input gomon-email" placeholder="Email" />
    </div>
    <div class="gomon-body">
      <textarea
        class="gomon-textarea"
        style="resize: none"
        v-model="rawContent"
        placeholder="支持Markdown语法"
      ></textarea>
    </div>

    <div class="gomon-emoji-container" v-show="showEmoji">
      <!-- a unicode emoji takes 4 bytes -->
      <span
        v-for="idx in Array.from(
          Array(RegisteredEmojis.length),
          (_, i) => i
        ).filter((v) => v % 2 == 0)"
        @click="appendEmoji(idx)"
        class="gomon-emoji-item"
        >{{ useEmoji(idx) }}</span
      >
    </div>

    <div class="gomon-footer">
      <div class="leftalign">
        <button class="button" @click="showEmoji = !showEmoji">表情</button>
        <button class="button">预览MD</button>
      </div>
      <div class="rightalign">
        <button class="button button-primary">发送评论</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import _useEmoji from '../hooks/useEmoji'

const RegisteredEmojis = '😀😁😅😂👀😍😎😏😟😭✌️👍👎🎉'
const useEmoji = _useEmoji(RegisteredEmojis)

export default defineComponent({
  setup(props) {
    let rawContent = ref('')
    let showEmoji = ref(false)

    function appendEmoji(idx: number) {
      rawContent.value += useEmoji(idx)
    }

    return {
      rawContent,
      RegisteredEmojis,
      useEmoji,
      appendEmoji,
      showEmoji,
    }
  },
})
</script>

<style scoped lang="less">
.gomon-commenteditor-container {
  font-size: 16px;
  width: 100%;
  border: 1px solid #ddd;
  box-shadow: 1px 1px 1px #ddd;
  border-radius: 5px;
  padding: 1rem;
  box-sizing: border-box;
}
.gomon-input {
  padding: 3px;
  outline: none;
  background: transparent;
  border: transparent;
  font-size: 1rem;
}
.gomon-textarea {
  margin-top: 1px;
  padding: 0.5rem;
  height: 5rem;
  width: 100%;
  box-sizing: border-box;
  border: 1px dotted #ccc;
}
.button {
  font-size: 1rem;
  border: none;
  outline: none;
  padding: 0.5rem;
  border-radius: 5px;
  margin: 3px 5px 5px 5px;
  cursor: pointer;
}
.button-primary {
  color: #fff;
  background-color: rgb(0, 131, 255);
}
.gomon-body {
  margin: 0.5rFem auto;
}
.gomon-emoji-container {
  display: flex;
  border: 1px solid #ddd;
  width: fit-content;
  max-width: 600px; // TODO
  transition: all 1s ease-in-out;
  .gomon-emoji-item {
    display: inline;
    width: fit-content;
    cursor: pointer;
    margin: 5px;
  }
}
.gomon-footer {
  display: flex;
  position: relative;
  .rightalign {
    position: absolute;
    right: 0.5rem;
  }
}
</style>
