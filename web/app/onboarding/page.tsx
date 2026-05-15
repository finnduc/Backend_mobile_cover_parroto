"use client"

import { useState } from "react"
import { useUser } from "@clerk/nextjs"
import { useRouter } from "next/navigation"
import { Button } from "@/components/ui/button"
import { completeSignUp } from "@/features/auth/services/auth.action"
import { Loader2, CheckCircle2 } from "lucide-react"

export default function OnboardingPage() {
  const { user, isLoaded } = useUser()
  const router = useRouter()
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const handleComplete = async () => {
    setLoading(true)
    setError(null)

    const res = await completeSignUp()

    if (res.error) {
      setError(res.error.message || "Có lỗi xảy ra, vui lòng thử lại.")
      setLoading(false)
      return
    }

    // Reload the user object and session token so the JWT picks up
    // the new role metadata written by the backend.
    // See: https://clerk.com/docs/guides/sessions/force-token-refresh
    await user?.reload()

    // Force a full navigation so the middleware re-runs with the updated claims.
    window.location.href = "/"
  }

  if (!isLoaded) {
    return (
      <main className="flex min-h-screen items-center justify-center">
        <Loader2 className="size-8 animate-spin text-muted-foreground" />
      </main>
    )
  }

  return (
    <main className="flex min-h-screen items-center justify-center bg-muted/30 px-4">
      <div className="w-full max-w-md space-y-6 rounded-2xl border bg-background p-8 shadow-sm">
        <div className="space-y-2 text-center">
          <h1 className="text-2xl font-semibold tracking-tight">
            Chào mừng bạn đến với Parroto!
          </h1>
          <p className="text-sm text-muted-foreground">
            Chỉ cần một bước nữa để hoàn tất đăng ký và bắt đầu học.
          </p>
        </div>

        <div className="flex items-center gap-4 rounded-xl border bg-muted/50 p-4">
          {user?.imageUrl ? (
            <img
              src={user.imageUrl}
              alt="avatar"
              className="size-12 rounded-full object-cover"
            />
          ) : (
            <div className="flex size-12 items-center justify-center rounded-full bg-primary/10 text-primary font-semibold">
              {user?.firstName?.charAt(0) || user?.emailAddresses[0]?.emailAddress?.charAt(0) || "?"}
            </div>
          )}
          <div>
            <p className="text-sm font-medium">
              {user?.fullName || user?.firstName || "Ngườı dùng"}
            </p>
            <p className="text-xs text-muted-foreground">
              {user?.emailAddresses[0]?.emailAddress}
            </p>
          </div>
        </div>

        <ul className="space-y-3 text-sm text-muted-foreground">
          <li className="flex items-start gap-3">
            <CheckCircle2 className="mt-0.5 size-4 shrink-0 text-primary" />
            Truy cập đầy đủ bài học và lộ trình học tập
          </li>
          <li className="flex items-start gap-3">
            <CheckCircle2 className="mt-0.5 size-4 shrink-0 text-primary" />
            Lưu lịch sử học tập và đánh dấu bài yêu thích
          </li>
          <li className="flex items-start gap-3">
            <CheckCircle2 className="mt-0.5 size-4 shrink-0 text-primary" />
            Theo dõi tiến độ và nhận gợi ý cá nhân hóa
          </li>
        </ul>

        {error && (
          <p className="rounded-lg bg-destructive/10 px-3 py-2 text-sm text-destructive">
            {error}
          </p>
        )}

        <Button
          className="w-full"
          size="lg"
          onClick={handleComplete}
          disabled={loading}
        >
          {loading ? (
            <>
              <Loader2 className="mr-2 size-4 animate-spin" />
              Đang xử lý...
            </>
          ) : (
            "Hoàn tất đăng ký"
          )}
        </Button>
      </div>
    </main>
  )
}
