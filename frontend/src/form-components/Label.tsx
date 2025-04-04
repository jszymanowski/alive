import type { ComponentProps } from "react";
import { Label as CnLabel } from "@/components/ui/label";

import { cn } from "@/lib/utils";

interface Props extends ComponentProps<"label"> {
  value?: string;
  children?: React.ReactNode;
  className?: string;
}

const Label = ({ value, children, className = "", ...props }: Props) => (
  <CnLabel
    className={cn("font-display font-normal text-base text-foreground", className)}
    data-testid="form-label"
    {...props}
  >
    {value || children}
  </CnLabel>
);

export { Label };
