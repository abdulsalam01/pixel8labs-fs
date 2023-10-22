import { type StaticImport } from "next/dist/shared/lib/get-img-props";
import Image from "next/image";

interface Props {
  url: string | StaticImport;
}

export const Visitor = ({ url }: Props) => {
  return <Image src={url} width={56} height={56} alt="Avatar" />;
};
