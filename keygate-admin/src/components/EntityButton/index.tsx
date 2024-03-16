"use client";
import React from "react";
import { useRouter } from "next/navigation";
import { Button } from "antd";

export default function EntityButton({
  name,
  endpoint,
}: Readonly<{
  name: string;
  endpoint: string;
}>) {
  const router = useRouter();

  return (
    <Button
      className="w-32"
      type="primary"
      onClick={() => {
        router.push(endpoint);
      }}
    >
      {name}
    </Button>
  );
}
