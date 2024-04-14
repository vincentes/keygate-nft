/** @type {import('next').NextConfig} */
const nextConfig = {
    env: {
        backendApiUrl: process.env.BACKEND_API_URL,
    }
};

export default nextConfig;
