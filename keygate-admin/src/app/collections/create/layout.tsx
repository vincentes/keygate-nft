"use client";

import { CloseOutlined } from "@ant-design/icons";
import { Button, Progress } from "antd";
import { usePathname, useRouter } from "next/navigation";
import {
  Dispatch,
  SetStateAction,
  createContext,
  useMemo,
  useState,
} from "react";

const getNextPage = (pathname: string) => {
  const pages = [
    "/collections",
    "/collections/create/details",
    "/collections/create/chain",
    "/collections/create/price",
    "/collections/create/review",
    "/collections",
  ];
  const currentPageIndex = pages.indexOf(pathname);
  return pages[currentPageIndex + 1];
};

const getPercentage = (pathname: string) => {
  const pages = [
    "/collections/create/details",
    "/collections/create/chain",
    "/collections/create/price",
    "/collections/create/review",
  ];
  const currentPageIndex = pages.indexOf(pathname);
  return (currentPageIndex / pages.length) * 100;
};

const isLastPage = (pathname: string) => {
  return pathname === "/collections/create/review";
};

export const CollectionCreateContext = createContext({
  image: "",
  setImage: {} as Dispatch<SetStateAction<string>>,
  name: "",
  setName: {} as Dispatch<SetStateAction<string>>,
  description: "",
  setDescription: {} as Dispatch<SetStateAction<string>>,
  price: 0,
  setPrice: {} as Dispatch<SetStateAction<number>>,
  blockchain: "",
  setBlockchain: {} as Dispatch<SetStateAction<string>>,
  receiptAddress: "",
  setReceiptAddress: {} as Dispatch<SetStateAction<string>>,
  isValid: false,
  setValid: {} as Dispatch<SetStateAction<boolean>>,
});

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const router = useRouter();
  const currentPath = usePathname();
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [image, setImage] = useState("");
  const [price, setPrice] = useState(0);
  const [blockchain, setBlockchain] = useState("");
  const [receiptAddress, setReceiptAddress] = useState("");
  const [isValid, setValid] = useState(false);

  const contextValue = useMemo(
    () => ({
      name,
      setName,
      description,
      setDescription,
      image,
      setImage,
      price,
      setPrice,
      blockchain,
      setBlockchain,
      receiptAddress,
      setReceiptAddress,
      isValid,
      setValid,
    }),
    [name, description, image, price, blockchain, receiptAddress, isValid]
  );

  return (
    <>
      <div className="pl-8 pt-8" onClick={() => router.push("/collections")}>
        <CloseOutlined
          className="hover:cursor-pointer"
          style={{ fontSize: 20, fontStyle: "bold" }}
        />
      </div>
      <CollectionCreateContext.Provider value={contextValue}>
        {children}
      </CollectionCreateContext.Provider>
      <div className="fixed w-full bottom-[40px]">
        <Progress
          gapPosition="right"
          percent={getPercentage(currentPath)}
          status="active"
          size={"small"}
        />
        <div className="flex flex-row mt-5 justify-between px-5">
          <Button type="primary" disabled onClick={() => router.back()}>
            Back
          </Button>
          <Button
            type="primary"
            onClick={() => {
              router.push(getNextPage(currentPath));
              setValid(false);
            }}
            disabled={!isValid}
          >
            {isLastPage(currentPath) ? "Deploy" : "Next"}
          </Button>
        </div>
      </div>
    </>
  );
}
