import { useState } from "react";
import DropdownCard from "./DropdownCard";

interface DropdownButtonProps {
  data: string[];
}

const DropdownButton = (props: DropdownButtonProps) => {
  const [open, setOpen] = useState(false);

  return (
    <div className="dropdown">
      <button onClick={() => setOpen(!open)} className="dropdown__button">
        Open
      </button>
      {open && <DropdownCard data={props.data} />}
    </div>
  );
};

export default DropdownButton;
