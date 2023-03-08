export default function useEmoji(emojis: string) {
  return function (idx: number) {
    return String.fromCodePoint(emojis.codePointAt(idx)!)
  }
}
