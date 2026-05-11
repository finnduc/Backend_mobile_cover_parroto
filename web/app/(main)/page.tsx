import { Button } from "@/components/ui/button"
import Link from "next/link"
import { ROUTES } from "@/lib/routes"

export default function Page() {
  return (
    <div className="flex min-h-svh flex-col">
      <main className="flex-1">
        <section className="flex flex-col items-center justify-center px-6 py-24 text-center md:py-32">
          <h1 className="text-4xl font-bold tracking-tight md:text-6xl">
            Learn languages with <span className="text-primary">Engflix</span>
          </h1>
          <p className="mt-4 max-w-md text-muted-foreground md:text-lg">
            Practice speaking, build vocabulary, and track your progress with interactive lessons.
          </p>
          <div className="mt-8">
            <Button asChild size="lg">
              <Link href={ROUTES.AUTH.REGISTER}>Get Started</Link>
            </Button>
          </div>
        </section>
      </main>
    </div>
  )
}
