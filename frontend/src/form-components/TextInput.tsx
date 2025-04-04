import { Flex, Text } from "@jszymanowski/breeze-react";
import { Input } from "@/components/ui/input";
import { Label } from "@/form-components/Label";

export interface TextInputProps extends React.ComponentProps<"input"> {
  name: string;
  type?: string;
  placeholder?: string;
  label?: string;
  labelOrientation?: "top" | "left";
  note?: string;
  className?: string;
}

const TextInput = ({
  name,
  type = "text",
  placeholder,
  label,
  labelOrientation = "top",
  note,
  className,
  ...props
}: TextInputProps) => {
  return (
    <Flex direction={labelOrientation === "left" ? "row" : "col"} gap="2" className="w-full">
      {label && (
        <Label htmlFor={name} className={labelOrientation === "left" ? "text-nowrap" : ""}>
          {label}
        </Label>
      )}
      <Input
        id={name}
        name={name}
        type={type}
        placeholder={placeholder}
        className={className}
        {...props}
      />
      {note && (
        <Text size="sm" variant="muted">
          {note}
        </Text>
      )}
    </Flex>
  );
};

export { TextInput };
