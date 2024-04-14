"use client";

import { Form, Input } from "antd";
import { useContext, useEffect } from "react";
import { CollectionCreateContext } from "../../../../context/CollectionCreate";

export default function Page() {
  const { price, setPrice, receiptAddress, setReceiptAddress, setValid } =
    useContext(CollectionCreateContext);

  const [form] = Form.useForm();
  const formValues = Form.useWatch([], form);

  useEffect(() => {
    form
      .validateFields({
        validateOnly: true,
      })
      .then(() => {
        setValid(true);
      })
      .catch(() => {
        setValid(false);
      });
  }, [form, formValues]);

  return (
    <main className="flex min-h-screen min-w-screen flex-col items-left">
      <div className="flex flex-col items-center p-16">
        <h1 className="text-3xl font-bold">Pricing</h1>
        <p className="text-lg">How much is your NFT going to cost?</p>
        <div
          className={`md:w-[600px] ease-in-out duration-300 cursor-pointer bg-white rounded-xl p-12 mt-5  shadow-sm border`}
        >
          <Form
            className="flex flex-col space-y-4"
            form={form}
            initialValues={{
              price: price ? price : "",
              receiptAddress,
            }}
            layout="vertical"
            onChange={() => {}}
          >
            <div className="font-bold">Pricing</div>
            <Form.Item
              label="Price"
              name="price"
              rules={[
                {
                  required: true,
                  message: "Please enter the price of the NFT.",
                },
              ]}
            >
              <Input
                addonAfter={
                  <div className="flex flex-row items-center opacity-50">
                    <span>ETH</span>
                  </div>
                }
                onChange={(e) => setPrice(parseFloat(e.target.value))}
              />
            </Form.Item>

            <Form.Item
              label="Receipt Address"
              name="receiptAddress"
              rules={[
                {
                  required: true,
                  message: "An address is required.",
                },
                {
                  len: 42,
                  message: "Please enter a valid EVM-based address.",
                },
              ]}
            >
              <Input
                onChange={(event) => {
                  setReceiptAddress(event.target.value);
                }}
              />
            </Form.Item>

            <div className="text-sm text-center">
              Enter the address where you want to receive the revenue from the
              sales of the NFTs. Funds may take a few days to arrive.
            </div>
          </Form>
        </div>
      </div>
    </main>
  );
}
