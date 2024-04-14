"use client";

import { Button, Table, notification } from "antd";
import { ColumnProps } from "antd/es/table";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { DeleteOutlined } from "@ant-design/icons";

type TokenRecord = {
  key: string;
  name: string;
  id: string;
  description: string;
  address: string;
};

export default function Tokens() {
  const router = useRouter();
  const [dataSource, setDataSource] = useState<TokenRecord[]>([]);
  const [api, contextHolder] = notification.useNotification();

  const columns: ColumnProps<TokenRecord>[] = [
    {
      title: "NAME",
      dataIndex: "name",
      key: "name",
    },
    {
      title: "TOKEN ID",
      dataIndex: "id",
      key: "id",
      render: (id: string) => {
        return id.substring(0, 23) + "...";
      },
    },
    {
      title: "DESCRIPTION",
      dataIndex: "description",
      key: "description",
    },
    {
      title: "ADDRESS",
      dataIndex: "address",
      key: "address",
      render: (address: string) => {
        return (
          <div>
            <Button
              onClick={() =>
                window.open(
                  `https://explorer.solana.com/address/${address}?cluster=devnet`,
                  "_blank"
                )
              }
            >
              {address.substring(0, 23) + "..."}
            </Button>
          </div>
        );
      },
    },
    {
      title: "ACTION",
      key: "action",
      render: (_: unknown, record: TokenRecord) => (
        <DeleteOutlined
          onClick={() => handleDelete(record.id)}
          style={{ color: "red", fontSize: "18px" }}
        />
      ),
    },
  ];

  useEffect(() => {
    fetchTokens();
  }, []);

  const fetchTokens = async () => {
    try {
      const response = await fetch(`${process.env.backendApiUrl}/tokens/`);
      const data = await response.json();

      if (data.status === "success") {
        const tokens: TokenRecord[] = data.data;
        setDataSource(tokens);
      } else {
        api.error({
          message: "Failed to fetch tokens",
          description: "Please try again later.",
          placement: "bottom",
        });
      }
    } catch (error) {
      api.error({
        message: "Error fetching tokens",
        description: "Please try again later.",
        placement: "bottom",
      });
    }
  };

  const handleDelete = async (id: string) => {
    try {
      const response = await fetch(`http://localhost:8080/tokens/${id}`, {
        method: "DELETE",
      });

      if (response.ok) {
        setDataSource((prevDataSource) =>
          prevDataSource.filter((record) => record.id !== id)
        );
        api.success({
          message: "Token deleted",
          description: "The token has been successfully deleted.",
          placement: "bottom",
        });
      } else {
        api.error({
          message: "Failed to delete token",
          description: "Please try again later.",
          placement: "bottom",
        });
      }
    } catch (error) {
      api.error({
        message: "Error deleting token",
        description: "Please try again later.",
        placement: "bottom",
      });
    }
  };

  return (
    <>
      {contextHolder}
      <main className="flex min-h-screen flex-col items-left p-16">
        <h1 className="text-3xl font-bold">Tokens</h1>
        <div className="flex flex-row justify-between">
          <p className="text-lg">Create and manage your tokens.</p>
          <Button
            type="primary"
            className="mt-4"
            onClick={() => router.push("/tokens/create")}
          >
            New token
          </Button>
        </div>
        <Table
          rowClassName={"cursor-pointer"}
          dataSource={dataSource}
          columns={columns}
        />
      </main>
    </>
  );
}
