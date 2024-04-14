"use client";

import { CloseOutlined, UploadOutlined } from "@ant-design/icons";
import { Button, Col, Form, Image, Input, Progress, Row, Upload } from "antd";
import FormItem from "antd/es/form/FormItem";
import TextArea from "antd/es/input/TextArea";
import Column from "antd/es/table/Column";
import { useRouter } from "next/navigation";

export default function Page() {
  const router = useRouter();
  return (
    <main className="flex min-h-screen min-w-screen flex-col items-left">
      <div className="pl-8 pt-8" onClick={() => router.back()}>
        <CloseOutlined
          className="hover:cursor-pointer"
          style={{ fontSize: 20, fontStyle: "bold" }}
        />
      </div>
      <div className="p-16 ">
        <h1 className="text-3xl font-bold">
          Detail the collection&apos;s information
        </h1>
        <p className="text-lg">
          There&apos;s no need to overthink this. This data is editable in
          future steps.
        </p>
        <div className="">
          <Form layout="vertical">
            <FormItem label="Cover Image" name="image">
              <div className="flex flex-row items-center w-64 h-64 hover:cursor-pointer border-2 rounded-xl border border-dashed border-black">
                <div className="w-full text-center">
                  <UploadOutlined style={{ fontSize: 30 }} />
                </div>
              </div>
            </FormItem>
            <FormItem label="Name" name="name">
              <Input />
            </FormItem>
            <FormItem label="Description" name="description">
              <TextArea />
            </FormItem>
            <FormItem>
              <Button type="primary">Submit</Button>
            </FormItem>
          </Form>
        </div>
      </div>
      <Progress
        gapPosition="right"
        percent={20}
        status="active"
        className="fixed bottom-[0px]"
      />
    </main>
  );
}
