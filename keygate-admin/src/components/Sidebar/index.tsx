"use client";

import { UserButton } from "@clerk/nextjs";
import { Divider, Menu, MenuProps } from "antd";
import Sider from "antd/es/layout/Sider";
import React from "react";
import { usePathname, useRouter } from "next/navigation";

const items: MenuProps["items"] = [
  {
    label: "Dashboard",
    key: "/dashboard",
  },
  {
    label: "Tokens",
    key: "/tokens",
  },
];

export default function Sidebar() {
  const pathname = usePathname();
  const router = useRouter();

  return (
    <Sider
      className="pt-5 pl-2 h-full overflow-clip"
      theme="light"
      width={300}
      breakpoint="lg"
      collapsedWidth="0"
      collapsed={pathname.includes("create")}
    >
      <div className="flex flex-row pl-2 items-center space-x-4">
        <UserButton />
        <div className="bold">Keygate</div>
      </div>
      <Divider />
      <Menu
        className="mt-5"
        mode="inline"
        defaultSelectedKeys={["1"]}
        defaultOpenKeys={["sub1"]}
        selectedKeys={[pathname]}
        onSelect={(e) => {
          router.push(e.key);
        }}
        style={{ height: "100%", borderRight: 0 }}
        items={items}
      />
    </Sider>
  );
}
