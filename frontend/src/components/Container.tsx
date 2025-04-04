import { Box, type BoxProps } from "@jszymanowski/breeze-react";

interface Props extends BoxProps {
  children: React.ReactNode;
  className?: string;
}

export default function Container({ children, className, ...props }: Props) {
  return (
    <Box width="full" className={`max-w-6xl p-12 pt-0 align-left ${className}`} {...props}>
      {children}
    </Box>
  );
}
