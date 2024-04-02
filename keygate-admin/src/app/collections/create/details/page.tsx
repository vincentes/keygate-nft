"use client";

import {
  FileImageOutlined,
  PlusOutlined,
  UploadOutlined,
} from "@ant-design/icons";
import { Button, Form, Input, Upload } from "antd";
import FormItem from "antd/es/form/FormItem";
import TextArea from "antd/es/input/TextArea";
import { useContext, useEffect, useState } from "react";
import { CollectionCreateContext } from "../layout";
import { RcFile } from "antd/es/upload";
import { calculateEtag } from "../../../../utils";

const normFile = (e: any) => {
  if (Array.isArray(e)) {
    return e;
  }
  return e?.fileList;
};

export default function Page() {
  const [form] = Form.useForm();
  const formValues = Form.useWatch([], form);

  const {
    name,
    description,
    image,
    imageURL,
    setName,
    setDescription,
    setImage,
    setImageURL,
    setValid,
  } = useContext(CollectionCreateContext);

  useEffect(() => {
    form
      .validateFields({
        validateOnly: true,
      })
      .then((o) => {
        // If image URL -> NFT image has been uploaded
        if (imageURL) {
          setValid(true);
        }
      })
      .catch((o) => {
        setValid(false);
      });
  }, [form, formValues, imageURL]);

  const uploadNFTImage = async (originFile: RcFile) => {
    const presignedReq = await fetch(
      "http://localhost:8080/images/upload-url",
      {
        headers: {
          "Content-Type": "application/json",
        },
        method: "POST",
        body: JSON.stringify({
          filename: "nft-image.png",
        }),
      }
    );

    const presignedBody = await presignedReq.json();
    const presignedURL = presignedBody.data.presigned_url;
    const url = presignedBody.data.url;

    const blob = originFile as Blob;
    const etag = await calculateEtag(blob);
    const putImage = await fetch(presignedURL, {
      method: "PUT",
      headers: {
        "Content-Type": "image/*",
        Etag: etag,
      },
      body: blob,
    });

    if (putImage.status === 200) {
      console.log("Image uploaded!");
    } else {
      console.log("Image upload failed");
    }

    setImageURL(url);
  };

  return (
    <main className="flex min-h-screen min-w-screen flex-col items-left">
      <div className="p-16 ">
        <h1 className="text-3xl font-bold">
          Detail the collection's information
        </h1>
        <p className="text-lg">
          There's no need to overthink this. This data is editable in future
          steps.
        </p>
        <div className="">
          <Form
            layout="vertical"
            form={form}
            initialValues={{
              name: name,
              description: description,
              image: image,
            }}
            validateMessages={{
              required: "${label} is required!",
            }}
          >
            <Form.Item
              name="image"
              label="Upload"
              valuePropName="file"
              getValueFromEvent={normFile}
              rules={[{ required: true }]}
              extra="This is the cover image your collection will show on OpenSea, Magic Eden, etc."
            >
              <Upload
                listType="picture-card"
                maxCount={1}
                onChange={(info: any) => {
                  // If image is selected
                  if (info.fileList.length > 0) {
                    const image = info.fileList[0].originFileObj;
                    setImage(image);
                    uploadNFTImage(image);
                  } else {
                    setImage("");
                  }
                }}
                fileList={
                  image
                    ? [
                        {
                          uid: "-1",
                          name: "image.png",
                          status: "done",
                          url: URL.createObjectURL(image),
                        },
                      ]
                    : []
                }
              >
                {!image && (
                  <div>
                    <FileImageOutlined />
                    <div style={{ marginTop: "10px" }}>Upload</div>
                  </div>
                )}
              </Upload>
            </Form.Item>
            <FormItem label="Name" name="name" rules={[{ required: true }]}>
              <Input
                onChange={(event) => {
                  setName(event.target.value);
                }}
                required
                value={name}
              />
            </FormItem>
            <FormItem
              label="Description"
              name="description"
              rules={[{ required: true }]}
            >
              <TextArea
                onChange={(event) => {
                  setDescription(event.target.value);
                }}
              />
            </FormItem>
          </Form>
        </div>
      </div>
    </main>
  );
}
