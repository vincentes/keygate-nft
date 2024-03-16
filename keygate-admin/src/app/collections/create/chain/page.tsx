"use client";

import { useContext, useEffect, useState } from "react";
import { CollectionCreateContext } from "../layout";

export default function Page() {
  const { blockchain, setBlockchain, setValid } = useContext(
    CollectionCreateContext
  );

  useEffect(() => {
    if (blockchain) {
      setValid(true);
    }
  }, [blockchain]);

  return (
    <main className="flex min-h-screen min-w-screen flex-col items-left">
      <div className="flex flex-col items-center p-16 text-center">
        <h1 className="text-3xl font-bold">Choose a chain</h1>
        <p className="text-lg">
          The collection will be deployed to the blockchain you choose.
        </p>
        <div
          className={`w-[320px] ease-in-out duration-300 cursor-pointer bg-white border-solid hover:border-sky-500 rounded-xl p-5 mt-5 font-bold shadow-sm border ${
            blockchain === "ethereum"
              ? "border-1 border-sky-500"
              : "border-white"
          }`}
          onClick={() => setBlockchain("ethereum")}
        >
          <img
            src="https://icons.iconarchive.com/icons/cjdowner/cryptocurrency-flat/256/Ethereum-ETH-icon.png"
            width={64}
          />
        </div>
        <div
          className={`w-[320px] ease-in-out duration-300 cursor-pointer bg-white border-solid hover:border-sky-500 rounded-xl p-5 mt-5 font-bold shadow-sm border ${
            blockchain === "bitcoin"
              ? "border-1 border-sky-500"
              : "border-white"
          }`}
          onClick={() => setBlockchain("bitcoin")}
        >
          <img
            src="https://upload.wikimedia.org/wikipedia/commons/thumb/4/46/Bitcoin.svg/1200px-Bitcoin.svg.png"
            width={64}
          />
        </div>
        <div
          className={`w-[320px] ease-in-out duration-300 cursor-pointer bg-white border-solid hover:border-sky-500 rounded-xl p-5 mt-5 font-bold shadow-sm border ${
            blockchain === "xrpl" ? "border-1 border-sky-500" : "border-white"
          }`}
          onClick={() => setBlockchain("xrpl")}
        >
          <img
            src="https://cryptologos.cc/logos/xrp-xrp-logo.png?v=029"
            width={64}
          />
        </div>
      </div>
    </main>
  );
}
