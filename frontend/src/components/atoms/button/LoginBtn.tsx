import { useLogin } from "@/hooks/login";
import { Btn } from "./btn";
import { Inter } from "next/font/google";
const inter = Inter({ subsets: ["latin"] });

export const LoginBtn = () => {
  const handleLogin = useLogin();
  const handleClick = () => {
    handleLogin().catch((err) => {
      console.log(err);
    });
  };

  return (
    <Btn
      onClick={handleClick}
      className={`${inter.className} font-semibold text-white bg-custom-pink rounded-lg px-[18px] py-[10px]`}
    >
      Login with Github
    </Btn>
  );
};
