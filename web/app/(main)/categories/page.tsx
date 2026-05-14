import { PageLayout } from "@/components/layouts/PageLayout"
import { getCategories } from "@/features/categories/services/categories-service"
import { ROUTES } from "@/lib/routes"
import Link from "next/link"

export default async function CategoriesPage() {
  const res = await getCategories()
  const categories = res.data ?? []

  return (
    <PageLayout
      title="Chủ đề"
      breadcrumbs={[{ label: "Chủ đề", href: ROUTES.USER.CATEGORIES }]}
    >
      <div className="grid gap-3 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
        {categories.map((category) => (
          <Link
            key={category.id}
            href={`${ROUTES.USER.LESSONS.LIST}?categoryId=${category.id}`}
            className="rounded-xl border bg-card p-4 text-card-foreground shadow-sm transition-colors hover:bg-muted/50"
          >
            <h3 className="font-medium">{category.name}</h3>
          </Link>
        ))}
      </div>
    </PageLayout>
  )
}
