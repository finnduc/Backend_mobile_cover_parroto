export interface ChatMessage {
  id: number
  userId: string
  userName?: string
  avatarUrl?: string
  content: string
  createdAt: string
}

export interface ChatHistory {
  messages: ChatMessage[]
  hasMore: boolean
  nextId?: number
}
