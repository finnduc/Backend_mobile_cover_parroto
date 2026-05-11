import { AuthAwareNav } from "@/components/common/AuthAwareNav"

export default function MainLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex min-h-svh flex-col">
      <AuthAwareNav />
      <main className="flex-1 p-6">
        {children}
      </main>
    </div>
  )
}
