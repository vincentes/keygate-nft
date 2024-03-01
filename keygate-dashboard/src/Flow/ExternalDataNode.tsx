import { memo, type FC } from "react";
import { Handle, Position, type NodeProps } from "reactflow";

const CustomNode: FC<NodeProps> = ({ data, id }) => {
  return (
    <div className="bg-white rounded-xl fira-mono-medium border-solid border border-gray-200">
      <div className="flex flex-col text-left">
        <div className="text-xs border-b border-solid border-gray-200 pl-2 pr-10 py-2">
          {data.label.toLowerCase()}
        </div>
        <div className="p-5 text-light">
          Allows the user to view the drops page.
        </div>
      </div>
      <Handle
        type="target"
        position={Position.Top}
        id={id}
        isConnectable={true}
      />
    </div>
  );
};

export default memo(CustomNode);
