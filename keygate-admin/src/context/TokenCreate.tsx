import { createContext } from "react";

export const TokenCreateContext = createContext({
  name: "",
  setName: {} as React.Dispatch<React.SetStateAction<string>>,
  description: "",
  setDescription: {} as React.Dispatch<React.SetStateAction<string>>,
});
