<template>
  <div class="gomon-comment-list">
    <CommentItem
      v-for="comment in comments"
      :username="comment['username']"
      :content="comment['content']"
      :datetime="comment['unixTime']"
    ></CommentItem>
    <button class="gomon-btn">Load More</button>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import CommentItem from './CommentItem.vue'
import { CommentInfo, api } from '../utils/gomon-types'
import { useAxios } from '../hooks/useAxios'

export default defineComponent({
  components: { CommentItem },
  props: {
    topicId: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    let comments = ref<CommentInfo[]>([])
    const { run, data } = useAxios(api(`get_comments/${props.topicId}`), {
      timeout: 5000,
    })
    run().then(() => {
      comments.value = (data.value as any)['payload']
    })
    return {
      comments,
    }
  },
})
</script>

<style scoped lang="less">
.gomon-btn {
  padding: 0.5em 1em;
  background: #fff;
  border: 1px solid #ccc;
  border-radius: 0.5em;
  cursor: pointer;
  color: #222;
  transition-duration: 0.3s;
  &:hover {
    color: #fff;
    background-color: rgb(38, 148, 250);
  }
}
</style>
