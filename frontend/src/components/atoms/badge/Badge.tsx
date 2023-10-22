import { Inter } from "next/font/google";
const inter = Inter({ subsets: ["latin"] });

interface Props {
  text: string;
  className?: string;
}

export const BadgeMd = ({ text }: Props) => {
  return (
    <div
      className={`${inter.className} text-sm bg-gray-50 rounded-full py-1 px-3`}
    >
      {text}
    </div>
  );
};

export const BadgeSm = ({ text, className }: Props) => {
  return (
    <div
      className={`${inter.className} bg-gray-50 text-xs rounded-full py-[2px] px-2 ${className}`}
    >
      {text}
    </div>
  );
};
