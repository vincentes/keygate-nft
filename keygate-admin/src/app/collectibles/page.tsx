"use client";

import { Button, Image, Table } from "antd";
import { ColumnProps } from "antd/es/table";
import { useRouter } from "next/navigation";

type Template = {
  title: string;
  image: string;
};

type TemplateRecord = {
  key: string;
  image: string;
  name: string;
  id: string;
  blockchain: string;
};

const dataSource = [
  {
    image: "https://via.placeholder.com/64",
    key: "1",
    name: "John Brown",
    id: "5748e143-1cbd-453e-a40e-87a573ed60d3",
    blockchain: "Polygon",
  },
  {
    image: "https://via.placeholder.com/64",
    key: "2",
    name: "John Brown 2",
    id: "5748e143-1cbd-453e-a40e-87a573ed60d3",
    blockchain: "Polygon",
  },
];

const renderImage = async (image: string) => {
  return <Image src={image} />;
};

const columns: ColumnProps<TemplateRecord>[] = [
  {
    title: "PREVIEW",
    dataIndex: "image",
    key: "image",
    // if NOT prerender -> do not use renderImage, use undefined
    onCell: (record: TemplateRecord) => {
      return {
        colSpan: 1,
      };
    },
    render: (_: unknown, record: TemplateRecord) => {
      return <Image src={record.image} />;
    },
  },
  {
    title: "NAME",
    dataIndex: "name",
    key: "name",
  },
  {
    title: "COLLECTIBLE ID",
    dataIndex: "id",
    key: "id",
    // limit to only first 20 chars
    render: (id: string) => {
      return id.substring(0, 23) + "...";
    },
  },
  {
    title: "BLOCKCHAIN",
    dataIndex: "blockchain",
    key: "blockchain",
  },
];

export default function Templates() {
  const router = useRouter();

  return (
    <main className="flex min-h-screen flex-col items-left p-16">
      <h1 className="text-3xl font-bold">Collectibles</h1>
      <div className="flex flex-row justify-between">
        <p className="text-lg">
          Design and specify shared properties between digital assets.
        </p>
        <Button
          type="primary"
          className="mt-4"
          onClick={() => router.push("/collectibles/create/details")}
        >
          New collectible
        </Button>
      </div>
      <Table dataSource={dataSource} columns={columns} />
    </main>
  );
}
