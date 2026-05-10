import { AppSidebar } from "@/components/common/Sidebar"

export default function MainLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex min-h-svh flex-col">
      <AppSidebar />
      <main className="flex-1 p-6">
        {children}
      </main>
    </div>
  )
}
