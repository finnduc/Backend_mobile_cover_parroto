export default function AuthLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex items-center justify-center min-h-screen">
      <main className="flex-1 p-6">{children}</main>
    </div>
  )
}