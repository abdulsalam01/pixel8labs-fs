import { useGetToken } from "@/hooks";
import { type NextPage } from "next";

const Callback: NextPage = () => {
  useGetToken();

  return (
    <div>
      <p>Processing OAuth callback...</p>
    </div>
  );
};

export default Callback;
