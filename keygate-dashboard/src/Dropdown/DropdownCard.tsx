interface DropdownCardProps {
  data: string[];
}

const DropdownCard = ({ data = [] }: DropdownCardProps) => (
  <div className="">
    {data.map((item: string) => (
      <div key={item} className="dropdown-card__item">
        {item}
      </div>
    ))}
  </div>
);

export default DropdownCard;
