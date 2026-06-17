import { normalizeWord } from "@/lib/utils"

export function computeWordDiff(correct: string[], input: string[]) {
  const result: { word: string; status: "correct" | "wrong" }[] = []

  for (let i = 0; i < input.length; i++) {
    const c = normalizeWord(correct[i] ?? "")
    const u = normalizeWord(input[i] ?? "")

    result.push({
      word: input[i],
      status: c === u ? "correct" : "wrong",
    })
  }

  return result
}
