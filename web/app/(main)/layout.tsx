import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar"
import { AppSidebar } from "@/components/common/Sidebar"
import { Topbar } from "@/components/common/Topbar"
import { ChatWidget } from "@/features/chat"

export default function MainLayout({ children }: { children: React.ReactNode }) {
  return (
    <SidebarProvider>
      <AppSidebar />
      <main className="flex flex-1 flex-col">
        <Topbar />
        <div className="flex-1 p-6">
          {children}
        </div>
      </main>
      <ChatWidget />
    </SidebarProvider>
  )
}
