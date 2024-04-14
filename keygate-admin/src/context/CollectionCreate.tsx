import { createContext, Dispatch, SetStateAction } from "react";

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
  imageURL: "",
  setImageURL: {} as Dispatch<SetStateAction<string>>,
});
