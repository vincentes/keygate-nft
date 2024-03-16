import { Switch } from "antd";
import { memo, useState, type FC } from "react";
import { Handle, Position, type NodeProps } from "reactflow";
import { FaKey } from "react-icons/fa";
import ActionDropdown from "./Actions/ActionDropdown";

const CustomNode: FC<NodeProps> = ({ data }) => {
  const [active, setActive] = useState(false);

  return (
    <div className="bg-white rounded-xl min-w-80 fira-mono-medium border-solid border border-gray-200">
      <div className="flex flex-col text-left w-full">
        <div className="text-xs border-b border-solid border-gray-200 pl-2 pr-2 py-2">
          <div className="flex flex-row justify-between">
            <div>
              <FaKey className="inline mr-4" />
              {data.label.toLowerCase()}
            </div>
            <ActionDropdown />
          </div>
        </div>
        <div className="py-2 px-2 flex flex-row space-x-2">
          <Switch onChange={setActive} />
          <div className="text-light">{active ? "active" : "inactive"}</div>
        </div>
      </div>
      <Handle type="source" position={Position.Bottom} id="b" />
    </div>
  );
};

export default memo(CustomNode);
