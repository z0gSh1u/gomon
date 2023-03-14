export const CommentEditorConfig = {
  topicId: String,
}
export type CommentEditorConfigT = typeof CommentEditorConfig

export interface CommentItemConfig {}

export interface CommentListConfig {}

export const RegisteredEmojis = '😀😁😅😂👀😍😎😏😟😭✌️👍👎🎉'

export const BaseURL = 'http://127.0.0.1:8080';

export function api(path: string) {
  return BaseURL + '/' + path
}

export interface CommentInfo {
  commentId: string
  topicId: string
  unixTime: string
  username: string
  email: string // gravatar: https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50
  content: string
}
