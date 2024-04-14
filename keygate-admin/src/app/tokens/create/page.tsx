"use client";

import { CloseOutlined } from "@ant-design/icons";
import { Button, Form, Input, notification } from "antd";
import { useRouter } from "next/navigation";
import { useMemo, useState } from "react";
import { TokenCreateContext } from "../../../context/TokenCreate";

const submitToken = async (name: string, description: string) => {
  return fetch("http://localhost:8080/tokens/", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      name,
      description,
    }),
  });
};

  

export default function CreateToken() {
  const router = useRouter();
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [api, contextHolder] = notification.useNotification();

  const submitError = () => {
    api.error({
      message: "Error",
      description: "Could not submit token.",
      placement: "bottom",
    });
  };

  const onFinish = async () => {
    submitToken(name, description)
      .then((response) => {
        if (response.ok) {
          router.push("/tokens");
        } else {
          submitError();
        }
      })
      .catch(() => {
        submitError();
      });
  };

  const contextValue = useMemo(
    () => ({
      name,
      setName,
      description,
      setDescription,
    }),
    [name, description]
  );

  return (
    <>
      <div className="pl-8 pt-8" onClick={() => router.push("/tokens")}>
        <CloseOutlined
          className="hover:cursor-pointer"
          style={{ fontSize: 20, fontStyle: "bold" }}
        />
      </div>
      <TokenCreateContext.Provider value={contextValue}>
        {contextHolder}
        <main className="flex min-h-screen flex-col items-left p-16">
          <h1 className="text-3xl font-bold">Create Token</h1>
          <Form onFinish={onFinish} layout="vertical" className="mt-8">
            <Form.Item
              label="Name"
              name="name"
              rules={[{ required: true, message: "Please enter a token name" }]}
            >
              <Input
                placeholder="Token Name"
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
            </Form.Item>
            <Form.Item
              label="Description"
              name="description"
              rules={[
                { required: true, message: "Please enter a token description" },
              ]}
            >
              <Input.TextArea
                placeholder="Token Description"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
              />
            </Form.Item>
            <Form.Item>
              <Button type="primary" htmlType="submit">
                Create Token
              </Button>
            </Form.Item>
          </Form>
        </main>
      </TokenCreateContext.Provider>
    </>
  );
}
