import { Button, Dropdown, MenuProps } from "antd";
import { FaPlus } from "react-icons/fa";
import { Panel, useNodes } from "reactflow";

const CreationActionButton = () => {
  const nodes = useNodes();

  const createKeyNode = () => {
    console.log("createKeyNode");

    const keyNode = {
      id: "key",
      type: "key",
      data: { label: "Key" },
      position: { x: 250, y: 250 },
    };

    nodes.push(keyNode);
  };

  const items: MenuProps["items"] = [
    {
      key: "1",
      label: <span>Asset</span>,
    },
    {
      key: "2",
      label: <button onClick={() => createKeyNode()}>Key</button>,
    },
    {
      key: "3",
      label: <span>Permission</span>,
    },
    {
      key: "4",
      label: <span>Vault</span>,
    },
  ];

  return (
    <Panel position="bottom-right" className="rounded-full">
      <Dropdown menu={{ items }} placement="topLeft" arrow>
        <Button className="text-sm rounded-full" size="large">
          <FaPlus />
        </Button>
      </Dropdown>
    </Panel>
  );
};

export default CreationActionButton;
