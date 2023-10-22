import { Inter } from "next/font/google";
const inter = Inter({ subsets: ["latin"] });

interface Props {
  text: string;
}

export const Heading = ({ text }: Props) => {
  return (
    <h1
      className={`${inter.className} flex items-center gap-3 text-2xl text-gray-800 font-bold`}
    >
      {text}
    </h1>
  );
};
