import axios from "axios";
import { memo, useState, type FC } from "react";
import { FaDoorClosed } from "react-icons/fa";
import { Handle, Position, type NodeProps } from "reactflow";

interface User {
  id: string;
  externalId: string;
}

const statusToBorder: Record<string, string> = {
  pending: "border-gray-200",
  ok: "border-green-500",
  notok: "border-red-500",
};

const PermissionNode: FC<NodeProps> = ({ data, id }) => {
  const [verification, setVerification] = useState<string>("pending");
  // const [users, setUsers] = useState<User[]>([]);
  const [externalUserId, setExternalUserId] = useState<string>("");

  const verifyUser = async () => {
    // get /users/{externalid}/permissions
    try {
      try {
        await axios({
          method: "HEAD",
          url: `http://localhost:8080/ext/users/${externalUserId}/permissions/${id}`,
        });
        setVerification("ok");
      } catch (error) {
        setVerification("notok");
      }

      setTimeout(() => {
        setVerification("pending");
      }, 1500);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div
      className={`bg-white rounded-xl fira-mono-medium border-solid border transition-all delay-400 ${statusToBorder[verification]}`}
    >
      <div className="flex flex-col text-left">
        <div
          className={`flex flex-row text-xs justify-between pl-2 pr-5 py-2 border-b border-solid border-gray-200 transition-all ${statusToBorder[verification]}`}
        >
          <div className="">
            <FaDoorClosed className="inline mr-4" />
            {data.label.toLowerCase()}
          </div>
          <div className="text-light space-x-2">+</div>
        </div>
        <div className="p-5 text-light">
          Allows the user to view the drops page.
        </div>
        <div className="flex flex-row p-5 text-light space-x-2 ">
          <input
            type="text"
            className={"w-full pl-4 transition-all duration-200 rounded-md"}
            placeholder="Enter an external user id"
            onChange={(e) => setExternalUserId(e.target.value)}
          />
          <button
            className="relative bg-blue-500 text-white rounded-md px-4 py-2"
            onClick={() => verifyUser()}
          >
            {/*
            float outside button
            */}
            Verify
          </button>
        </div>
        {/* <div className="flex flex-row p-5 text-light space-x-2">
          <div>Allows</div>
          <Select
            showSearch
            className="inline w-full"
            placeholder="Select an action"
            options={[
              { label: "View", value: "view" },
              { label: "Edit", value: "edit" },
              { label: "Delete", value: "delete" },
            ]}
          />
        </div>
        <div className="flex flex-row p-5 text-light space-x-2">
          <div>Entity</div>
          <Select
            showSearch
            className="inline w-full"
            placeholder="Select an entity"
            options={[{ label: "Drops", value: "drops" }]}
          />
        </div> */}
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

export default memo(PermissionNode);
