/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites() {
    return [
      {
        source: "/api/chat/events",
        destination: `${process.env.API_URL}/chat/events`,
      },
    ]
  },
}

export default nextConfig
