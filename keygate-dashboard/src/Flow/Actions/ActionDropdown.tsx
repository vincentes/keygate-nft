import { Button, Dropdown, MenuProps } from "antd";
import { FaEllipsisV } from "react-icons/fa";

const items: MenuProps["items"] = [
  {
    key: "1",
    label: <span>View blockchain data</span>,
  },
];

const ActionDropdown = () => (
  <Dropdown menu={{ items }} placement="bottomLeft" arrow>
    <Button className="text-sm border-none" size="small">
      <FaEllipsisV />
    </Button>
  </Dropdown>
);

export default ActionDropdown;
