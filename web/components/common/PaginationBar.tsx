import Link from "next/link"
import {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination"

function buildPageUrl(baseUrl: string, page: number, searchParams: URLSearchParams): string {
  const params = new URLSearchParams(searchParams)
  params.set("page", String(page))
  return `${baseUrl}?${params.toString()}`
}

function getPageNumbers(currentPage: number, totalPages: number): (number | "ellipsis")[] {
  if (totalPages <= 7) {
    return Array.from({ length: totalPages }, (_, i) => i + 1)
  }
  const pages: (number | "ellipsis")[] = [1]
  if (currentPage > 3) pages.push("ellipsis")
  for (let i = Math.max(2, currentPage - 1); i <= Math.min(totalPages - 1, currentPage + 1); i++) {
    pages.push(i)
  }
  if (currentPage < totalPages - 2) pages.push("ellipsis")
  pages.push(totalPages)
  return pages
}

export function PaginationBar({
  currentPage,
  totalPages,
  baseUrl,
  searchParams,
}: {
  currentPage: number
  totalPages: number
  baseUrl: string
  searchParams?: URLSearchParams
}) {
  if (totalPages <= 1) return null

  const params = searchParams ?? new URLSearchParams()

  return (
    <Pagination>
      <PaginationContent>
        <PaginationItem>
          <PaginationPrevious
            href={currentPage <= 1 ? "#" : buildPageUrl(baseUrl, currentPage - 1, params)}
            className={currentPage <= 1 ? "pointer-events-none opacity-50" : ""}
          />
        </PaginationItem>

        {getPageNumbers(currentPage, totalPages).map((p, i) =>
          p === "ellipsis" ? (
            <PaginationItem key={`ellipsis-${i}`}>
              <PaginationEllipsis />
            </PaginationItem>
          ) : (
            <PaginationItem key={p}>
              <PaginationLink
                href={buildPageUrl(baseUrl, p, params)}
                isActive={p === currentPage}
              >
                {p}
              </PaginationLink>
            </PaginationItem>
          )
        )}

        <PaginationItem>
          <PaginationNext
            href={currentPage >= totalPages ? "#" : buildPageUrl(baseUrl, currentPage + 1, params)}
            className={currentPage >= totalPages ? "pointer-events-none opacity-50" : ""}
          />
        </PaginationItem>
      </PaginationContent>
    </Pagination>
  )
}
