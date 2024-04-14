"use client";

import { useContext, useEffect } from "react";
import { CollectionCreateContext } from "../../../../context/CollectionCreate";

export default function Page() {
  const {
    name,
    description,
    image,
    price,
    blockchain,
    receiptAddress,
    setValid,
  } = useContext(CollectionCreateContext);

  useEffect(() => {
    setValid(true);
  }, []);

  return (
    <main className="flex min-h-screen min-w-screen flex-col items-left">
      <div className="flex flex-col p-16">
        <h1 className="text-xl font-bold">Review your collection details</h1>
        <p className="text-lg">
          Almost there! Take a final look and confirm your collection details
          before they get deployed to the blockchain.
        </p>
        <div className="flex flex-col space-y-2">
          {image && (
            // ignore error
            // @ts-ignore
            <img src={URL.createObjectURL(image)} className="w-64 h-64" />
          )}
          <div className="flex flex-row w-[300px] space-x-4 items-center">
            <div className="bg-light p-1">Name</div>
            <div>{name}</div>
          </div>
          <div className="flex flex-row w-[400px] space-x-4 items-start">
            <div className="bg-light p-1">Description</div>
            <div>{description}</div>
          </div>
          <div className="flex flex-row w-[200px] space-x-4 items-center">
            <div className="bg-light p-1">Price</div>
            <div>{price} ETH</div>
          </div>
          <div className="flex flex-row w-[200px] space-x-4 items-center">
            <div className="bg-light p-1">Blockchain</div>
            <div>{blockchain}</div>
          </div>
          <div className="flex flex-row w-[200px] space-x-4 items-center">
            <div className="bg-light p-1">Contract Type</div>
            <div>ERC721</div>
          </div>
          <div className="flex flex-row w-[200px] space-x-4 items-center">
            <div className="bg-light p-1">Receipt Address</div>
            <div>{receiptAddress}</div>
          </div>
        </div>
      </div>
    </main>
  );
}
