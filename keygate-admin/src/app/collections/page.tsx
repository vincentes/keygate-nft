"use client";

import { Button, Image, Table, notification } from "antd";
import { ColumnProps } from "antd/es/table";
import { useRouter } from "next/navigation";
import { Collection } from "../data/collection";
import { useEffect, useState } from "react";
import { DeleteOutlined } from "@ant-design/icons";

type CollectionRecord = {
  key: string;
  image: string;
  name: string;
  id: string;
  blockchain: string;
};

export default function Templates() {
  const router = useRouter();
  const [dataSource, setDataSource] = useState<CollectionRecord[]>([]);
  const [api, contextHolder] = notification.useNotification();

  const columns: ColumnProps<CollectionRecord>[] = [
    {
      title: "PREVIEW",
      dataIndex: "image",
      key: "image",
      onCell: (record: CollectionRecord) => {
        return {
          colSpan: 1,
        };
      },
      render: (_: unknown, record: CollectionRecord) => {
        return <Image src={record.image} width={64} height={64} />;
      },
    },
    {
      title: "NAME",
      dataIndex: "name",
      key: "name",
    },
    {
      title: "COLLECTION ID",
      dataIndex: "id",
      key: "id",
      render: (id: string) => {
        return id.substring(0, 23) + "...";
      },
    },
    {
      title: "BLOCKCHAIN",
      dataIndex: "blockchain",
      key: "blockchain",
    },
    {
      title: "ACTION",
      key: "action",
      render: (_: unknown, record: CollectionRecord) => (
        <DeleteOutlined
          onClick={() => handleDelete(record.id)}
          style={{ color: "red", fontSize: "18px" }}
        />
      ),
    },
  ];

  useEffect(() => {
    fetchCollections();
  }, []);

  const fetchCollections = async () => {
    try {
      const response = await fetch("http://localhost:8080/collections");
      const data = await response.json();

      if (data.status === "success") {
        const collections: Collection[] = data.data;
        const collectionRecords: CollectionRecord[] = collections.map(
          (collection) => ({
            key: collection.id,
            image: collection.image,
            name: collection.name,
            id: collection.id,
            blockchain: "Polygon",
          })
        );
        setDataSource(collectionRecords);
      } else {
        api.error({
          message: "Failed to fetch collections",
          description: "Please try again later.",
          placement: "bottom",
        });
      }
    } catch (error) {
      api.error({
        message: "Error fetching collections",
        description: "Please try again later.",
        placement: "bottom",
      });
    }
  };

  const handleDelete = async (id: string) => {
    try {
      const response = await fetch(`http://localhost:8080/collections/${id}`, {
        method: "DELETE",
      });

      if (response.ok) {
        setDataSource((prevDataSource) =>
          prevDataSource.filter((record) => record.id !== id)
        );
        api.success({
          message: "Collection deleted",
          description: "The collection has been successfully deleted.",
          placement: "bottom",
        });
      } else {
        api.error({
          message: "Failed to delete collection",
          description: "Please try again later.",
          placement: "bottom",
        });
      }
    } catch (error) {
      api.error({
        message: "Error deleting collection",
        description: "Please try again later.",
        placement: "bottom",
      });
    }
  };

  return (
    <>
      {contextHolder}
      <main className="flex min-h-screen flex-col items-left p-16">
        <h1 className="text-3xl font-bold">Collections</h1>
        <div className="flex flex-row justify-between">
          <p className="text-lg">
            Create ERC1155 and ERC721 collections to mint NFTs.
          </p>
          <Button
            type="primary"
            className="mt-4"
            onClick={() => router.push("/collections/create/details")}
          >
            New collection
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
