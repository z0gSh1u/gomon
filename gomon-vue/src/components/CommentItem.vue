<template>
  <div class="gomon-item-container">
    <div class="left-block">
      <img :src="gravatarURL" class="avatar" />
    </div>
    <div class="right-block">
      <p class="username" v-text="username"></p>
      <p class="time" v-text="formatDateTime(datetime)"></p>
      <p class="content" v-html="markdownContent"></p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed } from 'vue'
import MarkdownIt from 'markdown-it'
import CryptoJS from 'crypto-js'

export default defineComponent({
  props: {
    username: {
      type: String,
      required: true,
    },
    datetime: {
      type: String,
      required: true,
    },
    content: {
      type: String,
      default: '*xxx*',
    },
    email: {
      type: String,
      default: 'notExist@default.com',
    },
  },
  setup(props) {
    const markdown = MarkdownIt()
    const gravatarURL = computed(
      () => 'https://www.gravatar.com/avatar/' + CryptoJS.MD5(props.email)
    )
    const markdownContent = computed(() => markdown.render(props.content))
    function formatDateTime(datetime: string) {
      return datetime
    }
    return {
      gravatarURL,
      markdownContent,
      formatDateTime
    }
  },
})
</script>

<style scoped lang="less">
.gomon-item-container {
  font-size: 16px;
  width: 100%;
  display: flex;
  border-bottom: 1px dotted #bbb;
  margin-bottom: 4px;
  .left-block {
    .avatar {
      width: 36px;
      height: 36px;
    }
  }
  .right-block {
    margin-left: 1em;
    p {
      margin: 0;
    }
    .username {
      font-weight: bold;
    }
    .time {
      font-size: 0.75rem;
    }
    .content {
      word-wrap: break-word;
      word-break: break-all;
      line-height: 1.25em;
    }
  }
}
</style>
