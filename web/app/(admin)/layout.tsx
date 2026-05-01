import { SidebarProvider } from "@/components/ui/sidebar"
import { AdminSidebar } from "@/components/common/AdminSidebar"

export default function AdminLayout({ children }: { children: React.ReactNode }) {
  return (
    <SidebarProvider>
      <AdminSidebar />
      <main className="flex flex-1 flex-col">
        <header className="flex h-14 items-center border-b bg-background px-6">
          <h1 className="text-sm font-medium text-muted-foreground">Admin Dashboard</h1>
        </header>
        <div className="flex-1 p-6">
          {children}
        </div>
      </main>
    </SidebarProvider>
  )
}
