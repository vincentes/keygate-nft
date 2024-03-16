"use client";

import { Button } from "antd";
import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter();

  return (
    <main className="flex min-h-screen flex-col p-12 space-y-16">
      <h1 className="text-5xl font-bold">Keygate</h1>

      <div>
        <Button
          onClick={() => {
            router.push("/dashboard");
          }}
        >
          Start building
        </Button>
      </div>
    </main>
  );
}
